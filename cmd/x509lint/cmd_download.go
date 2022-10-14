package main

import (
	"crypto"
	"crypto/tls"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/peculiarventures/x509-linter/cmd/internal"
	"github.com/peculiarventures/x509-linter/pkg/url"
)

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

	// Disable SSL verification
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	links := strings.Split(string(raw), "\n")
	lintResults := []*url.LintUrlResultSet{}
	for _, link := range links {
		// use http links only
		if !strings.HasPrefix(link, "http") {
			continue
		}

		// lint url
		lintResult := url.LintUrl(link)
		lintResults = append(lintResults, lintResult)
		if lintResult.StatusCode != 200 {
			continue
		}

		certs := internal.ParseCertificates(lintResult.Body)
		files := []string{}
		for _, pemCert := range certs {
			cert := pemCert.Certificate
			if cert.IsCA && !includeCa {
				// skip CA cert if includeCa is false
				continue
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

		fmt.Printf("%-7s%s\n", "OK", link)
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	}

	reportDir := "x_report"
	if err := Mkdir(reportDir); err != nil {
		return fmt.Errorf("cannot create directory %s, %s", reportDir, err.Error())
	}
	urlDir := path.Join(reportDir, "url")
	if err := Mkdir(urlDir); err != nil {
		return fmt.Errorf("cannot create directory %s, %s", urlDir, err.Error())
	}

	file, err := os.Create(path.Join(urlDir, "README.md"))
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	PrintUrlSummary(file, lintResults)

	return nil
}

func GetStatusText(status url.LintUrlStatus) string {
	switch status {
	case url.Error:
		return "error"
	case url.Warn:
		return "warn"
	case url.Notice:
		return "notice"
	}
	return "unknown"
}

func PrintUrlSummary(w io.Writer, r []*url.LintUrlResultSet) {
	fmt.Fprintln(w, "# STIR/SHAKEN Certificate Repository Compliance")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "Participants in the STIR/SHAKEN ecosystem are required to publish the certificates they use for signing calls so that other participants can retrieve these certificates and use them for validating signatures.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "To ease meeting this requirement many of the certificate authorities publish the certificates for their customers and the majority of the STIR/SHAKEN deployments use these CA-provided repositories. There are still some cases where customers choose to host the certificate on their own.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "This report looks at what errors and compliance issues relying parties may experience when trying to use these certificate repositories.")
	fmt.Fprintln(w, "")

	for _, v := range r {
		fmt.Fprintf(w, "%s\n", v.Url)
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "| Code | Status | Source | Details |")
		fmt.Fprintln(w, "|------|--------|--------|---------|")
		counter := 0
		for code, l := range v.Results {
			if l.Status != url.Pass {
				counter += 1
				ruleInfo := url.GetRuleByName(code)
				fmt.Fprintf(w, "| %s | %s | %s | %s |\n", code, GetStatusText(l.Status), ruleInfo.Source, l.Details)
			}
		}
		fmt.Fprintln(w, "")
		if counter == 0 {
			fmt.Fprintf(w, "%d tests were ran and no warning or error level issues were found\n", len(v.Results))
			fmt.Fprintln(w, "")
		}
	}
}

func Mkdir(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory %s, %s", name, err.Error())
		}
	}

	return nil
}
