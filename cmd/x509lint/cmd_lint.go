package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/peculiarventures/x509-linter/cmd/internal"
	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

func RunLintCommand(certPath string, summary bool) error {
	if certPath == "" {
		return fmt.Errorf("cannot get certificate path, variable 'certPath' is empty")
	}

	// This returns an *os.FileInfo type
	file, err := os.Open(certPath)
	if err != nil {
		return fmt.Errorf("cannot open %s. %s", certPath, err.Error())
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("cannot get stat info from the %s. %w", certPath, err)
	}

	if fileInfo.IsDir() {
		// load root certs
		checkCerts := []*internal.PemCertificate{}
		rootPool := x509.NewCertPool()
		intermediatePool := x509.NewCertPool()
		rootPem, err := os.ReadFile("shaken-ca-list.pem")
		if err == nil {
			rootCerts := internal.ParseCertificates(rootPem)
			for _, cert := range rootCerts {
				rootPool.AddCert(cert.Certificate)
			}
		}

		files, err := ioutil.ReadDir(certPath)
		if err != nil {
			return fmt.Errorf("cannot read the directory %s, %w", certPath, err)
		}

		for _, file := range files {
			if !file.IsDir() {
				certPath := path.Join(certPath, file.Name())
				certRaw, err := os.ReadFile(certPath)
				if err != nil {
					continue
				}
				certs := internal.ParseCertificates(certRaw)
				for _, cert := range certs {
					if cert.Certificate.IsCA {
						// add intermediate certs into pool
						intermediatePool.AddCert(cert.Certificate)
					}
					// add leaf and intermediate certs into check list
					checkCerts = append(checkCerts, cert)
				}
			}
		}

		r := LintCertificates(checkCerts, &x509.VerifyOptions{
			Roots:         rootPool,
			Intermediates: intermediatePool,
		})

		reportDir := "x_report"
		if err := os.Mkdir(reportDir, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory %s, %s", reportDir, err.Error())
		}

		err = SaveTotalReport(r, reportDir)
		if err != nil {
			return err
		}
		err = SaveOrganizationReport(r, reportDir)
		if err != nil {
			return err
		}
		err = SaveCertificatesReport(r, reportDir)
		if err != nil {
			return err
		}
	} else {
		// file is not a directory

		raw, err := os.ReadFile(certPath)
		if err != nil {
			return fmt.Errorf("cannot read the file %s, %s", certPath, err.Error())
		}

		certs := internal.ParseCertificates(raw)
		if len(certs) == 0 {
			return fmt.Errorf("cannot read the certificate from the file %s, %s", certPath, err.Error())
		}

		result, err := LintCertificate(certs[0], nil)
		if err != nil {
			return fmt.Errorf("cannot lint the certificate, %w", err)
		}

		printResultMarkDown(os.Stdout, result)
	}

	return nil
}

type LintCertificateResult struct {
	Link         string
	Cert         *x509.Certificate
	Thumbprint   string
	Result       *zlint.ResultSet
	Organization string
}

var wellknownIssueNames = map[string]string{
	"Comcast":                          "Comcast",
	"GBSDTech":                         "GBSDTech",
	"Martini Security, LLC":            "Martini Security",
	"Metaswitch STI-CA SHAKEN Root":    "Metaswitch",
	"NetNumber Inc":                    "NetNumber",
	"Neustar Information Services Inc": "Neustar",
	"Peeringhub Inc":                   "Peeringhub",
	"Sansay Corporation":               "Sansay",
	"TMOBILE-USA":                      "T-Mobile",
	"TransNexus, Inc.":                 "TransNexus",
}

