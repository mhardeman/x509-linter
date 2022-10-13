package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_basicConstraints_Execute(t *testing.T) {
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
				Details: "STI certificates shall contain a BasicConstraints extension marked critical",
			},
		},
		{
			name: "extension is not critical",
			args: args{c: TEST_CERT_BASIC_CONSTRAINTS_UNCRITICAL},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall contain a BasicConstraints extension marked critical",
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
			b := atis1000080.NewBasicConstraints()
			if got := b.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("basicConstraints.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicConstraints_CheckApplies(t *testing.T) {
	CheckAppliesAllCertificates(t, "basicConstraints", atis1000080.NewBasicConstraints)
}
