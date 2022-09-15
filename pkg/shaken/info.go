package shaken

import (
	"time"

	"github.com/zmap/zlint/v3/lint"
)

var (
	ShakenPolicy        = lint.LintSource("SHAKEN")
	ShakenDate_1_0_Date = time.Date(2019, time.October, 22, 0, 0, 0, 0, time.UTC)
)
