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
		if err := Mkdir(reportDir); err != nil {
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

		PrintCertificateReport(os.Stdout, result)
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
			atis1000080.ATIS_Source,
			atis1000080.CPv1_3_Source,
			atis1000080.PKI_Source,
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
	Name         string
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

func NewLintCertificatesResult() *LintCertificatesResult {
	return &LintCertificatesResult{
		Issuers:  map[string]*LintOrganizationResult{},
		Findings: &Findings{},
	}
}

func (t *LintCertificatesResult) AppendCertificate(c *LintCertificateResult) {
	issuer := t.Issuers[c.Organization]
	if issuer == nil {
		issuer = &LintOrganizationResult{
			Name:     c.Organization,
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

type LintTotalResult struct {
	Issues           map[string]int
	LeafCertificates *LintCertificatesResult
	CaCertificates   *LintCertificatesResult
}

func NewLintTotalResult() *LintTotalResult {
	return &LintTotalResult{
		Issues:           map[string]int{},
		LeafCertificates: NewLintCertificatesResult(),
		CaCertificates:   NewLintCertificatesResult(),
	}
}

func (t *LintTotalResult) AppendCertificate(r *LintCertificateResult) {
	// update Issues counter
	for code, result := range r.Result.Results {
		if result.Status == lint.Error ||
			result.Status == lint.Warn ||
			result.Status == lint.Notice ||
			result.Status == lint.NE {
			t.Issues[code] += 1
		}
	}

	if r.Cert.IsCA {
		t.CaCertificates.AppendCertificate(r)
		return
	}

	t.LeafCertificates.AppendCertificate(r)
}

// GetOrganizationsNames returns ordered list of unique names fro Leaf and CA issuers
func (t *LintTotalResult) GetOrganizationsNames() []string {
	nameMap := map[string]bool{}
	// read all names for Leaf certs
	for n := range t.LeafCertificates.Issuers {
		nameMap[n] = true
	}
	// read all names for CA certs
	for n := range t.CaCertificates.Issuers {
		nameMap[n] = true
	}

	res := []string{}
	for n := range nameMap {
		res = append(res, n)
	}

	// sort names
	sort.Slice(res[:], func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
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

func LintCertificates(certs []*internal.PemCertificate, options *x509.VerifyOptions) *LintTotalResult {
	res := NewLintTotalResult()

	for _, cert := range certs {
		certRes, err := LintCertificate(cert, options)
		if err != nil {
			continue
		}

		res.AppendCertificate(certRes)
	}

	return res
}

// SaveOrganizationReport writes report for each organization
func SaveOrganizationReport(r *LintTotalResult, outDir string) error {

	names := r.GetOrganizationsNames()
	for _, name := range names {
		// create folder
		orgDir := path.Join(outDir, name)
		if err := Mkdir(orgDir); err != nil {
			return fmt.Errorf("cannot create directory %s, %s", name, err.Error())
		}

		// create file
		orgFile := path.Join(orgDir, "README.md")
		file, err := os.Create(orgFile)
		if err != nil {
			return fmt.Errorf("cannot create %s, %w", orgFile, err)
		}
		defer file.Close()

		PrintOrganizationReport(file, name, r)

		for code := range r.Issues {
			if HasIssue(name, code, r) {
				SaveIssueGroupReport(name, code, r, orgDir)
			}
		}
	}

	return nil
}

// HasIssue returns true if LeafCertificates or CaCertificates for selected orgName contains issueName and it has problems
func HasIssue(orgName string, issueName string, r *LintTotalResult) bool {
	certsSet := []*LintCertificatesResult{
		r.LeafCertificates,
		r.CaCertificates,
	}
	for _, certs := range certsSet {
		issuer := certs.Issuers[orgName]
		if issuer != nil {
			issue := issuer.Issues[issueName]
			if issue != nil {
				return true
			}
		}
	}

	return false
}

func SaveIssueGroupReport(orgName string, issueName string, r *LintTotalResult, orgDir string) error {
	// check if Leaf and CA has this issue

	issuesDir := path.Join(orgDir, "ISSUES")
	err := Mkdir(issuesDir)
	if err != nil {
		return fmt.Errorf("cannot save IssueGroup report, %w", err)
	}

	reportFile := path.Join(issuesDir, fmt.Sprintf("%s.md", issueName))
	file, err := os.Create(reportFile)
	if err != nil {
		return fmt.Errorf("cannot save IssueGroup report, %w", err)
	}
	defer file.Close()

	PrintIssueGroupReport(file, orgName, issueName, r)

	return nil
}

func SaveTotalReport(r *LintTotalResult, outDir string) error {
	file, err := os.Create(path.Join(outDir, "README.md"))
	if err != nil {
		return fmt.Errorf("cannot save README.md, %w", err)
	}
	defer file.Close()

	PrintTotalReport(file, r)

	return nil
}

// PrintFindingList prints the list of findings
func PrintFindingList(w io.Writer, f *Findings) {
	if f.LeafCertificates > 0 {
		fmt.Fprintf(w, "- Average validity span as of leaf certificates %d days\n", f.ValidityDays/int(f.LeafCertificates))
		fmt.Fprintf(w, "- Percentage of leaf certificates expiring in the next 30 days is %0.2f%%\n", percent(f.SoonExpiredCertificates, f.LeafCertificates))
	}
}

func SaveCertificatesReport(r *LintTotalResult, outDir string) error {
	names := r.GetOrganizationsNames()
	for _, name := range names {
		issuers := []*LintOrganizationResult{
			r.LeafCertificates.Issuers[name],
			r.CaCertificates.Issuers[name],
		}
		for _, issuer := range issuers {
			if issuer == nil {
				continue
			}
			for _, cert := range issuer.Certificates {
				// create folder
				certDir := path.Join(outDir, name, cert.Thumbprint)
				if err := Mkdir(certDir); err != nil {
					return fmt.Errorf("cannot create directory %s, %s", certDir, err.Error())
				}

				// create file
				certFile := path.Join(path.Join(certDir, "README.md"))
				file, err := os.Create(certFile)
				if err != nil {
					return fmt.Errorf("cannot create %s, %w", certFile, err)
				}
				defer file.Close()

				PrintCertificateReportForIssuer(file, name, cert)
			}
		}
	}

	return nil
}

func PrintFooter(w io.Writer) {
	now := time.Now()
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "Generated: %s at %s", now.Format("02/01/2006"), now.Format("15:04:05"))
}

// PrintTotalReport prints the total report
func PrintTotalReport(w io.Writer, r *LintTotalResult) {
	fmt.Fprintln(w, "# STIR/SHAKEN CA Ecosystem Compliance")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "[Approved Certificate Authorities](https://authenticate.iconectiv.com/approved-certification-authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "This report is generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main Zlint distribution but can be found [here](https://github.com/PeculiarVentures/x509-linter).")

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "## Summary")
	fmt.Fprintln(w, "")
	if r.LeafCertificates.Amount > 0 {
		fmt.Fprintln(w, "### Leaf Certificates")
		fmt.Fprintln(w, "")
		PrintFindingList(w, r.LeafCertificates.Findings)
		PrintOrganizationsTable(w, r.LeafCertificates, "leaf-certificates")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "\\* The percent of certificates per issuer is calculated against total certificates from all issuers\\")
		fmt.Fprintln(w, "\\*\\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer\\")
		fmt.Fprintln(w, "\\*\\*\\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time")
		fmt.Fprintln(w, "\\*\\*\\*\\* For information on failed certificate repository retrievals see this [report](url/README.md)")
	}
	if r.CaCertificates.Amount > 0 {
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "### CA Certificates")
		fmt.Fprintln(w, "")
		PrintOrganizationsTable(w, r.CaCertificates, "ca-certificates")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "\\* The percent of certificates per issuer is calculated against total certificates from all issuers\\")
		fmt.Fprintln(w, "\\*\\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer")
	}

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "## Key")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "| Type | Description |")
	fmt.Fprintln(w, "|------|-------------|")
	fmt.Fprintln(w, "| Error | Tests in which the specifications are unambiguous on what the expected behavior must be. |")
	fmt.Fprintln(w, "| Warning	| Tests in which the specifications are ambiguous or are provide only a recommendation. |")
	fmt.Fprintln(w, "| Notice | Tests in which industry best practices are not followed. |")
	fmt.Fprintln(w, "| Not Effective	| Tests that exist in the current specifications but were not in effect at the time of issuance. |")

	PrintFooter(w)
}

