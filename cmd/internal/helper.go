package internal

import (
	"crypto/x509"
	"encoding/pem"
)

// ParseCertificates parses certificates from the DER or PEM source
func ParseCertificates(source []byte) []*x509.Certificate {
	res := []*x509.Certificate{}

	// PEM
	// pem may contain multiple certificates
	for {
		pem, restBytes := pem.Decode(source)
		if pem == nil {
			break
		}
		cert, _ := x509.ParseCertificate(pem.Bytes)
		if cert != nil {
			res = append(res, cert)
		}
		source = restBytes
	}

	// DER
	cert, _ := x509.ParseCertificate(source)
	if cert != nil {
		res = append(res, cert)
	}

	return res
}
