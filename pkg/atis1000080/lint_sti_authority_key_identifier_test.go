package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_authorityKeyIdentifier_Execute(t *testing.T) {
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
			args: args{c: ParseCert("MIIBKzCB0qADAgECAgEBMAoGCCqGSM49BAMCMA0xCzAJBgNVBAMTAkNBMB4XDTIyMTAxMDE3MDUzMFoXDTIyMTAxMTE3MDUzMFowDzENMAsGA1UEAxMETGVhZjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGT9pIAnZXJjhmHw7fhQEdux9tFGuLZmzOuz3M8uPvQoM3pQpUiCh4fzn6GhWLf42NVFvPf4qo2gE9qvK1ZPaM+jITAfMB0GA1UdDgQWBBS20Rkse9fls6LwNqVgegtHAaqKQjAKBggqhkjOPQQDAgNIADBFAiA0IijaHQ6JBLje+n8cc9ubVeaiAs7qkEi302qLJQ4npwIhAPEh+/wP98++Fc5A2V0qaFX7oEQEUAU1BQs+tHzYrWwU")},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall contain an Authority Key Identifier extension",
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
			a := atis1000080.NewAuthorityKeyIdentifier()
			if got := a.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authorityKeyIdentifier.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorityKeyIdentifier_CheckApplies(t *testing.T) {
	CheckAppliesLeafCertificate(t, "authorityKeyIdentifier", atis1000080.NewAuthorityKeyIdentifier)
}
