package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type crlDistribution struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_sti_crl_distribution",
		Description:   "STI intermediate and End-Entity certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry",
		Citation:      "ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements",
		Source:        ATIS1000080_Source,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewCRLDistribution,
	})
}

func NewCRLDistribution() lint.LintInterface {
	return &crlDistribution{}
}

// CheckApplies implements lint.LintInterface
func (*crlDistribution) CheckApplies(c *x509.Certificate) bool {
	return true
}

// Execute implements lint.LintInterface
func (*crlDistribution) Execute(c *x509.Certificate) *lint.LintResult {
	ext := FindExtension(c, "2.5.29.31")
	if ext != nil {
		if len(c.CRLDistributionPoints) == 1 {
			return &lint.LintResult{
				Status: lint.Pass,
			}
		}

		return &lint.LintResult{
			Status:  lint.Error,
			Details: "CRL Distribution Points extension should contain a single DistributionPoint entry",
		}
	}

	return &lint.LintResult{
		Status:  lint.Error,
		Details: "STI intermediate and End-Entity certificates shall contain a CRL Distribution Points extension",
	}
}
