package main

import (
	"crypto"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/peculiarventures/x509-linter/cmd/internal"
	lintUrl "github.com/peculiarventures/x509-linter/pkg/url"
	"github.com/zmap/zcrypto/x509"
)

type LintUrlSummaryResult struct {
	Organizations map[string]*LintUrlOrgResult
	LintResult
}

func NewLintUrlSummaryResult() *LintUrlSummaryResult {
	return &LintUrlSummaryResult{
		Organizations: map[string]*LintUrlOrgResult{},
	}
}

func (t *LintUrlSummaryResult) AppendLink(name string, l *lintUrl.LintUrlResultSet) {
	org := t.Organizations[name]
	if org == nil {
		org = NewLintUrlOrgResult()
		org.Name = name
		t.Organizations[name] = org
	}

	org.AppendLink(l)

	// update counters
	t.Amount += 1
	if l.HasErrors {
		t.Errors += 1
	}
	if l.HasWarnings {
		t.Warnings += 1
	}
	if l.HasNotices {
		t.Notices += 1
	}
}

type LintUrlOrgResult struct {
	Name     string
	Links    []*lintUrl.LintUrlResultSet
	Problems map[string]int
	LintResult
}

func NewLintUrlOrgResult() *LintUrlOrgResult {
	return &LintUrlOrgResult{
		Problems: map[string]int{},
	}
}

func (t *LintUrlOrgResult) AppendLink(l *lintUrl.LintUrlResultSet) {
	t.Links = append(t.Links, l)

	for code, r := range l.Results {
		if r.Status != lintUrl.Pass {
			t.Problems[code] += 1
		}
	}

	// update counters
	t.Amount += 1
	if l.HasErrors {
		t.Errors += 1
	}
	if l.HasWarnings {
		t.Warnings += 1
	}
	if l.HasNotices {
		t.Notices += 1
	}
}

