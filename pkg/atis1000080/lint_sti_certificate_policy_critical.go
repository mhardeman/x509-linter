package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type certificatePolicyCritical struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "n_sti_certificate_policy_critical",
		Description:   "STI certificates should contain a CertificatePolicies extension marked uncritical",
		Citation:      ATIS1000080_STI_Citation,
		Source:        SHAKEN,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewCertificatePolicyCritical,
	})
}

func NewCertificatePolicyCritical() lint.LintInterface {
	return &certificatePolicyCritical{}
}

// CheckApplies implements lint.LintInterface
func (*certificatePolicyCritical) CheckApplies(c *x509.Certificate) bool {
	return FindExtension(c, "2.5.29.32") != nil
}

// Execute implements lint.LintInterface
func (*certificatePolicyCritical) Execute(c *x509.Certificate) *lint.LintResult {
	certPoliciesExt := FindExtension(c, "2.5.29.32")
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
