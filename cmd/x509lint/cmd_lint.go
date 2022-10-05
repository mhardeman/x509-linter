package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"sort"
	"time"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

type AuthorityReport struct {
	Certificates uint
	Warnings     uint
	Errors       uint
}

type SummaryReport struct {
	Authorities map[string]*AuthorityReport
	*AuthorityReport
}

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
		r, err := LintCertificates(certPath)
		if err != nil {
			return fmt.Errorf("cannot lint the directory with certificates, %w", err)
		}

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
		result, err := LintCertificate(certPath)
		if err != nil {
			return fmt.Errorf("cannot lint the certificate, %w", err)
		}

		printResultMarkDown(os.Stdout, result)
	}

	return nil
}

type LintCertificateResult struct {
	Link       string
	Cert       *x509.Certificate
	Thumbprint string
	Result     *zlint.ResultSet
}

func LintCertificate(certPath string) (*LintCertificateResult, error) {
	raw, err := os.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read the certificate file %s, %s", certPath, err.Error())
	}

	// Parse cert
	p, _ := pem.Decode(raw)
	if len(p.Bytes) > 0 {
		// try parse pem and if it's ok replace raw to bytes from the pem
		if p.Type != "CERTIFICATE" {
			return nil, fmt.Errorf("incorrect PEM data, tag value is %s, but should be 'CERTIFICATE'", p.Type)
		}
		raw = p.Bytes
	}
	cert, err := x509.ParseCertificate(raw)
	if err != nil {
		return nil, fmt.Errorf("cannot parse the certificate %s, %s", certPath, err.Error())
	}

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

	config, err := lint.NewConfigFromString(`
	[w_organization_identity]
	Some = "test"
	`)
	if err != nil {
		return nil, fmt.Errorf("cannot create config, error: %s", err.Error())
	}
	registry.SetConfiguration(config)

	link := p.Headers["Link"]
	if link == "" {
		link = certPath
	}

	return &LintCertificateResult{
		Link:       link,
		Cert:       cert,
		Thumbprint: computeCertThumbprint(cert),
		Result:     zlint.LintCertificateEx(cert, registry),
	}, nil
}

func computeCertThumbprint(c *x509.Certificate) string {
	thumbprint := sha1.New()
	thumbprint.Write(c.Raw)
	return hex.EncodeToString(thumbprint.Sum(nil))
}

func printResultMarkDown(w io.Writer, info *LintCertificateResult) {
	fmt.Fprintf(w, "### Certificate %s\n", info.Thumbprint)
	fmt.Fprintf(w, "Tested At: %s\\\n", time.Unix(info.Result.Timestamp, 0).String())
	fmt.Fprintf(w, "Subject: %s\\\n", info.Cert.Subject.String())
	fmt.Fprintf(w, "Issuer: %s\n\n", info.Cert.Issuer.String())
	fmt.Fprintf(w, "Link: %s\n\n", info.Link)
	fmt.Fprintf(w, "View: [Click to view](https://understandingwebpki.com/?cert=%s)\n\n", url.QueryEscape(base64.StdEncoding.EncodeToString(info.Cert.Raw)))
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "| Code | Type | Details |\n")
	fmt.Fprintf(w, "|------|------|---------|\n")
	for code, result := range info.Result.Results {
		if result.Status == lint.Error || result.Status == lint.Warn {
			fmt.Fprintf(w, "| %s | %s | %s |\n", code, result.Status, result.Details)
		}
	}

	if !info.Result.ErrorsPresent && !info.Result.WarningsPresent {
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "%d tests were ran and no issues were found\n", len(info.Result.Results))
	}

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "\\* The percent of certificates per issuer is calculated against total certificates from all issuers\\")
	fmt.Fprintln(w, "\\*\\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer")
}

func percent(a uint, b uint) float64 {
	return float64(a) / float64(b) * 100
}

type LintResult struct {
	Amount   uint
	Errors   uint
	Warnings uint
}

type LintIssue struct {
	Type   lint.LintStatus
	Amount uint
}

type LintOrganizationResult struct {
	Issues       map[string]*LintIssue
	Certificates []*LintCertificateResult
	LintResult
}

func (t *LintOrganizationResult) AppendCertificate(c *LintCertificateResult) {
	t.Certificates = append(t.Certificates, c)

	// Update Issues
	for code, result := range c.Result.Results {
		if result.Status != lint.Error && result.Status != lint.Warn {
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
	}

	// Update counters
	t.Amount += 1
	if c.Result.ErrorsPresent {
		t.Errors += 1
	}
	if c.Result.WarningsPresent {
		t.Warnings += 1
	}
}

type LintCertificatesResult struct {
	Issuers map[string]*LintOrganizationResult
	LintResult
}

func (t *LintCertificatesResult) AppendCertificate(c *LintCertificateResult) {
	organization := "Unknown"
	if len(c.Cert.Issuer.Organization) > 0 {
		organization = c.Cert.Issuer.Organization[0]
	}

	issuer := t.Issuers[organization]
	if issuer == nil {
		issuer = &LintOrganizationResult{
			Issues: map[string]*LintIssue{},
		}
		t.Issuers[organization] = issuer
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
}

func LintCertificates(dirPath string) (*LintCertificatesResult, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read the directory, %w", err)
	}

	res := &LintCertificatesResult{
		Issuers: map[string]*LintOrganizationResult{},
	}

	for _, file := range files {
		if !file.IsDir() {
			certPath := path.Join(dirPath, file.Name())
			certRes, err := LintCertificate(certPath)
			if err != nil {
				continue
			}

			res.AppendCertificate(certRes)
		}
	}

	return res, nil
}

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
		fmt.Fprintf(file, "Errors: %d\\\n", issuer.Errors)
		fmt.Fprintf(file, "Warnings: %d\n", issuer.Warnings)
		fmt.Fprintln(file, "")
		fmt.Fprintln(file, "| Status | Code | Amount |")
		fmt.Fprintln(file, "|--------|------|--------|")
		for code, issue := range issuer.Issues {
			fmt.Fprintf(file, "| %s | %s | %d |\n", issue.Type, code, issue.Amount)
		}

		// certificates
		fmt.Fprintln(file, "")
		fmt.Fprintln(file, "### Issued certificates")
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
	fmt.Fprintln(file, "| Issuers | Certificates | Errors | Warnings |")
	fmt.Fprintln(file, "|---------|--------------|--------|----------|")
	for issuerName, issuer := range r.Issuers {
		issuerNameLink := fmt.Sprintf("[%s](%s)", issuerName, url.PathEscape(path.Join(issuerName, "README.md")))
		fmt.Fprintf(file, "| %s | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", issuerNameLink, issuer.Amount, percent(issuer.Amount, r.Amount), issuer.Errors, percent(issuer.Errors, issuer.Amount), issuer.Warnings, percent(issuer.Warnings, issuer.Amount))
	}
	fmt.Fprintf(file, "| **Total** | %d (100%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", r.Amount, r.Errors, percent(r.Errors, r.Amount), r.Warnings, percent(r.Warnings, r.Amount))

	return nil
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
		}
	}

	return nil
}
