package atis1000080

import (
	"fmt"
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/v3/lint"
)

const id_TNAuthList = "1.3.6.1.5.5.7.1.26"
const id_AuthorityKeyIdentifier = "2.5.29.35"

// FindTNAuthListExtension returns TNAuthList pkix.Extension from the certificate, otherwise returns nil
func FindTNAuthListExtension(c *x509.Certificate) *pkix.Extension {
	return FindExtension(c, id_TNAuthList)
}

func FindExtension(c *x509.Certificate, id string) *pkix.Extension {
	for _, v := range c.Extensions {
		if v.Id.String() == id {
			return &v
		}
	}

	return nil
}

type ShakenPolicy = string

const SHAKEN_CP_v1_1 ShakenPolicy = "2.16.840.1.114569.1.1.1"
const SHAKEN_CP_v1_3 ShakenPolicy = "2.16.840.1.114569.1.1.3"

func HasShakenPolicy(c *x509.Certificate, policy ShakenPolicy) bool {
	for _, v := range c.PolicyIdentifiers {
		identifier := v.String()
		if identifier == policy {
			return true
		}
	}

	return false
}

func IsDateATIS1000080(c *x509.Certificate) bool {
	// ATIS-1000080.v004 was published on Monday, 18 October 2021.

	// If we apply the same rule of 90 days after approval for CAs to be conformant with new versions as is applied to CP->CPS
	// changes this means any changes would be required for certificates issued on or after Sunday, January 16, 2022.

	return !c.NotBefore.Before(time.Date(2022, time.January, 16, 0, 0, 0, 0, time.UTC))
}

func IsDateCP1_3(c *x509.Certificate) bool {
	// CPS 1.3 was approved on August 18, 2021.

	// If we assume that CAs CPSs were submitted for approval on the 45th day, the CPS was approved within 10 days
	// and the CA has 90 days to become conformant then certificates issued on or after "Monday, January 10, 2022"
	// should be evaluated against the new rules.

	// January 10, 2022
	return !c.NotBefore.Before(time.Date(2022, time.January, 10, 0, 0, 0, 0, time.UTC))
}

const downgrade_ATIS1000080 = "This certificate was issued before the latest ATIS-1000080 specification. The test result has been downgraded to notice as this test may not have existed in earlier versions of the specification."
const downgrade_cp1_3 = "This certificate was issued before the latest Certificate Policy. The test result has been downgraded to notice as this test may not have existed in earlier versions of the policy."

// DowngradeATIS1000080 Downgrades lint result for old certs
func DowngradeATIS1000080(c *x509.Certificate, r *lint.LintResult) *lint.LintResult {
	if !IsDateATIS1000080(c) && r.Status != lint.Pass {
		r.Status = lint.Notice
		r.Details = fmt.Sprintf("%s %s", downgrade_ATIS1000080, r.Details)
	}

	return r
}

// DowngradeCP1_3 Downgrades lint result for old certs
func DowngradeCP1_3(c *x509.Certificate, r *lint.LintResult) *lint.LintResult {
	if !IsDateCP1_3(c) && r.Status != lint.Pass {
		r.Status = lint.Notice
		r.Details = fmt.Sprintf("%s %s", downgrade_cp1_3, r.Details)
	}

	return r
}
