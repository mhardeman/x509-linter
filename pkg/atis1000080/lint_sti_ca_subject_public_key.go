package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type caSubjectPublicKey struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_ca_subject_public_key",
		Description:   subjectPublicKey_details,
		Citation:      ATIS1000080_STI_Citation,
		Source:        ATIS_Source,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewCaSubjectPublicKey,
	})
}

func NewCaSubjectPublicKey() lint.LintInterface {
	return &caSubjectPublicKey{}
}

// CheckApplies implements lint.LintInterface
func (*caSubjectPublicKey) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA
}

// Execute implements lint.LintInterface
func (*caSubjectPublicKey) Execute(c *x509.Certificate) *lint.LintResult {
	if err := assertSubjectPublicKey(c); err != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: err.Error(),
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
