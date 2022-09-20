package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_subjectSN_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "SERIALNUMBER is absent",
			args: args{c: TEST_CERT_VERSION_INCORRECT},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall include a ‘serialNumber’ attribute along with the CN",
			},
		},
		{
			name: "multiple SERIALNUMBER",
			args: args{c: TEST_CERT_SUBJECT_SERIALNUMBER_ODD},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall include a ‘serialNumber’ attribute along with the CN",
			},
		},
		{
			name: "SERIALNUMBER presents",
			args: args{c: TEST_CERT_SUBJECT_SERIALNUMBER},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := atis1000080.NewSubjectSN()
			if got := s.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subjectSN.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