func LintCertificate(cert *internal.PemCertificate, options *x509.VerifyOptions) (*LintCertificateResult, error) {
	// Initialize lint registry
	registry, err := lint.GlobalRegistry().Filter(lint.FilterOptions{
		// IncludeSources: lint.SourceList{shaken.ShakenPolicy},
		IncludeSources: lint.SourceList{
			lint.RFC5280,
			atis1000080.SHAKEN,
		},
		ExcludeNames: []string{
			"w_distribution_point_missing_ldap_or_uri",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cannot initialize the lint registry, %s", err.Error())
	}

	link := cert.Headers["Link"]
	if len(link) == 0 {
		link = ""
	}

	// get organization name
	organization := internal.GetOrganizationName(cert.Certificate)
	if options != nil {
		current, expired, never, err := cert.Certificate.Verify(*options)
		if err != nil {
			if len(current) > 0 {
				chain := current[0]
				organization = internal.GetOrganizationName(chain[len(chain)-1])
			} else if len(expired) > 0 {
				chain := expired[0]
				organization = internal.GetOrganizationName(chain[len(chain)-1])
			} else if len(never) > 0 {
				chain := never[0]
				organization = internal.GetOrganizationName(chain[len(chain)-1])
			}
		}
	}

	if len(wellknownIssueNames[organization]) != 0 {
		organization = wellknownIssueNames[organization]
	}

	thumbprint := computeCertThumbprint(cert.Certificate)
	fmt.Printf("Lint certificate %s issued by '%s' (%s)\n", thumbprint, organization, link)

	return &LintCertificateResult{
		Link:         link,
		Cert:         cert.Certificate,
		Thumbprint:   thumbprint,
		Result:       zlint.LintCertificateEx(cert.Certificate, registry),
		Organization: organization,
	}, nil
}

func computeCertThumbprint(c *x509.Certificate) string {
	thumbprint := sha1.New()
	thumbprint.Write(c.Raw)
	return hex.EncodeToString(thumbprint.Sum(nil))
}

func statusToString(s lint.LintStatus) string {
	switch s {
	case lint.Error:
		return "error"
	case lint.Warn:
		return "warn"
	case lint.Notice:
		return "notice"
	case lint.NE:
		return "not effective"
	default:
		return s.String()
	}
}

func printResultMarkDown(w io.Writer, info *LintCertificateResult) {
	fmt.Fprintf(w, "### Certificate %s\n", info.Thumbprint)
	fmt.Fprintf(w, "Tested At: %s\\\n", time.Unix(info.Result.Timestamp, 0).String())
	fmt.Fprintf(w, "Initial Validity Period: %d day(s)\\\n", internal.GetValidityDays(info.Cert))
	remainingDays := internal.GetRemainingDays(info.Cert, time.Now())
	remaining := fmt.Sprintf("%d day(s)", remainingDays)
	if remainingDays < 1 {
		remaining = "Expired"
	}
	fmt.Fprintf(w, "Remaining Validity Period: %s\\\n", remaining)
	fmt.Fprintf(w, "Subject: %s\\\n", strings.ReplaceAll(info.Cert.Subject.String(), "\\", "\\\\"))
	fmt.Fprintf(w, "Issuer: %s\n\n", strings.ReplaceAll(info.Cert.Issuer.String(), "\\", "\\\\"))
	fmt.Fprintf(w, "Link: %s\n\n", info.Link)
	fmt.Fprintf(w, "View: [Click to view](https://understandingwebpki.com/?cert=%s)\n\n", url.QueryEscape(base64.StdEncoding.EncodeToString(info.Cert.Raw)))
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "| Code | Type | Details |\n")
	fmt.Fprintf(w, "|------|------|---------|\n")
	for code, result := range info.Result.Results {
		if result.Status == lint.Error ||
			result.Status == lint.Warn ||
			result.Status == lint.Notice {
			fmt.Fprintf(w, "| %s | %s | %s |\n", code, statusToString(result.Status), result.Details)
		}
	}

	if !info.Result.ErrorsPresent && !info.Result.WarningsPresent {
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "%d tests were ran and non warning or error level issues were found\n", len(info.Result.Results))
	}

	neHeader := false
	for code, result := range info.Result.Results {
		if result.Status == lint.NE {
			if !neHeader {
				// Print header only once
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "### Not Effective")
				fmt.Fprintln(w, "")
				neHeader = true
			}

			fmt.Fprintf(w, "- %s\n", code)
		}
	}

	if neHeader {
		// Issue footer
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "\\* Tests use the ATIS 1000080 and Certificate Policy versions. Certificates issued before these dates are not evaluated as the rules may not have been enforce at the time")
	}
}

