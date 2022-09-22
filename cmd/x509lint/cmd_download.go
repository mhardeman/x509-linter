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

	"github.com/zmap/zcrypto/x509"
)

func RunDownloadCommand(listPath string, outDir string) error {
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
	for _, link := range links {
		// use http links only
		if !strings.HasPrefix(link, "http") {
			continue
		}

		// send GET request
		resp, err := http.Get(link)
		if err != nil {
			fmt.Printf("%-7s%s %s\n", "ERROR", link, err.Error())
			continue
		}

		// check status for GET response
		if resp.StatusCode == 200 {
			defer resp.Body.Close()

			certRaw, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("%-7s%s %s\n", "ERROR", link, err.Error())
				continue
			}

			if pem, err := pem.Decode(certRaw); err != nil {
				certRaw = pem.Bytes
			}
			cert, err := x509.ParseCertificate(certRaw)
			if err != nil {
				fmt.Printf("%-7s%s %s\n", "ERROR", link, err.Error())
				continue
			}
			certBlock := pem.Block{
				Type: "CERTIFICATE",
				Headers: map[string]string{
					"Link":    link,
					"Subject": cert.Subject.String(),
					"Issuer":  cert.Issuer.String(),
				},
				Bytes: certRaw,
			}
			pemCert := pem.EncodeToMemory(&certBlock)
			sha1 := crypto.SHA1.New()
			sha1.Write(certRaw)
			fileName := fmt.Sprintf("%s.pem", hex.EncodeToString(sha1.Sum(nil)))
			err = os.WriteFile(path.Join(outDir, fileName), pemCert, 0644)
			if err != nil {
				fmt.Printf("%-7s%s %s\n", "ERROR", link, err.Error())
				continue
			}

			fmt.Printf("%-7s%s\n", "OK", link)
		} else {
			fmt.Printf("%-7s%s error on HTTP request %d (%s)\n", "ERROR", link, resp.StatusCode, resp.Status)
		}
	}

	return nil
}
