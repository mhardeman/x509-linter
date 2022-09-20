package atis1000080

import (
	"time"

	"github.com/zmap/zlint/v3/lint"
)

var (
	SHAKEN                   = lint.LintSource("SHAKEN")
	ATIS1000080_v004_Date    = time.Date(2021, time.October, 5, 0, 0, 0, 0, time.UTC)
	ATIS1000080_STI_Citation = "ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements"
	CP_v1_3_Date             = time.Date(2021, time.August, 18, 0, 0, 0, 0, time.UTC)
	CP_v1_3_Citation         = "ToKENs (SHAKEN) Certificate Policy / "
)
