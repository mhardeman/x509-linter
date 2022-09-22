package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_serialNumber_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		// NOTE golang throws error on asn parsing
		//   cannot parse the certificate, error: asn1: structure error: integer not minimally-encoded [recovered]
		//
		// {
		// 	name: "SN 0x0000000000000001",
		// 	args: args{
		// 		c: ParseCert("MIH2MIGeoAMCAQICCAAAAAAAAAABMAoGCCqGSM49BAMCMAAwHhcNMjIwOTIyMTI1OTU0WhcNMjIwOTIzMTI1OTU0WjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEf1jf13iTb69+gQEa+uV0kwxKN1PKyWVyQ9Qr7Ib+/a8r8LsZ/OjV9CC5fNohxFjpU7UDsBNI8OvioMQD1/3uvqMCMAAwCgYIKoZIzj0EAwIDRwAwRAIgMv6rkyqtCHH2JF4wUvT/vyNRWLo9y25Cluhh+q/hA9gCICDHbT8ra17g+b//NUXdlbkBnSlaBaKro95Phm9ty1ce"),
		// 	},
		// 	want: &lint.LintResult{
		// 		Status: lint.Error,
		// 	},
		// },
		{
			name: "SN is negative", // -9223372036854776000
			args: args{
				c: ParseCert("MIH3MIGeoAMCAQICCIAAAAAAAAABMAoGCCqGSM49BAMCMAAwHhcNMjIwOTIyMTMwNjEwWhcNMjIwOTIzMTMwNjEwWjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHggTCbcixmrHtnAQyawuXvDnbn8ucB2/Osv0DFYcaJYIK4jA0HJY9+MGFxWtPKsk5bZdw0vEfEaaLME4PmImnaMCMAAwCgYIKoZIzj0EAwIDSAAwRQIgCopMOKXD7w4Wq9JPIAiESNKnHhj+lBjvjLpvg6jd7GACIQCvdcdnPFHMOPM9qLav5x7frLYcg00tITa7Gdf5Mq1LrA=="),
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG)",
			},
		},
		{
			name: "SN is 64-bit with 0 first octet", // ASN(hex): 02 08 0080000000000001
			args: args{
				c: ParseCert("MIH3MIGeoAMCAQICCACAAAAAAAABMAoGCCqGSM49BAMCMAAwHhcNMjIwOTIyMTMxMDIxWhcNMjIwOTIzMTMxMDIxWjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE//P/CaLZm8dXmFemUOb5XOHPAlIvAPq8JuRN69i+xHa1JIMIq3fiJ9QJfIjTX3ub3ZSXc/hIp7gacU02Ab0MlaMCMAAwCgYIKoZIzj0EAwIDSAAwRQIhAPCZsDcoCQEgv31CjOIKUJkZYq5wsHL95tVcK+7xg7dUAiB8MbfqSKVkDhPZnWzy0ZSPRa8Jzxc0AXM0MSWi1xNYTg=="),
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG)",
			},
		},
		{
			name: "SN is less than 64-bit", // ASN(hex): 02 07 70000000000001
			args: args{
				c: ParseCert("MIH3MIGeoAMCAQICCACAAAAAAAABMAoGCCqGSM49BAMCMAAwHhcNMjIwOTIyMTMxMzI1WhcNMjIwOTIzMTMxMzI1WjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEohe1I0oCMJKZh6IXxNQn0jf7UTxRPLzwC+OPVYo9ZNi+hnf7JCu1TfJqyrUZcCpG6QFTJ79JToW+IloaZIqnfqMCMAAwCgYIKoZIzj0EAwIDSAAwRQIgSbOPYLTSTMln8Xu7zGAUGz66kxwLDFqjpSOI83BFd+8CIQD368TanGFzBNZPiaip1Y79vqc99Fv3bnNCTExc4vQDzQ=="),
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG)",
			},
		},
		{
			name: "SN is correct", // ASN(hex): 02 08 0100000000000000
			args: args{
				c: ParseCert("MIH3MIGeoAMCAQICCAEAAAAAAAAAMAoGCCqGSM49BAMCMAAwHhcNMjIwOTIyMTMxOTU5WhcNMjIwOTIzMTMxOTU5WjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8yykAVGm9rSIrpKxfJdrn2C/OL4AA9uXAhfqGrSIxFOZstPglBTtBN0dS/hEAegr9c6B2V+YMg/0ZZdvYChCYKMCMAAwCgYIKoZIzj0EAwIDSAAwRQIgfpk6JLO/wZTicpsgLZM/dw5uFD87XtT5+140u0OIb6wCIQDEVKBMIEb1c11zTiZ+e/N4pyCLQWefVVxZrpXMYCmkAQ=="),
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