func PrintOrganizationReport(w io.Writer, name string, r *LintTotalResult) {
	// header
	fmt.Fprintln(w, "# STIR/SHAKEN CA Ecosystem Compliance")

	// summary
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "## %s\n", name)

	issuersKeys := []string{
		"Leaf Certificates",
		"CA Certificates",
	}
	issuers := map[string]*LintOrganizationResult{
		issuersKeys[0]: r.LeafCertificates.Issuers[name],
		issuersKeys[1]: r.CaCertificates.Issuers[name],
	}
	for _, issuerType := range issuersKeys {
		issuer := issuers[issuerType]
		if issuer == nil {
			continue
		}
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "### %s\n", issuerType)
		fmt.Fprintln(w, "")
		PrintFindingList(w, issuer.Findings)
		fmt.Fprintf(w, "- Certificates with Errors: %d\n", issuer.Errors)
		fmt.Fprintf(w, "- Certificates with Warnings: %d\n", issuer.Warnings)
		fmt.Fprintf(w, "- Certificates with Notices: %d\n", issuer.Notices)
		fmt.Fprintf(w, "- Certificates with tests not executed as the requirements were Not Effective at issuance time: %d\n", issuer.NE)
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "| Status | Code | Source | Instances |")
		fmt.Fprintln(w, "|--------|------|--------|-----------|")
		for code, issue := range issuer.Issues {
			rule := lint.GlobalRegistry().ByName(code)
			codeLink := fmt.Sprintf("[%s](ISSUES/%s.md#%s)", code, code, strings.ReplaceAll(strings.ToLower(issuerType), " ", "-"))
			fmt.Fprintf(w, "| %s | %s | %s | %d |\n", statusToString(issue.Type), codeLink, rule.Source, issue.Amount)
		}

		// summery footer
		// TODO don't show for CA
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "\\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time")

		// certificates
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "#### Issued certificates")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "| Created at | Name | Problems | Link |")
		fmt.Fprintln(w, "|------------|------|----------|------|")
		sort.Slice(issuer.Certificates[:], func(i, j int) bool {
			return issuer.Certificates[i].Cert.NotBefore.Unix() < issuer.Certificates[j].Cert.NotBefore.Unix()
		})
		for _, certReport := range issuer.Certificates {
			fmt.Fprintf(w, "| %s | %s | %s | %s |\n",
				certReport.Cert.NotBefore.Format(time.RFC822),                                            // created at
				certReport.Cert.Subject.CommonName,                                                       // name
				fmt.Sprintf("%t", certReport.Result.ErrorsPresent || certReport.Result.WarningsPresent),  // problems
				fmt.Sprintf("[view](%s)", url.PathEscape(path.Join(certReport.Thumbprint, "README.md"))), // link
			)
		}
	}

	PrintFooter(w)
}

