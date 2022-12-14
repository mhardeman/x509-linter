package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type version struct{}

var version_details = "STI certificates shall contain Version field specifying version 3"

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_version",
		Description:   version_details,
		Citation:      ATIS1000080_STI_Citation,
		Source:        SHAKEN,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewVersion,
	})
}

func NewVersion() lint.LintInterface {
	return &version{}
}

// CheckApplies implements lint.LintInterface
func (*version) CheckApplies(c *x509.Certificate) bool {
	return IsDateATIS1000080(c)
}

// Execute implements lint.LintInterface
func (*version) Execute(c *x509.Certificate) *lint.LintResult {
	if c.Version != 3 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: version_details,
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
