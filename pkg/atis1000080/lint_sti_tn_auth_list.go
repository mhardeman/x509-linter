package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type tnAuthList struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_sti_tn_auth_list",
		Description:   "STI End-Entity certificates shall contain a TNAuthList extension as specified in RFC 8226. The TNAuthList shall contain a single SPC value",
		Citation:      ATIS1000080_STI_Citation,
		Source:        ATIS1000080_Source,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewTnAuthList,
	})
}

func NewTnAuthList() lint.LintInterface {
	return &tnAuthList{}
}

// CheckApplies implements lint.LintInterface
func (*tnAuthList) CheckApplies(c *x509.Certificate) bool {
	return true
}

// Execute implements lint.LintInterface
func (*tnAuthList) Execute(c *x509.Certificate) *lint.LintResult {
	_, err := GetTNEntrySPC(c)
	if err != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: err.Error(),
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
