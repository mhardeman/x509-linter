package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type certificatePolicyCritical struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "n_shaken_sti_certificate_policy_critical",
		Description:   "STI certificates should contain a CertificatePolicies extension marked uncritical",
		Citation:      ATIS1000080_STI_Citation,
		Source:        ATIS_Source,
		EffectiveDate: ATIS1000080_v004_Leaf_Date,
		Lint:          NewCertificatePolicyCritical,
	})
}

func NewCertificatePolicyCritical() lint.LintInterface {
	return &certificatePolicyCritical{}
}

// CheckApplies implements lint.LintInterface
func (*certificatePolicyCritical) CheckApplies(c *x509.Certificate) bool {
	return !c.IsCA && FindExtension(c, id_CertificatePolicies) != nil
}

// Execute implements lint.LintInterface
func (*certificatePolicyCritical) Execute(c *x509.Certificate) *lint.LintResult {
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