// PrintCertificateReport prints the certificate report
func PrintCertificateReport(w io.Writer, r *LintCertificateResult) {
	fmt.Fprintf(w, "### Certificate %s\n", r.Thumbprint)
	fmt.Fprintf(w, "Tested At: %s\\\n", time.Unix(r.Result.Timestamp, 0).String())
	fmt.Fprintf(w, "Initial Validity Period: %d day(s)\\\n", internal.GetValidityDays(r.Cert))
	remainingDays := internal.GetRemainingDays(r.Cert, time.Now())
	remaining := fmt.Sprintf("%d day(s)", remainingDays)
	if remainingDays < 1 {
		remaining = "Expired"
	}
	fmt.Fprintf(w, "Remaining Validity Period: %s\\\n", remaining)
	fmt.Fprintf(w, "Subject: %s\\\n", strings.ReplaceAll(r.Cert.Subject.String(), "\\", "\\\\"))
	fmt.Fprintf(w, "Issuer: %s\n\n", strings.ReplaceAll(r.Cert.Issuer.String(), "\\", "\\\\"))
	fmt.Fprintf(w, "Link: %s\n\n", r.Link)
	fmt.Fprintf(w, "View: [Click to view](https://understandingwebpki.com/?cert=%s)\n\n", url.QueryEscape(base64.StdEncoding.EncodeToString(r.Cert.Raw)))
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "| Code | Type | Source | Details |\n")
	fmt.Fprintf(w, "|------|------|--------|---------|\n")
	for code, result := range r.Result.Results {
		if result.Status == lint.Error ||
			result.Status == lint.Warn ||
			result.Status == lint.Notice {
			rule := lint.GlobalRegistry().ByName(code)
			fmt.Fprintf(w, "| %s | %s | %s | %s |\n", code, statusToString(result.Status), rule.Source, result.Details)
		}
	}

	if !r.Result.ErrorsPresent && !r.Result.WarningsPresent {
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "%d tests were ran and no warning or error level issues were found\n", len(r.Result.Results))
	}

	neHeader := false
	for code, result := range r.Result.Results {
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
		fmt.Fprintln(w, "\\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time")
	}
}

