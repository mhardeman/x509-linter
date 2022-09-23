package atis1000080

import (
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type extensionUnknown struct{}

var shakenExtensions = []string{
	"2.5.29.14",          // subject key identifier
	"2.5.29.35",          // authority key identifier
	"2.5.29.19",          // basicConstraints
	"2.5.29.15",          // keyUsage
	"1.3.6.1.5.5.7.1.26", // tnAuthList
	"2.5.29.31",          // cRLDistributionPoints
	"2.5.29.32",          // certificatePolicies
}

func isShakenExtension(id string) bool {
	for _, extensionId := range shakenExtensions {
		if id == extensionId {
			return true
		}
	}
	return false
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_extension_unknown",
		Description:   "STI certificate shall not include extensions that are not specified",
		Citation:      "Citation",
		Source:        SHAKEN,
		EffectiveDate: time.Date(2019, time.October, 22, 0, 0, 0, 0, time.UTC),
		Lint:          NewExtensionUnknown,
	})
}

func NewExtensionUnknown() lint.LintInterface {
	return &extensionUnknown{}
}

// CheckApplies implements lint.LintInterface
func (*extensionUnknown) CheckApplies(c *x509.Certificate) bool {
	return IsDateATIS1000080(c)
}

// Execute implements lint.LintInterface
func (*extensionUnknown) Execute(c *x509.Certificate) *lint.LintResult {
	unknownExtensions := []string{}
	for _, extension := range c.Extensions {
		if isShakenExtension(extension.Id.String()) {
			continue
		}

		unknownExtensions = append(unknownExtensions, extension.Id.String())
	}

	if len(unknownExtensions) != 0 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "STI certificate shall not include extensions that are not specified",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
