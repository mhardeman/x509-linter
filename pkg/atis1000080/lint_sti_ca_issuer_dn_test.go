package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_caIssuerDn_CheckApplies(t *testing.T) {
	CheckAppliesRootOrIntermediateCertificate(t, "caIssuerDn", atis1000080.NewCaIssuerDn)
}

func Test_caIssuerDn_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "correct DN",
			args: args{
				c: TEST_CERT_CORRECT,
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "CN is missed",
			args: args{
				c: CERT_DN_MISSED_CN,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute",
			},
		},
		{
			name: "O is missed",
			args: args{
				c: CERT_DN_MISSED_O,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute",
			},
		},
		{
			name: "C is missed",
			args: args{
				c: CERT_ISSUER_DN_MISSED_C,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := atis1000080.NewCaIssuerDn()
			if got := i.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("issuerDn.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
