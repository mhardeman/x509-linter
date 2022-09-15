package main

import (
	"encoding/pem"
	"flag"
	"log"
	"os"

	"github.com/peculiarventures/x509-linter/pkg/shaken"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

func main() {
	flag.Parse()

	// Read certPath
	var certPath = flag.Arg(0)
	if certPath == "" {
		log.Fatalf("cannot get certificate path, variable 'certPath' is empty")
	}

	result := doLint(certPath)
	printLintResult(result)
}

func printLine(code string, status string, description string) {
	log.Printf("%s;%s;%s\n", code, status, description)
}

func printLintResult(result *zlint.ResultSet) {
	log.Default().SetFlags(0)
	log.Default().SetOutput(os.Stdout)
	printLine("Code", "Status", "Description")
	for i, v := range result.Results {
		switch v.Status {
		case lint.Error:
		case lint.Warn:
			{
				printLine(i, v.Status.String(), v.Details)
				break
			}
		}
	}
}

func doLint(certPath string) *zlint.ResultSet {
	raw, err := os.ReadFile(certPath)
	if err != nil {
		log.Fatalf("cannot read the certificate file %s, %s", certPath, err.Error())
	}

	// Parse cert
	p, _ := pem.Decode(raw)
	if len(p.Bytes) > 0 {
		// try parse pem and if it's ok replace raw to bytes from the pem
		if p.Type != "CERTIFICATE" {
			log.Fatalf("incorrect PEM data, tag value is %s, but should be 'CERTIFICATE'", p.Type)
		}
		raw = p.Bytes
	}
	cert, err := x509.ParseCertificate(raw)
	if err != nil {
		log.Fatalf("cannot parse the certificate %s, %s", certPath, err.Error())
	}

	// Initialize lint registry
	registry, err := lint.GlobalRegistry().Filter(lint.FilterOptions{
		IncludeSources: lint.SourceList{shaken.ShakenPolicy},
	})
	if err != nil {
		log.Fatalf("cannot initialize the lint registry, %s", err.Error())
	}

	return zlint.LintCertificateEx(cert, registry)
}
