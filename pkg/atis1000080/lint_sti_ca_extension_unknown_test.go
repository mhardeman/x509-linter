package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_caExtensionUnknown_CheckApplies(t *testing.T) {
	CheckAppliesIntermediateCertificate(t, "caExtensionUnknown", atis1000080.NewCaExtensionUnknown)
}

func Test_caExtensionUnknown_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "list of allowed extensions",
			args: args{
				c: ParseCert("MIIB6jCCAZCgAwIBAgIBATAKBggqhkjOPQQDAjAWMRQwEgYDVQQDEwtTSEFLRU4gUm9vdDAeFw0yMjEwMTMwOTUyNThaFw0yMjEwMTQwOTUyNThaMB4xHDAaBgNVBAMTE1NIQUtFTiBJbnRlcm1lZGlhdGUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASM68ds/3oWfpHDy+XyfpJnAJCy5QySZWEl+prmOyKM1t47IZFOd9p1MrY+SnT0kopSRkucWx9cqb0Hha/loOE/o4HGMIHDMA8GA1UdEwEB/wQFMAMBAf8wCwYDVR0PBAQDAgIEMBoGA1UdIAEB/wQQMA4wDAYKYIZIAYb/CQEBAzAfBgNVHSMEGDAWgBTnfGzb+g3jfI/nuR21zl+B7CVxPTAdBgNVHQ4EFgQUIxLII8DwDpY+S+/Dp+YBjZQ8yVYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAoGCCqGSM49BAMCA0gAMEUCIEftKf4rnzvABPnmqwLyp6i9C0fMVOj7JuW3esy6eHgkAiEAlZVSS1U5XxOWrRdaMqiKIyu6E35bHf4gA5Cj2N7QSmM="),
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "list with odd ext",
			args: args{
				c: ParseCert("MIIBNTCB3KADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA5NTQyN1oXDTIyMTAxNDA5NTQyN1owHjEcMBoGA1UEAxMTU0hBS0VOIEludGVybWVkaWF0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABItEI7mXcMWN9+zVkA3yAtixM3vWvk2LCWA/ohh5qexRDeGsCdRlh641ofaJfraP9R9GnVJZqDy34f6/hrDz0AqjEzARMA8GA1UdJQQIMAYGBCoDBAUwCgYIKoZIzj0EAwIDSAAwRQIhAKhk4QYBOzTRWgpw/IPgVLiLwgbydHMGL2fYLov0RW+iAiBOlOFZs0GBRq5rHvqP1ZRS437WjzcFwsFnTLcsru3eoQ=="),
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall not include extensions that are not specified",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCaExtensionUnknown()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caExtensionUnknown.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
