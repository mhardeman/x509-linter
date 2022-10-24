package atis1000080

import (
	"time"

	"github.com/zmap/zlint/v3/lint"
)

var (
	ATIS_Source   = lint.LintSource("ATIS1000080")
	CPv1_3_Source = lint.LintSource("CP 1.3")
	PKI_Source    = lint.LintSource("SHAKEN PKI Best Practice")

	ATIS1000080_v004_Date = time.Date(2021, time.October, 18, 0, 0, 0, 0, time.UTC)
	// ATIS-1000080.v004 was published on Monday, 18 October 2021.
	// If we apply the same rule of 90 days after approval for CAs to be conformant with new versions as is applied to CP->CPS
	// changes this means any changes would be required for certificates issued on or after Sunday,
	// March 12, 2022.
	ATIS1000080_v004_Leaf_Date = time.Date(2022, time.March, 12, 0, 0, 0, 0, time.UTC)
	ATIS1000080_STI_Citation   = "ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements"
	CPv1_3_Date                = time.Date(2021, time.August, 18, 0, 0, 0, 0, time.UTC)
	// CPS 1.3 was approved on August 18, 2021.
	// If we assume that CAs CPSs were submitted for approval on the 45th day, the CPS was approved within 10 days
	// and the CA has 90 days to become conformant then certificates issued on or after "Monday, January 10, 2022"
	// should be evaluated against the new rules.
	// June 12, 2022
	CPv1_3_Leaf_Date    = time.Date(2021, time.June, 12, 0, 0, 0, 0, time.UTC)
	CPv1_3_Citation     = "ToKENs (SHAKEN) Certificate Policy / "
	CPv1_3_Citation_4_9 = "ToKENs (SHAKEN) Certificate Policy / 4.9 Certificate Revocation and Suspension"
	PKI_Citation        = "SHAKEN PKI Best Practice"
)
