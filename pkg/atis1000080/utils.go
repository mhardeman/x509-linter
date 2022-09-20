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