func percent(a uint, b uint) float64 {
	return float64(a) / float64(b) * 100
}

type LintResult struct {
	Amount   uint
	Errors   uint
	Warnings uint
	Notices  uint
	NE       uint
}

type Findings struct {
	LeafCertificates        uint
	ValidityDays            int
	SoonExpiredCertificates uint
}

type LintIssue struct {
	Type   lint.LintStatus
	Amount uint
}

type LintOrganizationResult struct {
	Issues       map[string]*LintIssue
	Certificates []*LintCertificateResult
	LintResult
	*Findings
}

func (t *LintOrganizationResult) AppendCertificate(c *LintCertificateResult) {
	t.Certificates = append(t.Certificates, c)
	nePresents := false

	// Update Issues
	for code, result := range c.Result.Results {
		// filter results by Status
		if !(result.Status == lint.Error ||
			result.Status == lint.Warn ||
			result.Status == lint.Notice ||
			result.Status == lint.NE) {
			continue
		}
		issue := t.Issues[code]
		if issue == nil {
			issue = &LintIssue{
				Type: result.Status,
			}
			t.Issues[code] = issue
		}

		issue.Amount += 1

		// c.Result doesn't have NEPresents
		if !nePresents && result.Status == lint.NE {
			nePresents = true
		}
	}

	// Update counters
	t.Amount += 1
	if c.Result.ErrorsPresent {
		t.Errors += 1
	}
	if c.Result.WarningsPresent {
		t.Warnings += 1
	}
	if c.Result.NoticesPresent {
		t.Notices += 1
	}
	if nePresents {
		t.NE += 1
	}

	// finding
	if !c.Cert.IsCA {
		t.LeafCertificates += 1
		t.ValidityDays += internal.GetValidityDays(c.Cert)

		if time.Now().AddDate(0, 0, 30).After(c.Cert.NotAfter) {
			t.SoonExpiredCertificates += 1
		}
	}
}

type LintCertificatesResult struct {
	Issuers map[string]*LintOrganizationResult
	LintResult
	*Findings
}

func (t *LintCertificatesResult) AppendCertificate(c *LintCertificateResult) {
	issuer := t.Issuers[c.Organization]
	if issuer == nil {
		issuer = &LintOrganizationResult{
			Issues:   map[string]*LintIssue{},
			Findings: &Findings{},
		}
		t.Issuers[c.Organization] = issuer
	}

	issuer.AppendCertificate(c)

	// Update counters
	t.Amount += 1
	if c.Result.ErrorsPresent {
		t.Errors += 1
	}
	if c.Result.WarningsPresent {
		t.Warnings += 1
	}
	if c.Result.NoticesPresent {
		t.Notices += 1
	}
	for _, issue := range c.Result.Results {
		if issue.Status == lint.NE {
			t.NE += 1
			break
		}
	}

	// finding
	t.LeafCertificates += issuer.LeafCertificates
	t.ValidityDays += issuer.ValidityDays
	t.SoonExpiredCertificates += issuer.SoonExpiredCertificates
}

func ReadCertificatesDir(dirPath string) ([]*internal.PemCertificate, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read the '%s' directory, %w", dirPath, err)
	}

	res := []*internal.PemCertificate{}

	for _, file := range files {
		if !file.IsDir() {
			certPath := path.Join(dirPath, file.Name())
			certRaw, err := os.ReadFile(certPath)
			if err != nil {
				continue
			}

			certs := internal.ParseCertificates(certRaw)
			res = append(res, certs...)
		}
	}

	return res, nil
}

