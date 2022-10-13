package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type caVersion struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_ca_version",
		Description:   version_details,
		Citation:      ATIS1000080_STI_Citation,
		Source:        ATIS_Source,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewCaVersion,
	})
}

func NewCaVersion() lint.LintInterface {
	return &caVersion{}
}

// CheckApplies implements lint.LintInterface
func (*caVersion) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA
}

// Execute implements lint.LintInterface
func (*caVersion) Execute(c *x509.Certificate) *lint.LintResult {
	if err := assertVersion(c); err != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: err.Error(),
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