func RunDownloadCommand(listPath string, outDir string, includeCa bool) error {
	if listPath == "" {
		return fmt.Errorf("cannot get file path, variable 'listPath' is empty")
	}

	raw, err := os.ReadFile(listPath)
	if err != nil {
		return fmt.Errorf("cannot read file %s, %s", listPath, err.Error())
	}

	// create folder
	if _, err := os.ReadDir(outDir); err != nil {
		if err := os.Mkdir(outDir, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory %s, %s", outDir, err.Error())
		}
	}

	rootPool, err := ReadRootCertificates("shaken-ca-list.pem")
	if err != nil {
		return fmt.Errorf("cannot create root CAs, %s", err.Error())
	}

	links := strings.Split(string(raw), "\n")
	lintResults := NewLintUrlSummaryResult()
	for _, link := range links {
		// use http links only
		if !strings.HasPrefix(link, "http") {
			continue
		}

		// lint lintUrl
		lintResult := lintUrl.LintUrl(link)
		if lintResult.StatusCode != 200 {
			lintResults.AppendLink("Unknown", lintResult)
			continue
		}

		certs := internal.ParseCertificates(lintResult.Body)
		files := []string{}
		intermediatePool := x509.NewCertPool()
		var leafCert *x509.Certificate
		for _, pemCert := range certs {
			cert := pemCert.Certificate
			if cert.IsCA {
				intermediatePool.AddCert(cert)
				if !includeCa {
					// skip CA cert if includeCa is false
					continue
				}
			} else {
				leafCert = cert
			}

			// compute cert sha1 thumbprint
			sha1 := crypto.SHA1.New()
			sha1.Write(cert.Raw)

			fileName := fmt.Sprintf("%s.pem", hex.EncodeToString(sha1.Sum(nil)))
			filePath := path.Join(outDir, fileName)
			if _, err := os.Stat(filePath); err == nil {
				// file exists, skip it
				files = append(files, fmt.Sprintf("exists %s", filePath))
				continue
			}

			// write cert into pem file
			certBlock := pem.Block{
				Type: "CERTIFICATE",
				Headers: map[string]string{
					"Link":    link,
					"Subject": cert.Subject.String(),
					"Issuer":  cert.Issuer.String(),
				},
				Bytes: cert.Raw,
			}
			pemCert := pem.EncodeToMemory(&certBlock)
			err = os.WriteFile(filePath, pemCert, 0644)
			if err != nil {
				fmt.Printf("%-7s%s %s\n", "ERROR", link, err.Error())
				continue
			}
			files = append(files, fmt.Sprintf("write  %s", filePath))
		}

		orgName := "Unknown"
		if leafCert != nil {
			orgName = getOrganizationName(leafCert, &x509.VerifyOptions{
				Intermediates: intermediatePool,
				Roots:         rootPool,
			})
		}
		lintResults.AppendLink(orgName, lintResult)

		fmt.Printf("%-7s%s\n", "OK", link)
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	}

	reportDir := "x_report"
	if err := Mkdir(reportDir); err != nil {
		return fmt.Errorf("cannot create directory %s, %s", reportDir, err.Error())
	}

	file, err := os.Create(path.Join(reportDir, "URL.md"))
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	PrintUrlSummary(file, lintResults)

	for orgName, lintOrg := range lintResults.Organizations {
		// create org dir
		orgDir := path.Join(reportDir, orgName)
		err := Mkdir(orgDir)
		if err != nil {
			return fmt.Errorf("cannot create organization directory %s, %s", orgDir, err.Error())
		}

		// create org report file
		orgFile, err := os.Create(path.Join(orgDir, "URL.md"))
		if err != nil {
			return fmt.Errorf("cannot create organization report, %s", err.Error())
		}
		defer orgFile.Close()

		PrintUrlOrg(orgFile, lintOrg)

		// create issue report file
		issuesDir := path.Join(orgDir, "ISSUES")
		err = Mkdir(issuesDir)
		if err != nil {
			return fmt.Errorf("cannot create issues report, %s", err.Error())
		}
		for code := range lintOrg.Problems {
			issueFile, err := os.Create(path.Join(issuesDir, fmt.Sprintf("%s.md", code)))
			if err != nil {
				return fmt.Errorf("cannot create issues report, %s", err.Error())
			}
			defer issueFile.Close()
			PrintUrlIssue(issueFile, code, lintOrg)
		}
	}

	return nil
}

func GetStatusText(status lintUrl.LintUrlStatus) string {
	switch status {
	case lintUrl.Error:
		return "error"
	case lintUrl.Warn:
		return "warn"
	case lintUrl.Notice:
		return "notice"
	}
	return "unknown"
}

func PrintUrlOrg(w io.Writer, r *LintUrlOrgResult) {
	fmt.Fprintln(w, "# STIR/SHAKEN Certificate Repository Compliance")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "## %s\n", r.Name)
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "| Code | Source | Instances |")
	fmt.Fprintln(w, "|------|--------|-----------|")
	for code, instances := range r.Problems {
		rule := lintUrl.GetRuleByName(code)
		codeLink := fmt.Sprintf("[%s](ISSUES/%s.md)", code, code)
		fmt.Fprintf(w, "| %s | %s | %d |\n", codeLink, rule.Source, instances)
	}
	fmt.Fprintln(w, "")

	for _, v := range r.Links {
		fmt.Fprintf(w, "### %s\n", v.Url)
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "| Code | Status | Source | Details |")
		fmt.Fprintln(w, "|------|--------|--------|---------|")
		counter := 0
		for code, l := range v.Results {
			if l.Status != lintUrl.Pass {
				counter += 1
				ruleInfo := lintUrl.GetRuleByName(code)
				fmt.Fprintf(w, "| %s | %s | %s | %s |\n", code, GetStatusText(l.Status), ruleInfo.Source, l.Details)
			}
		}
		fmt.Fprintln(w, "")
		if counter == 0 {
			fmt.Fprintf(w, "%d tests were ran and no warning or error level issues were found\n", len(v.Results))
			fmt.Fprintln(w, "")
		}
	}

	PrintFooter(w)
}

