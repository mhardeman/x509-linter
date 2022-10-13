package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

var CERT_SUBJECT_WITHOUT_CN = ParseCert("MIIBTzCB96ADAgECAgEBMAoGCCqGSM49BAMCMB4xHDAaBgNVBAMTE1NIQUtFTiBJbnRlcm1lZGlhdGUwHhcNMjIxMDEzMTMxNjUxWhcNMjIxMDE0MTMxNjUxWjAgMREwDwYDVQQKEwhTb21lIG9yZzELMAkGA1UEAxMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASV6bm4Z1rg9GHoLwVDKjJxbCuUc1x+PR9sfSVVJ6y127+r/CL1AGFutQGoP3vRKa4iIOpW6uobgpII2bkZklLMoyQwIjAPBgNVHRMBAf8EBTADAQH/MA8GA1UdJQQIMAYGBCoDBAUwCgYIKoZIzj0EAwIDRwAwRAIgCrf09OT7mge+vYLOQS4j74cs55y2FC/kitSuYGz1qpYCIGU4Gxb1de6ua6Y8EI0jqUBPFS+EP5DIo2h0Op2SXvYY")
var CERT_SUBJECT_WITHOUT_O = ParseCert("MIIBVDCB+qADAgECAgEBMAoGCCqGSM49BAMCMB4xHDAaBgNVBAMTE1NIQUtFTiBJbnRlcm1lZGlhdGUwHhcNMjIxMDEzMTMxNzQ4WhcNMjIxMDE0MTMxNzQ4WjAjMRQwEgYDVQQDEwtTSEFLRU4gTGVhZjELMAkGA1UEAxMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS+vd5oj/4byiZgfaZhVIwfKg7mmvf4XENRgT0ih2n9PCl/p+CW1zoIJHJ6YPo3oYpuKYVx0XmTWkydOaLijTMNoyQwIjAPBgNVHRMBAf8EBTADAQH/MA8GA1UdJQQIMAYGBCoDBAUwCgYIKoZIzj0EAwIDSQAwRgIhAIKbfXaZSQhfLyItemyOM402Xu0VAarRu3SW3iVJwggbAiEAgdvAPjoFkY4Zqe9GgbKOk+67E9iOue6wsP+1+3lGE/g=")
var CERT_SUBJECT_WITHOUT_C = ParseCert("MIIBWjCCAQCgAwIBAgIBATAKBggqhkjOPQQDAjAeMRwwGgYDVQQDExNTSEFLRU4gSW50ZXJtZWRpYXRlMB4XDTIyMTAxMzEzMTgyOFoXDTIyMTAxNDEzMTgyOFowKTEUMBIGA1UEAxMLU0hBS0VOIExlYWYxETAPBgNVBAoTCFNvbWUgb3JnMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEqTdo0Y9clZ4cFi0kt2+loHyHgPhIv1ITdzaZwr1L9Rv3Kl2RCd4YlyElOkFPtJd6klByQh8C0mgSSolIgKr9N6MkMCIwDwYDVR0TAQH/BAUwAwEB/zAPBgNVHSUECDAGBgQqAwQFMAoGCCqGSM49BAMCA0gAMEUCIQCcaUCwi2mRNt2OFZAONUj/P99RYXEZNCDH4iLPwxMODQIgSy2RQ7PYWiZOWhFb7KvKXhxRRBoHCXgRhn3iyXpQQR8=")

func Test_subject_CheckApplies(t *testing.T) {
	CheckAppliesLeafCertificate(t, "subject", atis1000080.NewSubject)
}

func Test_subject_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "CN is missed",
			args: args{
				c: CERT_SUBJECT_WITHOUT_CN,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute",
			},
		},
		{
			name: "O is missed",
			args: args{
				c: CERT_SUBJECT_WITHOUT_O,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute",
			},
		},
		{
			name: "C is missed",
			args: args{
				c: CERT_SUBJECT_WITHOUT_C,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute",
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
			v := atis1000080.NewSubject()
			if got := v.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("version.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
