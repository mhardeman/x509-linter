package atis1000080

import (
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type ambiguousIdentifiers struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_cp1_3_ambiguous_identifier",
		Description:   "Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject",
		Citation:      CPv1_3_Citation,
		Source:        CPv1_3_Source,
		EffectiveDate: CPv1_3_Leaf_Date,
		Lint:          NewAmbiguousIdentifiers,
	})
}

func NewAmbiguousIdentifiers() lint.LintInterface {
	return &ambiguousIdentifiers{}
}

// CheckApplies implements lint.LintInterface
func (*ambiguousIdentifiers) CheckApplies(c *x509.Certificate) bool {
	return !c.IsCA
}

// Execute implements lint.LintInterface
func (*ambiguousIdentifiers) Execute(c *x509.Certificate) *lint.LintResult {
	// For linters to be effective tools they must take an opinionated interpretation of the associated requirements.
	// Both the CP and ATIS-1000080 have language that allow for extra values to be included in the CN and subject name
	// however the CP also has language that states names must be unambiguous. The only way to test for this lack of ambiguity
	// is to state that the name must follow a specific form.

	spc, err := GetTNEntrySPC(c)
	if err != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: err.Error(),
		}
	}

	name := fmt.Sprintf("SHAKEN %s", spc)
	if c.Subject.CommonName != name {
		return DowngradeCP1_3(c, &lint.LintResult{
			Status:  lint.Error,
			Details: "Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject",
		})
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
