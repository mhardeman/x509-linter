package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/v3/lint"
)

func TestFindTNAuthListExtension(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *pkix.Extension
	}{
		{
			name: "TNAuthList extensions exists",
			args: args{
				c: TEST_CERT_CORRECT,
			},
			want: &pkix.Extension{
				Id:       asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 26},
				Critical: false,
				Value:    []byte{48, 8, 160, 6, 22, 4, 55, 48, 57, 74},
			},
		},
		{
			name: "TNAuthList extensions does not exist",
			args: args{
				c: TEST_CERT_VERSION_INCORRECT,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := atis1000080.FindTNAuthListExtension(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTNAuthListExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

type CheckAppliesArgs struct {
	c *x509.Certificate
}

type CheckAppliesVector struct {
	name string
	args CheckAppliesArgs
	want bool
}

func CheckApplies(t *testing.T, name string, linter func() lint.LintInterface, vectors []CheckAppliesVector) {
	for _, tt := range vectors {
		t.Run(tt.name, func(t *testing.T) {
			c := linter()
			if got := c.CheckApplies(tt.args.c); got != tt.want {
				t.Errorf("%s.CheckApplies() = %v, want %v", name, got, tt.want)
			}
		})
	}
}

var CERT_ROOT = ParseCert("MIIBLDCB1KADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA4MTMwOFoXDTIyMTAxNDA4MTMwOFowFjEUMBIGA1UEAxMLU0hBS0VOIFJvb3QwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQJfWAUcCzEneKUSYm8ukN29xObdIJnDeFyqxNw+dNMuEedQqVELDCXUAPNjfkx8Bc63Bf+oILi0cq0GikEKwRwoxMwETAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0cAMEQCIA7IFaiTWTRlIegoljUIxd5Gy1OuudHWhjoLWmS2OK57AiB52s9O9poNqAszV389XEgEjwgiC0HJDEY0yqUZeImRzQ==")
var CERT_CA = ParseCert("MIIBNTCB3KADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA4MTIzMloXDTIyMTAxNDA4MTIzMlowHjEcMBoGA1UEAxMTU0hBS0VOIEludGVybWVkaWF0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGz8fnDG0E/iSNwpVtOQDlCmm+b8T1a1WllUtdrIGkdYKVFcz4CC/IJpdNOkhRfMfeZM5/C5Yl79kwtlUyVoQo+jEzARMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIgYj7XEyu2Hepe7AhlyKAzFum8gaAE22sQXE4CSKi4jK0CIQDCknDyjoBjk1LLXf4IiFiEvqm7ECoDH0yYnam+7Wri+g==")
var CERT_LEAF = ParseCert("MIIBMzCB2aADAgECAgEBMAoGCCqGSM49BAMCMB4xHDAaBgNVBAMTE1NIQUtFTiBJbnRlcm1lZGlhdGUwHhcNMjIxMDEzMDgxNDAwWhcNMjIxMDE0MDgxNDAwWjAWMRQwEgYDVQQDEwtTSEFLRU4gTGVhZjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDNQ8LyzClMvRtlgG3Ub9traeAzzYM21tvXfXH+ge7km3RXqux9A1PzmIsbaKDYIYJvvw3lI/Q7RayyHvlqRuIejEDAOMAwGA1UdEwEB/wQCMAAwCgYIKoZIzj0EAwIDSQAwRgIhAMFrxYOGps1nltW6YFop9r08lG7d+tbWVXCasH0WU4HfAiEApBufv4sviBZU6ae0Gwf3trcjKbAM75n7rqzac9q+CJM=")

func CheckAppliesCertificate(t *testing.T, name string, linter func() lint.LintInterface, leaf bool, intermediate bool, root bool) {
	CheckApplies(t, name, linter, []CheckAppliesVector{
		{
			name: "Leaf certificate",
			args: CheckAppliesArgs{
				c: CERT_LEAF,
			},
			want: leaf,
		},
		{
			name: "Intermediate certificate",
			args: CheckAppliesArgs{
				c: CERT_CA,
			},
			want: intermediate,
		},
		{
			name: "Root certificate",
			args: CheckAppliesArgs{
				c: CERT_ROOT,
			},
			want: root,
		},
	})
}

func CheckAppliesLeafCertificate(t *testing.T, name string, linter func() lint.LintInterface) {
	CheckAppliesCertificate(t, name, linter, true, false, false)
}

func CheckAppliesIntermediateCertificate(t *testing.T, name string, linter func() lint.LintInterface) {
	CheckAppliesCertificate(t, name, linter, false, true, false)
}

func CheckAppliesRootCertificate(t *testing.T, name string, linter func() lint.LintInterface) {
	CheckAppliesCertificate(t, name, linter, false, false, true)
}

func CheckAppliesRootOrIntermediateCertificate(t *testing.T, name string, linter func() lint.LintInterface) {
	CheckAppliesCertificate(t, name, linter, false, true, true)
}

func CheckAppliesAllCertificates(t *testing.T, name string, linter func() lint.LintInterface) {
	CheckAppliesCertificate(t, name, linter, true, true, true)
}
