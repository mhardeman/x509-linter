package atis1000080

import (
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/v3/lint"
)

const subjectRdn_details = "Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network."

type subjectRdnUnknown struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_shaken_cp1_3_subject_rdn_unknown",
		Description:   subjectRdn_details,
		Citation:      CPv1_3_Citation,
		Source:        CPv1_3_Source,
		EffectiveDate: CPv1_3_Leaf_Date,
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
			return fmt.Errorf("No unknown RDNs are allowed as they may introduce ambiguity")
		}
	}

	return nil
}
