package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

/*
 STI intermediate and End-Entity certificates shall contain an Authority Key Identifier extension (this extension is
 optional for root certificates). For root certificates that contain an Authority Key Identifier extension, the Authority
 Key Identifier shall contain a keyIdentifier field with a value that matches the Subject Key Identifier value of the
 same root certificate. For intermediate and End-Entity certificates, the Authority Key Identifier extension shall
 contain a keyIdentifier field with a value that matches the Subject Key Identifier value of the parent certificate.
*/

type authorityKeyIdentifier struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_authority_key_identifier",
		Description:   "STI certificates shall contain an Authority Key Identifier extension",
		Citation:      ATIS1000080_STI_Citation,
		Source:        ATIS_Source,
		EffectiveDate: ATIS1000080_v004_Leaf_Date,
		Lint:          NewAuthorityKeyIdentifier,
	})
}

func NewAuthorityKeyIdentifier() lint.LintInterface {
	return &authorityKeyIdentifier{}
}

// CheckApplies implements lint.LintInterface
func (*authorityKeyIdentifier) CheckApplies(c *x509.Certificate) bool {
	return !c.IsCA
}

// Execute implements lint.LintInterface
func (*authorityKeyIdentifier) Execute(c *x509.Certificate) *lint.LintResult {
	if ext := FindExtension(c, id_AuthorityKeyIdentifier); ext == nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "STI certificates shall contain an Authority Key Identifier extension",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
