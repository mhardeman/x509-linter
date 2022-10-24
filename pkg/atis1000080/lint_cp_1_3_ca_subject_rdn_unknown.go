package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type caSubjectRdnUnknown struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_shaken_cp1_3_ca_subject_rdn_unknown",
		Description:   subjectRdn_details,
		Citation:      CPv1_3_Citation,
		Source:        CPv1_3_Source,
		EffectiveDate: CPv1_3_Date,
		Lint:          NewCaSubjectRdnUnknown,
	})
}

func NewCaSubjectRdnUnknown() lint.LintInterface {
	return &caSubjectRdnUnknown{}
}

// CheckApplies implements lint.LintInterface
func (*caSubjectRdnUnknown) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA
}

// Execute implements lint.LintInterface
func (*caSubjectRdnUnknown) Execute(c *x509.Certificate) *lint.LintResult {
	list := newStringList([]string{
		"2.5.4.3",  // commonName
		"2.5.4.6",  // countryName
		"2.5.4.10", // organization
	})
	for _, name := range c.Subject.Names {
		if !list.Contains(name.Type.String()) {
			return &lint.LintResult{
				Status:  lint.Warn,
				Details: "Only CN, C, and O can be included. Additional RNDs may introduce ambiguity and may not be verifiable",
			}
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
