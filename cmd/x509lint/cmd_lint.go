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
	"strings"
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

		err = SaveTotalReport(r, "")
		if err != nil {
			return err
		}
		err = SaveCertificatesReport(r, "")
		if err != nil {
			return err
		}
		err = SaveOrganizationReport(r, "")
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
	Link   string
	Cert   *x509.Certificate
	Result *zlint.ResultSet
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
		Link:   link,
		Cert:   cert,
		Result: zlint.LintCertificateEx(cert, registry),
	}, nil
}

func printResultMarkDown(w io.Writer, info *LintCertificateResult) {
	thumbprint := sha1.New()
	thumbprint.Write(info.Cert.Raw)
	fmt.Fprintf(w, "### Certificate %s\n", hex.EncodeToString(thumbprint.Sum(nil)))
	fmt.Fprintf(w, "Tested At: %s\n\n", time.Unix(info.Result.Timestamp, 0).String())
	fmt.Fprintf(w, "Subject: %s\n\n", info.Cert.Subject.String())
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
	fmt.Fprintln(w, "* The percent of certificates per issuer is calculated against total certificates from all issuers")
	fmt.Fprintln(w, "** The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer")
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
	file, err := os.Create(path.Join(outDir, "organization.md"))
	if err != nil {
		return fmt.Errorf("cannot save organization.md, %w", err)
	}
	defer file.Close()

	fmt.Fprintln(file, "# SHAKEN COMPLIANCE SUMMARY")
	fmt.Fprintln(file, "## Organizations")

	for name, issuer := range r.Issuers {
		fmt.Fprintln(file, "")
		fmt.Fprintf(file, "### %s\n", name)
		fmt.Fprintln(file, "")
		fmt.Fprintf(file, "Errors: %d\n", issuer.Errors)
		fmt.Fprintln(file, "")
		fmt.Fprintf(file, "Warnings: %d\n", issuer.Warnings)
		fmt.Fprintln(file, "")
		fmt.Fprintln(file, "| Status | Code | Amount |")
		fmt.Fprintln(file, "|--------|------|--------|")
		for code, issue := range issuer.Issues {
			fmt.Fprintf(file, "| %s | %s | %d |\n", issue.Type, code, issue.Amount)
		}
	}

	return nil
}

func SaveTotalReport(r *LintCertificatesResult, outDir string) error {
	file, err := os.Create(path.Join(outDir, "total.md"))
	if err != nil {
		return fmt.Errorf("cannot save total.md, %w", err)
	}
	defer file.Close()

	fmt.Fprintln(file, "# SHAKEN COMPLIANCE SUMMARY")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "| Issuers | Certificates | Errors | Warnings |")
	fmt.Fprintln(file, "|---------|--------------|--------|----------|")
	for issuerName, issuer := range r.Issuers {
		issuerNameLink := fmt.Sprintf("[%s](organization.md#%s)", issuerName, strings.ReplaceAll(issuerName, " ", "-"))
		fmt.Fprintf(file, "| %s | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", issuerNameLink, issuer.Amount, percent(issuer.Amount, r.Amount), issuer.Errors, percent(issuer.Errors, issuer.Amount), issuer.Warnings, percent(issuer.Warnings, issuer.Amount))
	}
	fmt.Fprintf(file, "| **Total** | %d (100%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", r.Amount, r.Errors, percent(r.Errors, r.Amount), r.Warnings, percent(r.Warnings, r.Amount))

	return nil
}

func SaveCertificatesReport(r *LintCertificatesResult, outDir string) error {
	file, err := os.Create(path.Join(outDir, "certs.md"))
	if err != nil {
		return fmt.Errorf("cannot save certs.md, %w", err)
	}
	defer file.Close()

	fmt.Fprintln(file, "# SHAKEN COMPLIANCE SUMMARY")
	fmt.Fprintln(file, "## Certificates")

	for _, issuer := range r.Issuers {
		for _, cert := range issuer.Certificates {
			fmt.Fprintln(file, "")
			printResultMarkDown(file, cert)
		}
	}

	return nil
}
