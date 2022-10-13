package atis1000080

import (
	"time"

	"github.com/zmap/zlint/v3/lint"
)

var (
	SHAKEN                = lint.LintSource("SHAKEN")
	ATIS1000080_v004_Date = time.Date(2021, time.October, 18, 0, 0, 0, 0, time.UTC)
	// ATIS-1000080.v004 was published on Monday, 18 October 2021.
	// If we apply the same rule of 90 days after approval for CAs to be conformant with new versions as is applied to CP->CPS
	// changes this means any changes would be required for certificates issued on or after Sunday, January 16, 2022.
	ATIS1000080_v004_Leaf_Date = time.Date(2022, time.January, 16, 0, 0, 0, 0, time.UTC)
	ATIS1000080_STI_Citation   = "ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements"
	// CPS 1.3 was approved on August 18, 2021.
	// If we assume that CAs CPSs were submitted for approval on the 45th day, the CPS was approved within 10 days
	// and the CA has 90 days to become conformant then certificates issued on or after "Monday, January 10, 2022"
	// should be evaluated against the new rules.
	// January 10, 2022
	CP_v1_3_Date     = time.Date(2022, time.January, 10, 0, 0, 0, 0, time.UTC)
	CP_v1_3_Citation = "ToKENs (SHAKEN) Certificate Policy / "
)