// PrintCertificateReportForIssuer prints the certificate report for specified issuer
func PrintCertificateReportForIssuer(w io.Writer, issuerName string, r *LintCertificateResult) {
	fmt.Fprintln(w, "# STIR/SHAKEN CA Ecosystem Compliance")
	fmt.Fprintf(w, "## %s\n", issuerName)
	fmt.Fprintln(w, "")
	PrintCertificateReport(w, r)

	PrintFooter(w)
}

func PrintOrganizationsTable(w io.Writer, r *LintCertificatesResult, anchor string) {
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |")
	fmt.Fprintln(w, "|---------|--------------|--------|----------|---------|---------------|")

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
		issuerNameLink := fmt.Sprintf("[%s](%s)", key, fmt.Sprintf("%s#%s", url.PathEscape(path.Join(key, "README.md")), anchor))
		fmt.Fprintf(w, "| %s | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", issuerNameLink, issuer.Amount, percent(issuer.Amount, r.Amount), issuer.Errors, percent(issuer.Errors, issuer.Amount), issuer.Warnings, percent(issuer.Warnings, issuer.Amount), issuer.Notices, percent(issuer.Notices, issuer.Amount), issuer.NE, percent(issuer.NE, issuer.Amount))
	}
	fmt.Fprintf(w, "| **Total** | %d (100%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", r.Amount, r.Errors, percent(r.Errors, r.Amount), r.Warnings, percent(r.Warnings, r.Amount), r.Notices, percent(r.Notices, r.Amount), r.NE, percent(r.NE, r.Amount))
	// TODO add a footnote
	// fmt.Fprintln(w, "")
	// fmt.Fprintf(w, "* 140 tests are included in this test suite")
}

func PrintIssueGroupReport(w io.Writer, orgName string, issueName string, r *LintTotalResult) {
	// header
	fmt.Fprintln(w, "# STIR/SHAKEN CA Ecosystem Compliance")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "## %s", orgName)
	fmt.Fprintln(w, "")
	issueInfo := lint.GlobalRegistry().ByName(issueName)
	fmt.Fprintf(w, "Name: %s\\\n", issueInfo.Name)
	fmt.Fprintf(w, "Source: %s\\\n", issueInfo.Source)
	fmt.Fprintf(w, "Citation: %s\\\n", issueInfo.Citation)
	fmt.Fprintf(w, "Effective Date: %s\\\n", issueInfo.EffectiveDate.Format(time.RFC822))
	fmt.Fprintf(w, "Description: %s\n", issueInfo.Description)
	fmt.Fprintln(w, "")

	leafIssuer := r.LeafCertificates.Issuers[orgName]
	if leafIssuer != nil {
		fmt.Fprintln(w, "### Leaf Certificates")
		fmt.Fprintln(w, "")
		PrintIssueGroupCertificateTable(w, issueName, leafIssuer)
	}

	caIssuer := r.CaCertificates.Issuers[orgName]
	if caIssuer != nil {
		fmt.Fprintln(w, "### CA Certificates")
		fmt.Fprintln(w, "")
		PrintIssueGroupCertificateTable(w, issueName, caIssuer)
	}

	PrintFooter(w)
}

func PrintIssueGroupCertificateTable(w io.Writer, issueName string, org *LintOrganizationResult) {
	fmt.Fprintln(w, "| Status | Subject | Link | Details |")
	fmt.Fprintln(w, "|--------|---------|------|---------|")
	counter := 0
	for _, cert := range org.Certificates {
		issue := cert.Result.Results[issueName]
		if issue == nil || issue.Status == lint.Pass || issue.Status == lint.Fatal || issue.Status == lint.NA || issue.Status == lint.Reserved {
			continue
		}
		subject := strings.ReplaceAll(cert.Cert.Subject.String(), "\\", "\\\\")
		link := fmt.Sprintf("[view](%s)", path.Join("..", computeCertThumbprint(cert.Cert), "README.md"))
		fmt.Fprintf(w, "| %s | %s | %s | %s |\n", statusToString(issue.Status), subject, link, issue.Details)
		counter += 1
	}

	fmt.Fprintln(w, "")
	if counter == 0 {
		fmt.Fprintln(w, "no warning, or error, or not effective date level issues were found")
		fmt.Fprintln(w, "")
	}
}
