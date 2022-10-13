package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type rootCertificatePolicies struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_root_certificate_policies",
		Description:   "STI Root certificates shall not contain a Certificate Policies extension",
		Citation:      ATIS1000080_STI_Citation,
		Source:        SHAKEN,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewRootCertificatePolicies,
	})
}

func NewRootCertificatePolicies() lint.LintInterface {
	return &rootCertificatePolicies{}
}

// CheckApplies implements lint.LintInterface
func (*rootCertificatePolicies) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA && c.SelfSigned
}

// Execute implements lint.LintInterface
func (*rootCertificatePolicies) Execute(c *x509.Certificate) *lint.LintResult {
	if ext := FindExtension(c, id_CertificatePolicies); ext != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "STI Root certificates shall not contain a Certificate Policy extension",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
