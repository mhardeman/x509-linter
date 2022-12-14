package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_certificatePolicyCritical_CheckApplies(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "extension is absent",
			args: args{c: TEST_CERT_VERSION_INCORRECT},
			want: false,
		},
		{
			name: "extension presents",
			args: args{c: TEST_CERT_CORRECT},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCertificatePolicyCritical()
			if got := c.CheckApplies(tt.args.c); got != tt.want {
				t.Errorf("certificatePolicyCritical.CheckApplies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_certificatePolicyCritical_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "CertificatePolicies extension is critical",
			args: args{c: ParseCert("MIH+MIGloAMCAQICAQEwCgYIKoZIzj0EAwIwADAeFw0yMjEwMDMwODI5NTRaFw0yMjEwMDQwODI5NTRaMAAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATZyy/7sJBc0mA36ryKIgIvJSxdyKAgbPuPsL+vf2BrCkLna6e9EQtZF+NJnRqG23UPu6Iz6cwMn77pnP/yWLOQoxAwDjAMBgNVHSABAf8EAjAAMAoGCCqGSM49BAMCA0gAMEUCIQCMI2It8BabbIA6Z+wTkz/YJsDHPFP3F+oQpiYrqGfGDAIgSZGloF0AvzL3xnEm3jr8/Nm4fVybO/PoyBc+AwT0824=")},
			want: &lint.LintResult{
				Status:  lint.Notice,
				Details: "STI certificates should contain a CertificatePolicies extension marked uncritical",
			},
		},
		{
			name: "CertificatePolicies extension is uncritical",
			args: args{c: TEST_CERT_CORRECT},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCertificatePolicyCritical()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("certificatePolicyCritical.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
