package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_subjectPublicKey_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "incorrect key algorithm",
			args: args{
				c: TEST_CERT_SUBJECT_KEY_INCORRECT_CURVE,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall contain a Subject Public Key Info field specifying a Public Key Algorithm of \"id-ecPublicKey\" and containing a 256-bit public key",
			},
		},
		{
			name: "incorrect key algorithm",
			args: args{
				c: TEST_CERT_SUBJECT_KEY_INCORRECT_ALG,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall contain a Subject Public Key Info field specifying a Public Key Algorithm of \"id-ecPublicKey\" and containing a 256-bit public key",
			},
		},
		{
			name: "correct certificate version",
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
			s := atis1000080.NewSubjectPublicKey()
			if got := s.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subjectPublicKey.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
