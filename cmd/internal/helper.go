package internal

import (
	"encoding/pem"

	"github.com/zmap/zcrypto/x509"
)

type PemCertificate struct {
	Headers     map[string]string
	Certificate *x509.Certificate
}

// ParseCertificates parses certificates from the DER or PEM source
func ParseCertificates(source []byte) []*PemCertificate {
	res := []*PemCertificate{}

	// PEM
	// pem may contain multiple certificates
	for {
		pem, restBytes := pem.Decode(source)
		if pem == nil {
			break
		}
		cert, _ := x509.ParseCertificate(pem.Bytes)
		if cert != nil {
			res = append(res, &PemCertificate{
				Headers:     pem.Headers,
				Certificate: cert,
			})
		}
		source = restBytes
	}

	// DER
	cert, _ := x509.ParseCertificate(source)
	if cert != nil {
		res = append(res, &PemCertificate{
			Headers:     map[string]string{},
			Certificate: cert,
		})
	}

	return res
}