func LintCertificates(certs []*internal.PemCertificate, options *x509.VerifyOptions) *LintCertificatesResult {
	res := &LintCertificatesResult{
		Issuers:  map[string]*LintOrganizationResult{},
		Findings: &Findings{},
	}

	for _, cert := range certs {
		// TODO ignore CA certs
		if cert.Certificate.IsCA {
			continue
		}

		certRes, err := LintCertificate(cert, options)
		if err != nil {
			continue
		}

		res.AppendCertificate(certRes)
	}

	return res
}

// SaveOrganizationReport writes report for each organization
func SaveOrganizationReport(r *LintCertificatesResult, outDir string) error {

	for name, issuer := range r.Issuers {
		// create folder
		orgDir := path.Join(outDir, name)
		if err := os.Mkdir(orgDir, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory %s, %s", name, err.Error())
		}

		// create file
		orgFile := path.Join(orgDir, "README.md")
		file, err := os.Create(orgFile)
		if err != nil {
			return fmt.Errorf("cannot create %s, %w", orgFile, err)
		}
		defer file.Close()

		// header
		fmt.Fprintln(file, "# STIR/SHAKEN CA Ecosystem Compliance")

		// summary
		fmt.Fprintln(file, "")
		fmt.Fprintf(file, "## %s\n", name)

		fmt.Fprintln(file, "")
		PrintFindingList(file, issuer.Findings)
		fmt.Fprintf(file, "- Errors: %d\n", issuer.Errors)
		fmt.Fprintf(file, "- Warnings: %d\n", issuer.Warnings)
		fmt.Fprintf(file, "- Notices: %d\n", issuer.Notices)
		fmt.Fprintf(file, "- Not Effective: %d\n", issuer.NE)
		fmt.Fprintln(file, "")
		fmt.Fprintln(file, "| Status | Code | Instances |")
		fmt.Fprintln(file, "|--------|------|-----------|")
		for code, issue := range issuer.Issues {
			fmt.Fprintf(file, "| %s | %s | %d |\n", statusToString(issue.Type), code, issue.Amount)
		}

		// summery footer
		fmt.Fprintln(file, "")
		fmt.Fprintln(file, "\\* Tests use the ATIS 1000080 and Certificate Policy versions. Certificates issued before these dates are not evaluated as the rules may not have been enforce at the time")

		// certificates
		fmt.Fprintln(file, "")
		fmt.Fprintln(file, "## Issued certificates")
		fmt.Fprintln(file, "")
		fmt.Fprintln(file, "| Created at | Name | Problems | Link |")
		fmt.Fprintln(file, "|------------|------|----------|------|")
		sort.Slice(issuer.Certificates[:], func(i, j int) bool {
			return issuer.Certificates[i].Cert.NotBefore.Unix() < issuer.Certificates[j].Cert.NotBefore.Unix()
		})
		for _, certReport := range issuer.Certificates {
			fmt.Fprintf(file, "| %s | %s | %s | %s |\n",
				certReport.Cert.NotBefore.Format(time.RFC822),                                            // created at
				certReport.Cert.Subject.CommonName,                                                       // name
				fmt.Sprintf("%t", certReport.Result.ErrorsPresent || certReport.Result.WarningsPresent),  // problems
				fmt.Sprintf("[view](%s)", url.PathEscape(path.Join(certReport.Thumbprint, "README.md"))), // link
			)
		}

		PrintFooter(file)
	}

	return nil
}

