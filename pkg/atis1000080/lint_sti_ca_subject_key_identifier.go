package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type caSubjectKeyIdentifier struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_ca_subject_key_identifier",
		Description:   subjectKeyIdentifier_details,
		Citation:      ATIS1000080_STI_Citation,
		Source:        ATIS_Source,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewCaSubjectKeyIdentifier,
	})
}

func NewCaSubjectKeyIdentifier() lint.LintInterface {
	return &caSubjectKeyIdentifier{}
}

// CheckApplies implements lint.LintInterface
func (*caSubjectKeyIdentifier) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA
}

// Execute implements lint.LintInterface
func (*caSubjectKeyIdentifier) Execute(c *x509.Certificate) *lint.LintResult {
	if err := assertSubjectKeyIdentifier(c); err != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: err.Error(),
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
