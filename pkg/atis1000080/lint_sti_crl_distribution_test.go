package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_crlDistribution_Execute(t *testing.T) {
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
				Details: "STI intermediate and End-Entity certificates shall contain a CRL Distribution Points extension",
			},
		},
		{
			name: "CRL is multiple",
			args: args{c: TEST_CERT_CRL_MULTIPLE},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "CRL Distribution Points extension should contain a single DistributionPoint entry",
			},
		},
		{
			name: "CRL is reachable",
			args: args{c: TEST_CERT_CRL_REACHABLE},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs",
			},
		},
		{
			name: "extension is correct",
			args: args{c: TEST_CERT_CORRECT},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCRLDistribution()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("crlDistribution.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
