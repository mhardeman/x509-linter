package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_serialNumber_CheckApplies(t *testing.T) {
	CheckAppliesLeafCertificate(t, "serialNumber", atis1000080.NewSerialNumber)
}

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
