package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type signatureAlgorithm struct{}

var signatureAlgorithm_details = "STI certificates shall contain a Signature Algorithm field with the value 'ecdsa-with-SHA256'"

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_sti_signature_algorithm",
		Description:   signatureAlgorithm_details,
		Citation:      "ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements",
		Source:        ATIS1000080_Source,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewSignatureAlgorithm,
	})
}

func NewSignatureAlgorithm() lint.LintInterface {
	return &signatureAlgorithm{}
}

// CheckApplies implements lint.LintInterface
func (*signatureAlgorithm) CheckApplies(c *x509.Certificate) bool {
	return true
}

// Execute implements lint.LintInterface
func (*signatureAlgorithm) Execute(c *x509.Certificate) *lint.LintResult {
	if c.SignatureAlgorithmOID.String() != "1.2.840.10045.4.3.2" {
		return &lint.LintResult{
			Status:  lint.Warn,
			Details: signatureAlgorithm_details,
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
