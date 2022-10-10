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
			args: args{c: ParseCert("MIIBDDCBs6ADAgECAgEBMAoGCCqGSM49BAMCMA0xCzAJBgNVBAMTAkNBMB4XDTIyMTAxMDE3MjAyMVoXDTIyMTAxMTE3MjAyMVowDzENMAsGA1UEAxMETGVhZjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGG0PGLskEbiBY0xDuYg8vPMzcTyLvhBPl40uARkuu4uTScK2fQL5VP7d3t3JkPEmstVQS4Cqc/fvVJRQIdKV06jAjAAMAoGCCqGSM49BAMCA0gAMEUCIQCU2PrUfAocvNNzP2du/77S+fBR4wLu7ug3XVwvTISJVwIgEOOFivZJa0MqmWkwqB9iA1KwiyfgW6k2tATHEw7aafo=")},
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
