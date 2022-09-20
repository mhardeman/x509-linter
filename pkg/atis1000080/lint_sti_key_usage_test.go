package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_keyUsage_Execute(t *testing.T) {
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
				Details: "STI certificates shall contain a Key Usage extension marked as critical",
			},
		},
		{
			name: "keyUsage is uncritical",
			args: args{c: TEST_CERT_KEY_USAGE_UNCRITICAL},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall contain a Key Usage extension marked as critical",
			},
		},
		{
			name: "keyUsage use multiple flags",
			args: args{c: TEST_CERT_KEY_USAGE_MULTIPLE_FLAGS},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The Key Usage extension shall contain a single key usage value of digitalSignature",
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
			k := atis1000080.NewKeyUsage()
			if got := k.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keyUsage.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
