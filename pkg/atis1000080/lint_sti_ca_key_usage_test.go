package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

var CERT_KU_CA_keyCertSign = ParseCert("MIIBETCBuKADAgECAgEBMAoGCCqGSM49BAMCMAAwHhcNMjIxMDAzMTUwNjUzWhcNMjIxMDA0MTUwNjUzWjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEZE4tMNzgxSiOD58aQHC4Kk2SSUpTCpqbCgaMJPeYBZMNNn2lAvaX+E/1ZWZmb+ElbieLSEgrY9nuUC//5BephqMjMCEwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAgQwCgYIKoZIzj0EAwIDSAAwRQIgQsoxlZWBhveYlvW3JA0Q6sQqY2jT2ijUjhBj8kfS39ACIQC5+oaKj538OFCxqhdyK85hUCTh2TYW9WDl+gXJHARFTw==")
var CERT_KU_CA_WITHOUT_keyCertSign = ParseCert("MIIBEjCBuKADAgECAgEBMAoGCCqGSM49BAMCMAAwHhcNMjIxMDAzMTUxMjUxWhcNMjIxMDA0MTUxMjUxWjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEMR86UkOcy2aZYuASNz0QcR3RrD0U6AkoFVDSwJv8eGZMu0MVL6lsAJr9L0LAcLUZ4q7nmZ61seGrQL5Gx+BRnaMjMCEwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAI1xmeTmKE6ryG/6U74/j96NzEdaco9B8djI8h/6ZKQLAiEA+7u434VC6tYemgqYqGlX3OfjN2Sz29syTQElI1YxXxk=")
var CERT_KU_CA_keyCertSign_digitalSignature = ParseCert("MIIBETCBuKADAgECAgEBMAoGCCqGSM49BAMCMAAwHhcNMjIxMDAzMTUxMjA2WhcNMjIxMDA0MTUxMjA2WjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7kp9whB/AdqlOSUntQ7WBbW3AYdlc/ppe3daPsQnKYwZc+I8KWmUE1FefH2e77P4Vh3RVsBUI/UXYpBwRYNR7qMjMCEwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAoQwCgYIKoZIzj0EAwIDSAAwRQIgaOdJM0woR2cVGENr+ljXwMedasasrThmC3ftKfHk0HQCIQDoDNTiwSOK2xcSrhp1XXMShjqKnOQK+al16cvEjuN6QQ==")

func Test_caKeyUsage_CheckApplies(t *testing.T) {
	CheckAppliesRootOrIntermediateCertificate(t, "caKeyUsage", atis1000080.NewCaKeyUsage)
}

func Test_caKeyUsage_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "KeyUsage with keyCertSign flag",
			args: args{
				c: CERT_KU_CA_keyCertSign,
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "KeyUsage with keyCertSign, crlSign flags",
			args: args{
				c: ParseCert("MIIBETCBuKADAgECAgEBMAoGCCqGSM49BAMCMAAwHhcNMjIxMDAzMTUwOTI4WhcNMjIxMDA0MTUwOTI4WjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbfvKFTyMfc0uQOcruMkb5Qp8onD4nwhOVnnFVECT3PI5du8UXUreOw1Mp0z/bO9YTDvMEhjGTR03gW3shk7jt6MjMCEwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwCgYIKoZIzj0EAwIDSAAwRQIhAP7+pascNAa2Ivqli8daA9lQ7tZQjlbM83G1SOCCjAOWAiAuaipKBkoW3F4T2FdgxpbWxfTbV/sxKTsckRQ51U6IPw=="),
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "KeyUsage with keyCertSign, crlSign, digitalSignature flags",
			args: args{
				c: ParseCert("MIIBETCBuKADAgECAgEBMAoGCCqGSM49BAMCMAAwHhcNMjIxMDAzMTUxMDE3WhcNMjIxMDA0MTUxMDE3WjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGRGaZQsM9yRE5LmZvDWKQGwYe94JSG1Rcr9uthrNYiJht1XTE7SYHl4teRo2e5tDenkWVzRCmmOJ7a7pHzF6S6MjMCEwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAYYwCgYIKoZIzj0EAwIDSAAwRQIhALuE/vQ3pi9B5FvVMZUCsDHJIScwvo51a3fXfRIFCj79AiBn/Vh5p+yDoy0mAn2iGfERC9IRsCHHq129gXMWHwVGng=="),
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "KeyUsage with keyCertSign, digitalSignature flags",
			args: args{
				c: CERT_KU_CA_keyCertSign_digitalSignature,
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "KeyUsage without keyCertSign flags",
			args: args{
				c: CERT_KU_CA_WITHOUT_keyCertSign,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The Key Usage extension shall contain the key usage value keyCertSign, and may contain the key usage values digitalSignature and/or cRLSign",
			},
		},
		{
			name: "KeyUsage with odd flag",
			args: args{
				c: ParseCert("MIIBETCBuKADAgECAgEBMAoGCCqGSM49BAMCMAAwHhcNMjIxMDAzMTUxNDM2WhcNMjIxMDA0MTUxNDM2WjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAESGp0fzaAGRQQ4ncXr1Gvrjyov+0NOLJMKTUlp+n3sRpa3IAelaaK6zIk/hJcxgAUoFKYA7scNTiEQgf6kRvbEKMjMCEwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAkQwCgYIKoZIzj0EAwIDSAAwRQIgah1nkwC5+0M96caNcqbsmFwYbM65ke4hyS9ZlKXYRSoCIQDP0qfqDm7dSJfh30o0qN92/33CxsfLzR3N5BTg34xNWg=="),
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The Key Usage extension shall contain the key usage value keyCertSign, and may contain the key usage values digitalSignature and/or cRLSign",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCaKeyUsage()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caKeyUsage.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
