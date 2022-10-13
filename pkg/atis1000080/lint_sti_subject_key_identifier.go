package atis1000080

import (
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

const id_SubjectKeyIdentifier = "2.5.29.14"
const subjectKeyIdentifier_details = "STI certificates shall contain a Subject Key Identifier extension"

type subjectKeyIdentifier struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_subject_key_identifier",
		Description:   subjectKeyIdentifier_details,
		Citation:      ATIS1000080_STI_Citation,
		Source:        ATIS_Source,
		EffectiveDate: ATIS1000080_v004_Leaf_Date,
		Lint:          NewSubjectKeyIdentifier,
	})
}

func NewSubjectKeyIdentifier() lint.LintInterface {
	return &subjectKeyIdentifier{}
}

// CheckApplies implements lint.LintInterface
func (*subjectKeyIdentifier) CheckApplies(c *x509.Certificate) bool {
	return !c.IsCA
}

// Execute implements lint.LintInterface
func (*subjectKeyIdentifier) Execute(c *x509.Certificate) *lint.LintResult {
	if err := assertSubjectKeyIdentifier(c); err != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: err.Error(),
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}

func assertSubjectKeyIdentifier(c *x509.Certificate) error {
	ext := FindExtension(c, id_SubjectKeyIdentifier)
	if ext == nil {
		return fmt.Errorf(subjectKeyIdentifier_details)
	}

	return nil
}
