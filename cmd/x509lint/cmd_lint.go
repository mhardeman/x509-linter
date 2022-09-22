package main

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

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

		counter := map[string]int{}
		allCerts := 0
		invalidCerts := 0

		for _, file := range files {
			if !file.IsDir() {
				filePath := path.Join(certPath, file.Name())
				result, err := doLint(filePath)
				if err != nil {
					continue
				}
				allCerts = allCerts + 1
				if result.ErrorsPresent || result.WarningsPresent || result.FatalsPresent {
					invalidCerts = invalidCerts + 1
					fmt.Println(filePath)
					printLintResult(result)
					for code, v := range result.Results {
						if v.Status != lint.Pass {
							counter[code] = counter[code] + 1
						}
					}
				}
			}
		}

		// Print common report
		log.Printf("%s;%s\n", "Name", "Amount")
		log.Printf("%s;%d\n", "Verified", allCerts)
		log.Printf("%s;%d\n", "Invalid", invalidCerts)
		if len(counter) > 0 {
			log.Printf("%s;%s\n", "Code", "Amount")
			for code, v := range counter {
				log.Printf("%s;%d\n", code, v)
			}
		}
	} else {
		// file is not a directory
		result, err := doLint(certPath)
		if err != nil {
			return fmt.Errorf("cannot lint the certificate, %w", err)
		}
		printLintResult(result)
	}

	return nil
}

func printLine(code string, status string, description string) {
	log.Printf("%s;%s;%s\n", code, status, description)
}

func printLintResult(result *zlint.ResultSet) {
	log.Default().SetFlags(0)
	log.Default().SetOutput(os.Stdout)
	printLine("Code", "Status", "Description")
	for i, v := range result.Results {
		if v.Status == lint.Error || v.Status == lint.Warn {
			printLine(i, v.Status.String(), v.Details)
		}
	}
}

func doLint(certPath string) (*zlint.ResultSet, error) {
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

	return zlint.LintCertificateEx(cert, registry), nil
}
