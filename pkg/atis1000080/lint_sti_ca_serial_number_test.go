package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_caSerialNumber_CheckApplies(t *testing.T) {
	CheckAppliesRootOrIntermediateCertificate(t, "caSerialNumber", atis1000080.NewCaSerialNumber)
}

var CERT_SN_NEGATIVE = ParseCert("MIH3MIGeoAMCAQICCIAAAAAAAAABMAoGCCqGSM49BAMCMAAwHhcNMjIwOTIyMTMwNjEwWhcNMjIwOTIzMTMwNjEwWjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHggTCbcixmrHtnAQyawuXvDnbn8ucB2/Osv0DFYcaJYIK4jA0HJY9+MGFxWtPKsk5bZdw0vEfEaaLME4PmImnaMCMAAwCgYIKoZIzj0EAwIDSAAwRQIgCopMOKXD7w4Wq9JPIAiESNKnHhj+lBjvjLpvg6jd7GACIQCvdcdnPFHMOPM9qLav5x7frLYcg00tITa7Gdf5Mq1LrA==")
var CERT_SN_INCORRECT_FIRST_BYTE = ParseCert("MIH3MIGeoAMCAQICCACAAAAAAAABMAoGCCqGSM49BAMCMAAwHhcNMjIwOTIyMTMxMDIxWhcNMjIwOTIzMTMxMDIxWjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE//P/CaLZm8dXmFemUOb5XOHPAlIvAPq8JuRN69i+xHa1JIMIq3fiJ9QJfIjTX3ub3ZSXc/hIp7gacU02Ab0MlaMCMAAwCgYIKoZIzj0EAwIDSAAwRQIhAPCZsDcoCQEgv31CjOIKUJkZYq5wsHL95tVcK+7xg7dUAiB8MbfqSKVkDhPZnWzy0ZSPRa8Jzxc0AXM0MSWi1xNYTg==")
var CERT_SN_LESS_64BITS = ParseCert("MIH3MIGeoAMCAQICCACAAAAAAAABMAoGCCqGSM49BAMCMAAwHhcNMjIwOTIyMTMxMzI1WhcNMjIwOTIzMTMxMzI1WjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEohe1I0oCMJKZh6IXxNQn0jf7UTxRPLzwC+OPVYo9ZNi+hnf7JCu1TfJqyrUZcCpG6QFTJ79JToW+IloaZIqnfqMCMAAwCgYIKoZIzj0EAwIDSAAwRQIgSbOPYLTSTMln8Xu7zGAUGz66kxwLDFqjpSOI83BFd+8CIQD368TanGFzBNZPiaip1Y79vqc99Fv3bnNCTExc4vQDzQ==")

func Test_caSerialNumber_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "SN is negative", // -9223372036854776000
			args: args{
				c: CERT_SN_NEGATIVE,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG)",
			},
		},
		{
			name: "SN is 64-bit with 0 first octet", // ASN(hex): 02 08 0080000000000001
			args: args{
				c: CERT_SN_INCORRECT_FIRST_BYTE,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG)",
			},
		},
		{
			name: "SN is less than 64-bit", // ASN(hex): 02 07 70000000000001
			args: args{
				c: CERT_SN_LESS_64BITS,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG)",
			},
		},
		{
			name: "SN is correct", // ASN(hex): 02 08 0100000000000000
			args: args{
				c: TEST_CERT_CORRECT,
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := atis1000080.NewSerialNumber()
			if got := s.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serialNumber.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
