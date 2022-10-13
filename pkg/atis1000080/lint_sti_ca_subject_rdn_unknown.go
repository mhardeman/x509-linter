package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type caSubjectRdnUnknown struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_shaken_sti_ca_subject_rdn_unknown",
		Description:   subjectRdn_details,
		Citation:      "Citation",
		Source:        ATIS_Source,
		EffectiveDate: ATIS1000080_v004_Date,
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
	if err := assertNameUnknown(c.Subject.Names, []string{
		"2.5.4.3",  // commonName
		"2.5.4.6",  // countryName
		"2.5.4.10", // organization
	}); err != nil {
		return &lint.LintResult{
			Status:  lint.Warn,
			Details: err.Error(),
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
