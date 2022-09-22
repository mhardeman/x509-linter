package atis1000080

import (
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type subjectRdnUnknown struct{}

var shakenRdn = []string{
	"2.5.4.3",  // commonName
	"2.5.4.6",  // countryName
	"2.5.4.10", // organization
	"2.5.4.5",  // SERIALNUMBER
}

func isShakenRdn(id string) bool {
	for _, rdn := range shakenRdn {
		if id == rdn {
			return true
		}
	}
	return false
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_subject_rdn_unknown",
		Description:   "STI certificate shall not include RDNs that are not specified",
		Citation:      "Citation",
		Source:        SHAKEN,
		EffectiveDate: time.Date(2019, time.October, 22, 0, 0, 0, 0, time.UTC),
		Lint:          NewSubjectRdnUnknown,
	})
}

func NewSubjectRdnUnknown() lint.LintInterface {
	return &subjectRdnUnknown{}
}

// CheckApplies implements lint.LintInterface
func (*subjectRdnUnknown) CheckApplies(c *x509.Certificate) bool {
	return true
}

// Execute implements lint.LintInterface
func (*subjectRdnUnknown) Execute(c *x509.Certificate) *lint.LintResult {
	unknownRdn := []string{}
	for _, name := range c.Subject.Names {
		nameType := name.Type.String()
		if isShakenRdn(nameType) {
			continue
		}

		unknownRdn = append(unknownRdn, nameType)
	}

	if len(unknownRdn) != 0 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "STI certificate shall not include RDNs that are not specified",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
