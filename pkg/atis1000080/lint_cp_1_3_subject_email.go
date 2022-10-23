package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type subjectEmail struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_shaken_cp_1_3_subject_email",
		Description:   "Information that is not verified shall not be included in certificates",
		Citation:      CPv1_3_Citation,
		Source:        CPv1_3_Source,
		EffectiveDate: CPv1_3_Leaf_Date,
		Lint:          NewSubjectEmail,
	})
}

func NewSubjectEmail() lint.LintInterface {
	return &subjectEmail{}
}

// CheckApplies implements lint.LintInterface
func (*subjectEmail) CheckApplies(c *x509.Certificate) bool {
	return !c.IsCA
}

// Execute implements lint.LintInterface
func (*subjectEmail) Execute(c *x509.Certificate) *lint.LintResult {
	list := newStringList(nil)

	for _, name := range c.Subject.Names {
		list.Items = append(list.Items, name.Type.String())
	}

	if list.Contains("1.2.840.113549.1.9.1") {
		return &lint.LintResult{
			Status:  lint.Warn,
			Details: "Email addresses are not allowed as the CP does not specify how to validate them",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
