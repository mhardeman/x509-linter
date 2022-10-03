package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type caKeyUsage struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_ca_key_usage",
		Description:   "STI certificates shall contain a Key Usage extension marked as critical. For root and intermediate certificates, the Key Usage extension shall contain the key usage value keyCertSign (5), and may contain the key usage values digitalSignature (0) and/or cRLSign (6)",
		Citation:      ATIS1000080_STI_Citation,
		Source:        SHAKEN,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewCaKeyUsage,
	})
}

func NewCaKeyUsage() lint.LintInterface {
	return &caKeyUsage{}
}

// CheckApplies implements lint.LintInterface
func (*caKeyUsage) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA
}

// Execute implements lint.LintInterface
func (*caKeyUsage) Execute(c *x509.Certificate) *lint.LintResult {
	ext := FindExtension(c, "2.5.29.15")
	if ext != nil && ext.Critical {
		flag := x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature
		if c.KeyUsage&x509.KeyUsageCertSign == x509.KeyUsageCertSign &&
			(c.KeyUsage|flag)^flag == 0 {
			return &lint.LintResult{
				Status: lint.Pass,
			}
		}

		return &lint.LintResult{
			Status:  lint.Error,
			Details: "The Key Usage extension shall contain the key usage value keyCertSign, and may contain the key usage values digitalSignature and/or cRLSign",
		}
	}

	return &lint.LintResult{
		Status:  lint.Error,
		Details: "STI certificates shall contain a Key Usage extension marked as critical",
	}
}
