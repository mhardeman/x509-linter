package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_certificatePolicies_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "extension is absent",
			args: args{c: TEST_CERT_VERSION_INCORRECT},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy",
			},
		},
		{
			name: "SHAKEN policy is absent",
			args: args{c: TEST_CERT_SHAKEN_POLICY_ABSENT},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy",
			},
		},
		{
			name: "SHAKEN policy is not single",
			args: args{c: TEST_CERT_SHAKEN_POLICY_NOT_SINGLE},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy",
			},
		},
		{
			name: "SHAKEN policy is correct",
			args: args{c: TEST_CERT_SHAKEN_POLICY_v1_1},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy",
			},
		},
		{
			name: "SHAKEN policy is correct",
			args: args{c: TEST_CERT_CORRECT},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCertificatePolicies()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("certificatePolicies.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
