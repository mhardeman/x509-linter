package shaken

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type serialNumber struct{}

const (
	serialNumberShallBeIncluded         = "The ‘serialNumber’ attribute shall be included along with the CN."
	serialNumberShallBeIncludedOnlyOnce = "The ‘serialNumber’ attribute shall be included along with the CN only time."
)

/************************************************
 Subsection 3.1.1 Types of names

The ‘serialNumber’ attribute shall be included along with the CN (to form a terminal relative distinguished name set), to distinguish among successive instances of certificates associated with the same entity.
 ************************************************/

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_types_of_names_serial_number",
		Description:   "The ‘serialNumber’ attribute shall be included along with the CN.",
		Citation:      "ToKENs (SHAKEN) Certificate Policy / Section 3.1.1",
		Source:        ShakenPolicy,
		EffectiveDate: ShakenDate_1_0_Date,
		Lint:          NewTypesOfNamesSerialNumberRequired,
	})
}

func NewTypesOfNamesSerialNumberRequired() lint.LintInterface {
	return &serialNumber{}
}

func (l *serialNumber) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *serialNumber) Execute(c *x509.Certificate) *lint.LintResult {
	values := make([]interface{}, 0)
	for _, v := range c.Subject.Names {
		if v.Type.String() == "2.5.4.5" {
			values = append(values, v.Value)
		}
	}

	if len(values) == 0 {
		return &lint.LintResult{
			Status:  lint.Warn,
			Details: serialNumberShallBeIncluded,
		}
	}

	if len(values) > 1 {
		return &lint.LintResult{
			Status:  lint.Warn,
			Details: serialNumberShallBeIncludedOnlyOnce,
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
