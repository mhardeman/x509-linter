package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_subjectKeyIdentifier_CheckApplies(t *testing.T) {
	CheckAppliesLeafCertificate(t, "subjectKeyIdentifier", atis1000080.NewSubjectKeyIdentifier)
}

func Test_subjectKeyIdentifier_Execute(t *testing.T) {
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
			args: args{c: CERT_LEAF},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall contain a Subject Key Identifier extension",
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
			s := atis1000080.NewSubjectKeyIdentifier()
			if got := s.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subjectPublicKeyIdentifier.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
