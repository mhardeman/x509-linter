package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type certificatePolicies struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_sti_certificate_policies",
		Description:   "STI intermediate and End-Entity certificates shall include a Certificate Policies extension containing a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA",
		Citation:      ATIS1000080_STI_Citation,
		Source:        SHAKEN,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewCertificatePolicies,
	})
}

func NewCertificatePolicies() lint.LintInterface {
	return &certificatePolicies{}
}

// CheckApplies implements lint.LintInterface
func (*certificatePolicies) CheckApplies(c *x509.Certificate) bool {
	return true
}

// Execute implements lint.LintInterface
func (*certificatePolicies) Execute(c *x509.Certificate) *lint.LintResult {
	if len(c.PolicyIdentifiers) == 1 && c.PolicyIdentifiers[0].String() == "2.16.840.1.114569.1.1.3" {
		return &lint.LintResult{
			Status: lint.Pass,
		}
	}

	return &lint.LintResult{
		Status:  lint.Error,
		Details: "STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy",
	}
}
