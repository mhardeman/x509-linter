package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_crlCaDistribution_CheckApplies(t *testing.T) {
	CheckAppliesIntermediateCertificate(t, "caCrlDistribution", atis1000080.NewCaCrlDistribution)
}

func Test_crlCaDistribution_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "extension is correct",
			args: args{
				c: TEST_CERT_CORRECT,
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "extension is absent",
			args: args{
				c: CERT_CA,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI Intermediate certificates shall contain a CRL Distribution Points extension",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCaCrlDistribution()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caCrlDistribution.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
