package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_subjectEmail_CheckApplies(t *testing.T) {
	CheckAppliesLeafCertificate(t, "subjectEmail", atis1000080.NewSubjectRdnUnknown)
}

func Test_subjectEmail_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "name without Email",
			args: args{c: TEST_CERT_CORRECT},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "name with Email",
			args: args{c: ParseCert("MIIBcjCCARigAwIBAgIBATAKBggqhkjOPQQDAjAeMRwwGgYDVQQDExNTSEFLRU4gSW50ZXJtZWRpYXRlMB4XDTIyMTAyMzA2MjMyNFoXDTIyMTAyNDA2MjMyNFowVTEUMBIGA1UEAxMLU0hBS0VOIHNvbWUxETAPBgNVBAoTCFNvbWUgb3JnMQswCQYDVQQGEwJVUzEdMBsGCSqGSIb3DQEJARYOc29tZUBlbWFpbC5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR4uIiaTeeIcpA9BCCQxdUpdBhvSc/gexPM554Jr9UOHPJr3ouOxRRd4u0a2P9pVEH98SNT0Qt9CwPsoSx4BsxFoxAwDjAMBgNVHRMBAf8EAjAAMAoGCCqGSM49BAMCA0gAMEUCIQDKvQxjHNpHM6nnP1+FixGxjVb4gGClBZYVboOy72fMCgIgOKguQ3L+J07Mily9jHtAaI27Y2O2q3FdiT7RzBVKrlY=")},
			want: &lint.LintResult{
				Status:  lint.Warn,
				Details: "Email addresses are not allowed as the CP does not specify how to validate them",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := atis1000080.NewSubjectEmail()
			if got := s.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subjectEmail.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
