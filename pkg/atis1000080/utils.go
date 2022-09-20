package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
)

const id_TNAuthList = "1.3.6.1.5.5.7.1.26"

// FindTNAuthListExtension returns TNAuthList pkix.Extension from the certificate, otherwise returns nil
func FindTNAuthListExtension(c *x509.Certificate) *pkix.Extension {
	return FindExtension(c, id_TNAuthList)
}

func FindExtension(c *x509.Certificate, id string) *pkix.Extension {
	for _, v := range c.Extensions {
		if v.Id.String() == id {
			return &v
		}
	}

	return nil
}

type ShakenPolicy = string

const SHAKEN_CP_v1_1 ShakenPolicy = "2.16.840.1.114569.1.1.1"
const SHAKEN_CP_v1_3 ShakenPolicy = "2.16.840.1.114569.1.1.3"

func HasShakenPolicy(c *x509.Certificate, policy ShakenPolicy) bool {
	for _, v := range c.PolicyIdentifiers {
		identifier := v.String()
		if identifier == policy {
			return true
		}
	}

	return false
}
