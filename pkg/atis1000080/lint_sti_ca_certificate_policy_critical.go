package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type caCertificatePolicyCritical struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "n_shaken_sti_ca_certificate_policy_critical",
		Description:   "STI certificates should contain a CertificatePolicies extension marked uncritical",
		Citation:      ATIS1000080_STI_Citation,
		Source:        ATIS_Source,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewCaCertificatePolicyCritical,
	})
}

func NewCaCertificatePolicyCritical() lint.LintInterface {
	return &caCertificatePolicyCritical{}
}

// CheckApplies implements lint.LintInterface
func (*caCertificatePolicyCritical) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA && !c.SelfSigned && FindExtension(c, id_CertificatePolicies) != nil
}

// Execute implements lint.LintInterface
func (*caCertificatePolicyCritical) Execute(c *x509.Certificate) *lint.LintResult {
	certPoliciesExt := FindExtension(c, id_CertificatePolicies)
	if certPoliciesExt != nil && certPoliciesExt.Critical {
		return &lint.LintResult{
			Status:  lint.Notice,
			Details: "STI certificates should contain a CertificatePolicies extension marked uncritical",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
