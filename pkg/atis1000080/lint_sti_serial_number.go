package atis1000080

import (
	"math/big"
	"strings"
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type serialNumber struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_shaken_sti_serial_number",
		Description:   "STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG)",
		Citation:      "ATIS-1000080", // TODO check this name
		Source:        SHAKEN,
		EffectiveDate: time.Date(2022, time.September, 22, 0, 0, 0, 0, time.UTC), // TODO check date
		Lint:          NewSerialNumber,
	})
}

func NewSerialNumber() lint.LintInterface {
	return &serialNumber{}
}

// CheckApplies implements lint.LintInterface
func (*serialNumber) CheckApplies(c *x509.Certificate) bool {
	return IsDateATIS1000080(c)
}

// Execute implements lint.LintInterface
func (*serialNumber) Execute(c *x509.Certificate) *lint.LintResult {
	if strings.HasPrefix(c.SerialNumber.String(), "-") || c.SerialNumber.Cmp(big.NewInt(0x0100000000000000)) == -1 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG)",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
