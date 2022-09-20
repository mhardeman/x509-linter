package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_ambiguousIdentifiers_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "CN is empty",
			args: args{c: TEST_CERT_VERSION_INCORRECT},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall contain TNAuthorizationList extension",
			},
		},
		{
			name: "CN is ambiguous", // CN=Some SHAKEN 123J
			args: args{c: TEST_CERT_AMBIGUOUS},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject",
			},
		},
		{
			name: "correct CN",
			args: args{c: TEST_CERT_CORRECT},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := atis1000080.NewAmbiguousIdentifiers()
			if got := a.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ambiguousIdentifiers.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
