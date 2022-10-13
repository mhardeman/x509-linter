package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_issuer_CheckApplies(t *testing.T) {
	CheckAppliesLeafCertificate(t, "issuerDn", atis1000080.NewIssuerDn)
}

var CERT_DN_MISSED_CN = ParseCert("MIIBZzCCAQygAwIBAgIBATAKBggqhkjOPQQDAjAgMREwDwYDVQQKEwhTb21lIG9yZzELMAkGA1UEAxMCVVMwHhcNMjIxMDEzMTAyODQ1WhcNMjIxMDE0MTAyODQ1WjA2MRQwEgYDVQQDEwtTSEFLRU4gc29tZTERMA8GA1UEChMIU29tZSBvcmcxCzAJBgNVBAMTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHmGsXXN+YqFRnUeEq9EfVRQ8cPUdiGKJCrN+WFqXn61B5GDjKf+ihIE3CTnBUHyzphmrAKp1c7pMeNFePcmK8aMhMB8wDAYDVR0TAQH/BAIwADAPBgNVHSUECDAGBgQqAwQFMAoGCCqGSM49BAMCA0kAMEYCIQC3i8xxIODyclJ+ypOnXOi69gzZV6z9zVWSzhG8A61rAwIhAJx+eDLnW2B6O6GZYG15qnYNOS9phIhHZyLIXeBDJ6Ni")
var CERT_DN_MISSED_O = ParseCert("MIIBcTCCARegAwIBAgIBATAKBggqhkjOPQQDAjArMRwwGgYDVQQDExNTSEFLRU4gSW50ZXJtZWRpYXRlMQswCQYDVQQDEwJVUzAeFw0yMjEwMTMxMDI5MDZaFw0yMjEwMTQxMDI5MDZaMDYxFDASBgNVBAMTC1NIQUtFTiBzb21lMREwDwYDVQQKEwhTb21lIG9yZzELMAkGA1UEAxMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARXjeOc+8KPswu2vxZ1pQNsyZSFStOpCe6I0b2LBLsqKIhSNfO8O3UyqeKe2rrGRnUwHqYktYTx+Kxdbu6BiSbKoyEwHzAMBgNVHRMBAf8EAjAAMA8GA1UdJQQIMAYGBCoDBAUwCgYIKoZIzj0EAwIDSAAwRQIhAI71YBv4b0ZVRD3w9Y8CATbLD+8wV44o28N4Psgk6C6rAiA8P23NCTN+PdrVvCTKq3QHM4FskLso8FXCYp3mBSUkgw==")
var CERT_ISSUER_DN_MISSED_C = ParseCert("MIIBdjCCAR2gAwIBAgIBATAKBggqhkjOPQQDAjAxMRwwGgYDVQQDExNTSEFLRU4gSW50ZXJtZWRpYXRlMREwDwYDVQQKEwhTb21lIG9yZzAeFw0yMjEwMTMxMDI5MjVaFw0yMjEwMTQxMDI5MjVaMDYxFDASBgNVBAMTC1NIQUtFTiBzb21lMREwDwYDVQQKEwhTb21lIG9yZzELMAkGA1UEAxMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARt3LvsTfQ4O3rC79XMwlBfun0Y+N2y/ecHRCbH2U5TApAt6wvVUHkt76q/NHlElcVnYp18o7ubXsKX0Dqr7c1aoyEwHzAMBgNVHRMBAf8EAjAAMA8GA1UdJQQIMAYGBCoDBAUwCgYIKoZIzj0EAwIDRwAwRAIgKAYP/LfMrvBdLIut5o4a+pWdTF6f+h2/1JaBADA8oL0CIDJcX8AwqZPRuBRRyrkWcRhQ40HFwi7ejwHxN/GkaHNK")

func Test_issuer_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "correct DN",
			args: args{
				c: TEST_CERT_CORRECT,
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "CN is missed",
			args: args{
				c: CERT_DN_MISSED_CN,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute",
			},
		},
		{
			name: "O is missed",
			args: args{
				c: CERT_DN_MISSED_O,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute",
			},
		},
		{
			name: "C is missed",
			args: args{
				c: CERT_ISSUER_DN_MISSED_C,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := atis1000080.NewIssuerDn()
			if got := i.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("issuerDn.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
