package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type subjectKeyIdentifier struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_sti_subject_key_identifier",
		Description:   "STI certificates shall contain a Subject Key Identifier extension",
		Citation:      ATIS1000080_STI_Citation,
		Source:        SHAKEN,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewSubjectKeyIdentifier,
	})
}

func NewSubjectKeyIdentifier() lint.LintInterface {
	return &subjectKeyIdentifier{}
}

// CheckApplies implements lint.LintInterface
func (*subjectKeyIdentifier) CheckApplies(c *x509.Certificate) bool {
	return true
}

// Execute implements lint.LintInterface
func (*subjectKeyIdentifier) Execute(c *x509.Certificate) *lint.LintResult {
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
