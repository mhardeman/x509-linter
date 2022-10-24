package atis1000080

import (
	"github.com/zmap/zcrypto/x509"
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
	list := newStringList([]string{
		"2.5.4.3",  // commonName from ATIS
		"2.5.4.6",  // countryName from ATIS
		"2.5.4.10", // organization from ATIS
		"2.5.4.5",  // SERIALNUMBER from CP1.1
	})
	for _, name := range c.Subject.Names {
		if !list.Contains(name.Type.String()) {
			return &lint.LintResult{
				Status:  lint.Warn,
				Details: "Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable",
			}
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
