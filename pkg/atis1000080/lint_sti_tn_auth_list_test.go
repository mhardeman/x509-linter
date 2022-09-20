package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_tnAuthList_Execute(t *testing.T) {
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
				Details: "STI certificate shall contain TNAuthorizationList extension",
			},
		},
		{
			name: "TNAuthList has multiple TN entries",
			args: args{c: TEST_CERT_TN_ENTRY_MULTIPLE},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "TNAuthorizationList shall have only one TN Entry",
			},
		},
		{
			name: "TNAuthList has zero-length TN entry",
			args: args{c: TEST_CERT_TN_ENTRY_ZERO_LEN},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "TN Entry shall contain a SPC value",
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
			tr := atis1000080.NewTnAuthList()
			if got := tr.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tnAuthList.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
