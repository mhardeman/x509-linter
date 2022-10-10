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

func GetOrganizationName(c *x509.Certificate) string {
	org := "Unknown"
	if len(c.Issuer.Organization) > 0 {
		org = c.Issuer.Organization[0]
	} else if len(c.Subject.Organization) > 0 {
		org = c.Subject.Organization[0]
	} else if len(c.Subject.CommonName) > 0 {
		org = c.Subject.CommonName
	}

	return org
}

// GetValidityDays returns validity period of the certificate in days
func GetValidityDays(c *x509.Certificate) int {
	return c.ValidityPeriod / 86400
}
