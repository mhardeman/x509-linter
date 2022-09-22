package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

func RunLintCommand(certPath string) error {
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
		// file is a directory
		files, err := ioutil.ReadDir(certPath)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("# SHAKEN lint report")
		fmt.Println("## Certificates")

		for _, file := range files {
			if !file.IsDir() {

				filePath := path.Join(certPath, file.Name())
				result, err := doLint(filePath)
				if err != nil {
					continue
				}
				fmt.Println("")
				printResultMarkDown(result)
			}
		}
	} else {
		// file is not a directory
		result, err := doLint(certPath)
		if err != nil {
			return fmt.Errorf("cannot lint the certificate, %w", err)
		}

		printResultMarkDown(result)
	}

	return nil
}

type LintResult struct {
	Link   string
	Cert   *x509.Certificate
	Result *zlint.ResultSet
}

func doLint(certPath string) (*LintResult, error) {
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

	return &LintResult{
		Link:   link,
		Cert:   cert,
		Result: zlint.LintCertificateEx(cert, registry),
	}, nil
}

func printResultMarkDown(info *LintResult) {
	thumbprint := sha1.New()
	thumbprint.Write(info.Cert.Raw)
	fmt.Printf("### Certificate %s\n", hex.EncodeToString(thumbprint.Sum(nil)))
	fmt.Printf("Timestamp: %s\n\n", time.Unix(info.Result.Timestamp, 0).String())
	fmt.Printf("Subject: %s\n\n", info.Cert.Subject.String())
	fmt.Printf("Issuer: %s\n\n", info.Cert.Issuer.String())
	fmt.Printf("Link: %s\n\n", info.Link)
	fmt.Printf("View: [Click to view](https://understandingwebpki.com/?cert=%s)\n\n", url.QueryEscape(base64.StdEncoding.EncodeToString(info.Cert.Raw)))
	fmt.Println("")
	fmt.Printf("| Code | Type | Details |\n")
	fmt.Printf("|------|------|---------|\n")
	for code, result := range info.Result.Results {
		if result.Status == lint.Error || result.Status == lint.Warn {
			fmt.Printf("| %s | %s | %s |\n", code, result.Status, result.Details)
		}
	}
}
