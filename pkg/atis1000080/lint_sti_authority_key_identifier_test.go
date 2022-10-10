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
		{
			name: "root cert without ext",
			args: args{c: ParseCert("MIIBKTCB0KADAgECAgEBMAoGCCqGSM49BAMCMA0xCzAJBgNVBAMTAkNBMB4XDTIyMTAxMDE3MDY1M1oXDTIyMTAxMTE3MDY1M1owDTELMAkGA1UEAxMCQ0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQl9iA2CwEV2iLLN4RAjccqaTrA0sPDrD/JJE+BgmWWanp2+tYDWa2V8Jxfg+942gmjjJ8z/D2OOXwxFVLpnRtGoyEwHzAdBgNVHQ4EFgQUmpp9LzXh4+/O50XZ1QEIUz+E9p4wCgYIKoZIzj0EAwIDSAAwRQIhAL2HPpp34AeEUENFx8nd+/kHGCR/iJ7dQC2GKoTGVBgPAiBFNAFVfWX28HTBt7NBbDeS4SN3SdkstXkNz6AcZM3/Wg==")},
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
