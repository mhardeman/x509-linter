package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_caSubjectRdnUnknown_CheckApplies(t *testing.T) {
	CheckAppliesRootOrIntermediateCertificate(t, "caSubjectRdnUnknown", atis1000080.NewCaSubjectRdnUnknown)
}

func Test_caSubjectRdnUnknown_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "RDN is correct", // CN, C, O
			args: args{
				c: ParseCert("MIIBZjCCAQ2gAwIBAgIBATAKBggqhkjOPQQDAjAWMRQwEgYDVQQDEwtTSEFLRU4gUm9vdDAeFw0yMjEwMTMxMTM0NTdaFw0yMjEwMTQxMTM0NTdaMD4xHDAaBgNVBAMTE1NIQUtFTiBJbnRlcm1lZGlhdGUxETAPBgNVBAoTCFNvbWUgb3JnMQswCQYDVQQDEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABBgsKpePEDU+V8xQHBS8359bD2SuHs+ZE2yXu93pes+DQqLFVyyyv35yuUeuHMHc9O2gcK/vikXlpmnWMsAMOl2jJDAiMA8GA1UdEwEB/wQFMAMBAf8wDwYDVR0lBAgwBgYEKgMEBTAKBggqhkjOPQQDAgNHADBEAiA9pWaTu69uBd8OGLFJIAVxIoOKbQ7cgdT9JmJaMVAFZwIgAeAWv++lZNlQWJgNsAjrGB83DY52UfODN+rtV9854S0="),
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "odd name", // CN, C, O, E
			args: args{
				c: ParseCert("MIIBhjCCASygAwIBAgIBATAKBggqhkjOPQQDAjAWMRQwEgYDVQQDEwtTSEFLRU4gUm9vdDAeFw0yMjEwMTMxMTM1NDNaFw0yMjEwMTQxMTM1NDNaMF0xHDAaBgNVBAMTE1NIQUtFTiBJbnRlcm1lZGlhdGUxETAPBgNVBAoTCFNvbWUgb3JnMQswCQYDVQQDEwJVUzEdMBsGCSqGSIb3DQEJARYOc29tZUBlbWFpbC5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQiyMVByfKmrynuD5oyFdkMObug96tHSA23oftau0ydh6Q0P1uFAW8z0xLH2FH15lxsHGrkrkn5W34jlFapt+1HoyQwIjAPBgNVHRMBAf8EBTADAQH/MA8GA1UdJQQIMAYGBCoDBAUwCgYIKoZIzj0EAwIDSAAwRQIhANKjK2EfLsEwjXEOHAWOxxuw6max9FtIwYuALwKty3naAiBf/JOHgvrwAtALI3ONaFus+fJ5o4r34qNvQW5c9qak8g=="),
			},
			want: &lint.LintResult{
				Status:  lint.Warn,
				Details: "Only CN, C, and O can be included. Additional RNDs may introduce ambiguity and may not be verifiable",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := atis1000080.NewCaSubjectRdnUnknown()
			if got := s.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subjectRdnUnknown.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
