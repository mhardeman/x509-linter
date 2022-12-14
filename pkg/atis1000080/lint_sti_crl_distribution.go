package atis1000080

import (
	"net/http"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type crlDistribution struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_crl_distribution",
		Description:   "STI intermediate and End-Entity certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry",
		Citation:      ATIS1000080_STI_Citation,
		Source:        SHAKEN,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewCRLDistribution,
	})
}

func NewCRLDistribution() lint.LintInterface {
	return &crlDistribution{}
}

// CheckApplies implements lint.LintInterface
func (*crlDistribution) CheckApplies(c *x509.Certificate) bool {
	return IsDateATIS1000080(c)
}

// Execute implements lint.LintInterface
func (*crlDistribution) Execute(c *x509.Certificate) *lint.LintResult {
	ext := FindExtension(c, "2.5.29.31")
	if ext != nil {
		if len(c.CRLDistributionPoints) == 1 {
			_, err := http.Get(c.CRLDistributionPoints[0])
			if err == nil {
				return &lint.LintResult{
					Status:  lint.Error,
					Details: "CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs",
				}
			}

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
