package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type subjectPublicKeyIdentifier struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_sti_subject_public_key_identifier",
		Description:   "STI certificates shall contain a Subject Key Identifier extension.",
		Citation:      "ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements",
		Source:        ATIS1000080_Source,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewSubjectPublicKeyIdentifier,
	})
}

func NewSubjectPublicKeyIdentifier() lint.LintInterface {
	return &subjectPublicKeyIdentifier{}
}

// CheckApplies implements lint.LintInterface
func (*subjectPublicKeyIdentifier) CheckApplies(c *x509.Certificate) bool {
	return true
}

// Execute implements lint.LintInterface
func (*subjectPublicKeyIdentifier) Execute(c *x509.Certificate) *lint.LintResult {
	ext := FindExtension(c, "2.5.29.14")
	if ext != nil {
		return &lint.LintResult{
			Status: lint.Pass,
		}
	}

	return &lint.LintResult{
		Status:  lint.Error,
		Details: "STI certificates shall contain a Subject Key Identifier extension",
	}
}
