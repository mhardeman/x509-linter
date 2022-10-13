package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_rootExtensionUnknown_CheckApplies(t *testing.T) {
	CheckAppliesRootCertificate(t, "rootExtensionUnknown", atis1000080.NewRootExtensionUnknown)
}

func Test_rootExtensionUnknown_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "list of allowed extension",
			args: args{
				c: ParseCert("MIIBezCCASGgAwIBAgIBATAKBggqhkjOPQQDAjAWMRQwEgYDVQQDEwtTSEFLRU4gUm9vdDAeFw0yMjEwMTMwOTU3MThaFw0yMjEwMTQwOTU3MThaMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEptTanEUzlVyx38gXsRJy3XGOLjbi1i4cldf+Acdmo1dgUzkliifmuNLXBSXwiwNo5pivsHVe4vNC5NGTDxZBtaNgMF4wDwYDVR0TAQH/BAUwAwEB/zALBgNVHQ8EBAMCAgQwHwYDVR0jBBgwFoAUOAjYd5vAegAHIl4qfA6ZoiQvvkQwHQYDVR0OBBYEFDgI2HebwHoAByJeKnwOmaIkL75EMAoGCCqGSM49BAMCA0gAMEUCIQCnVlD3eZIoXh3Y3pjwXasMG5feLtsO3X4/1mP7XnIbWQIgGBurRUVUMZ+G9XOhdhpdxwss9FxZXakRO7pqf5NBpkk="),
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "list with odd ext",
			args: args{
				c: ParseCert("MIIBizCCATKgAwIBAgIBATAKBggqhkjOPQQDAjAWMRQwEgYDVQQDEwtTSEFLRU4gUm9vdDAeFw0yMjEwMTMxMDAzMDVaFw0yMjEwMTQxMDAzMDVaMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEgbn5Bx0yCEaNyTngnlQzjijpOH9iGb3tCyt/28n2rAgRMwsuhKf9Ks6Z19MhlpXnXCHSvl9lt+iiPP1AO8rJLKNxMG8wDwYDVR0TAQH/BAUwAwEB/zALBgNVHQ8EBAMCAgQwHwYDVR0jBBgwFoAUMrHkj1131BjzPLCZSYLGJQxnCPAwHQYDVR0OBBYEFDKx5I9dd9QY8zywmUmCxiUMZwjwMA8GA1UdJQQIMAYGBCoDBAUwCgYIKoZIzj0EAwIDRwAwRAIgBzcS1hpNhfWax3yjCIKKewc/Qw8w45+OZsO3AFEfN1kCIDOIWKylCZLvdzkLvNr+ox2ASxmFFH+OgF3ILqGf9RNQ"),
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall not include extensions that are not specified",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := atis1000080.NewRootExtensionUnknown()
			if got := r.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rootExtensionUnknown.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