func SaveTotalReport(r *LintCertificatesResult, outDir string) error {
	file, err := os.Create(path.Join(outDir, "README.md"))
	if err != nil {
		return fmt.Errorf("cannot save README.md, %w", err)
	}
	defer file.Close()

	fmt.Fprintln(file, "# STIR/SHAKEN CA Ecosystem Compliance")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "[Approved Certificate Authorities](https://authenticate.iconectiv.com/approved-certification-authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "This report is generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main Zlint distribution but can be found [here](https://github.com/PeculiarVentures/x509-linter).")

	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "## Summary")
	fmt.Fprintln(file, "")
	PrintFindingList(file, r.Findings)
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |")
	fmt.Fprintln(file, "|---------|--------------|--------|----------|---------|---------------|")

	// order r.Issuers keys
	keys := make([]string, 0, len(r.Issuers))
	for k := range r.Issuers {
		keys = append(keys, k)
	}
	sort.Slice(keys[:], func(a int, b int) bool {
		return keys[a] < keys[b]
	})

	for _, key := range keys {
		issuer := r.Issuers[key]
		issuerNameLink := fmt.Sprintf("[%s](%s)", key, url.PathEscape(path.Join(key, "README.md")))
		fmt.Fprintf(file, "| %s | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", issuerNameLink, issuer.Amount, percent(issuer.Amount, r.Amount), issuer.Errors, percent(issuer.Errors, issuer.Amount), issuer.Warnings, percent(issuer.Warnings, issuer.Amount), issuer.Notices, percent(issuer.Notices, issuer.Amount), issuer.NE, percent(issuer.NE, issuer.Amount))
	}
	fmt.Fprintf(file, "| **Total** | %d (100%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", r.Amount, r.Errors, percent(r.Errors, r.Amount), r.Warnings, percent(r.Warnings, r.Amount), r.Notices, percent(r.Notices, r.Amount), r.NE, percent(r.NE, r.Amount))

	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "\\* The percent of certificates per issuer is calculated against total certificates from all issuers\\")
	fmt.Fprintln(file, "\\*\\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer\\")
	fmt.Fprintln(file, "\\*\\*\\* Tests use the ATIS 1000080 and Certificate Policy versions. Certificates issued before these dates are not evaluated as the rules may not have been enforce at the time")

	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "## Key")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "| Type | Description |")
	fmt.Fprintln(file, "|------|-------------|")
	fmt.Fprintln(file, "| Error | Tests in which the specifications are unambiguous on what the expected behavior must be. |")
	fmt.Fprintln(file, "| Warning	| Tests in which the specifications are ambiguous or are provide only a recommendation. |")
	fmt.Fprintln(file, "| Notice | Tests in which industry best practices are not followed. |")
	fmt.Fprintln(file, "| Not Effective	| Tests that exist in the current specifications but were not in effect at the time of issuance. |")

	PrintFooter(file)

	return nil
}

func PrintFindingList(file *os.File, findings *Findings) {
	if findings.LeafCertificates > 0 {
		fmt.Fprintf(file, "- Average validity span as of leaf certificates %d days\n", findings.ValidityDays/int(findings.LeafCertificates))
		fmt.Fprintf(file, "- Percentage of leaf certificates expiring in the next 30 days is %0.2f%%\n", percent(findings.SoonExpiredCertificates, findings.LeafCertificates))
	}
}

func SaveCertificatesReport(r *LintCertificatesResult, outDir string) error {

	for issuerName, issuer := range r.Issuers {
		for _, cert := range issuer.Certificates {
			// create folder
			certDir := path.Join(outDir, issuerName, cert.Thumbprint)
			if err := os.Mkdir(certDir, os.ModePerm); err != nil {
				return fmt.Errorf("cannot create directory %s, %s", certDir, err.Error())
			}

			// create file
			certFile := path.Join(path.Join(certDir, "README.md"))
			file, err := os.Create(certFile)
			if err != nil {
				return fmt.Errorf("cannot create %s, %w", certFile, err)
			}
			defer file.Close()

			// header
			fmt.Fprintln(file, "# STIR/SHAKEN CA Ecosystem Compliance")
			fmt.Fprintf(file, "## %s\n", issuerName)
			fmt.Fprintln(file, "")
			printResultMarkDown(file, cert)

			PrintFooter(file)
		}
	}

	return nil
}

func PrintFooter(file *os.File) {
	now := time.Now()
	fmt.Fprintln(file, "")
	fmt.Fprintf(file, "Generated: %s at %s", now.Format("02/01/2006"), now.Format("15:04:05"))
}
