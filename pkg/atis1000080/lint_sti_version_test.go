package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_version_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "incorrect certificate version",
			args: args{
				c: TEST_CERT_VERSION_INCORRECT,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall contain Version field specifying version 3",
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
			v := atis1000080.NewVersion()
			if got := v.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("version.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
