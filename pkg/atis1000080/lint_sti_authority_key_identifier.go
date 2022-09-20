package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type authorityKeyIdentifier struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_sti_authority_key_identifier",
		Description:   "STI certificates shall contain an Authority Key Identifier extension",
		Citation:      ATIS1000080_STI_Citation,
		Source:        SHAKEN,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewAuthorityKeyIdentifier,
	})
}

func NewAuthorityKeyIdentifier() lint.LintInterface {
	return &authorityKeyIdentifier{}
}

// CheckApplies implements lint.LintInterface
func (*authorityKeyIdentifier) CheckApplies(c *x509.Certificate) bool {
	return true
}

// Execute implements lint.LintInterface
func (*authorityKeyIdentifier) Execute(c *x509.Certificate) *lint.LintResult {
	ext := FindExtension(c, "2.5.29.35")
	if ext != nil {
		return &lint.LintResult{
			Status: lint.Pass,
		}
	}

	return &lint.LintResult{
		Status:  lint.Error,
		Details: "STI certificates shall contain an Authority Key Identifier extension",
	}
}
