package shaken

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

const organizationIdentityCountryName = "the certificate shall contain the 'countryName' field"

type organizationIdentity struct{}

/************************************************
Subsection 3.2.2 Authentication of organization identity

The certificate shall contain the ‘countryName’ field and other Subject Identity Information. The CA shall verify the
identity of the SP and the authenticity of the SP Applicant Representative’s certificate request using a verification
process meeting the requirements of this and that is described in the CA’s CPS. Specifically, the CA shall ensure
that the SP has a valid SPC token.
************************************************/

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_organization_identity",
		Description:   "The certificate shall contain the ‘countryName’ field and other Subject Identity Information",
		Citation:      "ToKENs (SHAKEN) Certificate Policy / Section 3.2.2",
		Source:        ShakenPolicy,
		EffectiveDate: ShakenDate_1_0_Date,
		Lint:          NewOrganizationIdentity,
	})
}

func NewOrganizationIdentity() lint.LintInterface {
	return &organizationIdentity{}
}

func (l *organizationIdentity) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *organizationIdentity) Execute(c *x509.Certificate) *lint.LintResult {
	if len(c.Subject.Country) == 0 {
		return &lint.LintResult{
			Status:  lint.Warn,
			Details: organizationIdentityCountryName,
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
