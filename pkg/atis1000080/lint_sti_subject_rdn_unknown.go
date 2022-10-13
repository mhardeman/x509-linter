package atis1000080

import (
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/v3/lint"
)

const subjectRdn_details = "STI certificate shall not include RDNs that are not specified"

type subjectRdnUnknown struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_shaken_sti_subject_rdn_unknown",
		Description:   subjectRdn_details,
		Citation:      "Citation",
		Source:        SHAKEN,
		EffectiveDate: ATIS1000080_v004_Leaf_Date,
		Lint:          NewSubjectRdnUnknown,
	})
}

func NewSubjectRdnUnknown() lint.LintInterface {
	return &subjectRdnUnknown{}
}

// CheckApplies implements lint.LintInterface
func (*subjectRdnUnknown) CheckApplies(c *x509.Certificate) bool {
	return !c.IsCA
}

// Execute implements lint.LintInterface
func (*subjectRdnUnknown) Execute(c *x509.Certificate) *lint.LintResult {
	if err := assertNameUnknown(c.Subject.Names, []string{
		"2.5.4.3",  // commonName from ATIS
		"2.5.4.6",  // countryName from ATIS
		"2.5.4.10", // organization from ATIS
		"2.5.4.5",  // SERIALNUMBER from CP1.1
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

func assertNameUnknown(n []pkix.AttributeTypeAndValue, knownAttrs []string) error {
	list := newStringList(knownAttrs)
	for _, name := range n {
		if !list.Contains(name.Type.String()) {
			return fmt.Errorf(subjectRdn_details)
		}
	}

	return nil
}