func PrintUrlSummary(w io.Writer, r *LintUrlSummaryResult) {
	fmt.Fprintln(w, "# STIR/SHAKEN Certificate Repository Compliance")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "Participants in the STIR/SHAKEN ecosystem are required to publish the certificates they use for signing calls so that other participants can retrieve these certificates and use them for validating signatures.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "To ease meeting this requirement many of the certificate authorities publish the certificates for their customers and the majority of the STIR/SHAKEN deployments use these CA-provided repositories. There are still some cases where customers choose to host the certificate on their own.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "This report looks at what errors and compliance issues relying parties may experience when trying to use these certificate repositories.")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "| Issuers | Links | Errors | Warnings | Notices |")
	fmt.Fprintln(w, "|---------|-------|--------|----------|---------|")

	// order r.Issuers keys
	keys := make([]string, 0, len(r.Organizations))
	for k := range r.Organizations {
		keys = append(keys, k)
	}
	sort.Slice(keys[:], func(a int, b int) bool {
		return keys[a] < keys[b]
	})

	for _, key := range keys {
		issuer := r.Organizations[key]
		issuerNameLink := fmt.Sprintf("[%s](%s)", key, url.PathEscape(path.Join(key, "URL.md")))
		fmt.Fprintf(w, "| %s | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", issuerNameLink, issuer.Amount, percent(issuer.Amount, r.Amount), issuer.Errors, percent(issuer.Errors, issuer.Amount), issuer.Warnings, percent(issuer.Warnings, issuer.Amount), issuer.Notices, percent(issuer.Notices, issuer.Amount))
	}
	fmt.Fprintf(w, "| **Total** | %d (100%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", r.Amount, r.Errors, percent(r.Errors, r.Amount), r.Warnings, percent(r.Warnings, r.Amount), r.Notices, percent(r.Notices, r.Amount))

	PrintFooter(w)
}

func Mkdir(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory %s, %s", name, err.Error())
		}
	}

	return nil
}

func ReadRootCertificates(path string) (*x509.CertPool, error) {
	p := x509.NewCertPool()
	rootPem, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	rootCerts := internal.ParseCertificates(rootPem)
	for _, cert := range rootCerts {
		p.AddCert(cert.Certificate)
	}

	return p, nil
}

func PrintUrlIssue(w io.Writer, code string, r *LintUrlOrgResult) {
	fmt.Fprintln(w, "# STIR/SHAKEN Certificate Repository Compliance")
	fmt.Fprintln(w, "")

	fmt.Fprintf(w, "## %s\n", r.Name)
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "Code: %s\\\n", code)
	rule := lintUrl.GetRuleByName(code)
	if rule != nil {
		fmt.Fprintf(w, "Source: %s\\\n", rule.Source)
		fmt.Fprintf(w, "Description: %s\n", rule.Description)
	}

	for _, link := range r.Links {
		if !link.HasProblem(code) {
			continue
		}

		fmt.Fprintf(w, "### %s\n", link.Url)
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "| Code | Status | Source | Details |")
		fmt.Fprintln(w, "|------|--------|--------|---------|")
		counter := 0
		for code, l := range link.Results {
			if l.Status != lintUrl.Pass {
				counter += 1
				ruleInfo := lintUrl.GetRuleByName(code)
				fmt.Fprintf(w, "| %s | %s | %s | %s |\n", code, GetStatusText(l.Status), ruleInfo.Source, l.Details)
			}
		}
		fmt.Fprintln(w, "")
		if counter == 0 {
			fmt.Fprintf(w, "%d tests were ran and no warning or error level issues were found\n", len(link.Results))
			fmt.Fprintln(w, "")
		}
	}

	PrintFooter(w)
}
