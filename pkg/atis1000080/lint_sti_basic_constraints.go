package atis1000080

import (
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type basicConstraints struct{}

type asnBasicConstraints struct {
	IsCA       bool `asn1:"optional"`
	MaxPathLen int  `asn1:"optional,default:-1"`
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_sti_basic_constraints",
		Description:   "STI certificates shall contain a BasicConstraints extension marked critical. For root and intermediate certificates, the BasicConstraints CA boolean shall be set to TRUE, while for End-Entity certificates, the CA boolean shall be set to FALSE.",
		Citation:      ATIS1000080_STI_Citation,
		Source:        ATIS1000080_Source,
		EffectiveDate: ATIS1000080_v004_Date,
		Lint:          NewBasicConstraints,
	})
}

func NewBasicConstraints() lint.LintInterface {
	return &basicConstraints{}
}

// CheckApplies implements lint.LintInterface
func (*basicConstraints) CheckApplies(c *x509.Certificate) bool {
	return true
}

// Execute implements lint.LintInterface
func (*basicConstraints) Execute(c *x509.Certificate) *lint.LintResult {
	ext := FindExtension(c, "2.5.29.19")

	if ext != nil && ext.Critical {
		basicConstraints := asnBasicConstraints{}
		if _, err := asn1.Unmarshal(ext.Value, &basicConstraints); err != nil {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: err.Error(), // "bad BasicConstraints ASN.1 value",
			}
		}

		return &lint.LintResult{
			Status: lint.Pass,
		}
	}

	return &lint.LintResult{
		Status:  lint.Error,
		Details: "STI certificates shall contain a BasicConstraints extension marked critical",
	}
}
