package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type caCertificatePolicies struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_ca_certificate_policies",
		Description:   "STI Intermediate certificates shall include a Certificate Policies extension containing a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA",
		Citation:      ATIS1000080_STI_Citation,
		Source:        SHAKEN,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewCertificatePolicies,
	})
}

func NewCaCertificatePolicies() lint.LintInterface {
	return &caCertificatePolicies{}
}

// CheckApplies implements lint.LintInterface
func (*caCertificatePolicies) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA && !c.SelfSigned
}

// Execute implements lint.LintInterface
func (*caCertificatePolicies) Execute(c *x509.Certificate) *lint.LintResult {
	if len(c.PolicyIdentifiers) == 1 {
		if IsDateCP1_3(c) && c.PolicyIdentifiers[0].String() != SHAKEN_CP_v1_3 {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall contain '2.16.840.1.114569.1.1.3' policy",
			}
		}

		return &lint.LintResult{
			Status: lint.Pass,
		}
	}

	return &lint.LintResult{
		Status:  lint.Error,
		Details: "STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy",
	}
}
