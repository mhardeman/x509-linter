package shaken

import (
	"reflect"
	"testing"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_organizationIdentity_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}

	certWithoutSerialNumber := parseCertificate(t, "certWithoutSerialNumber", "MIHxMIGXoAMCAQICAQEwCgYIKoZIzj0EAwIwADAeFw0yMjA5MTUwOTA1MzNaFw0yMjA5MTYwOTA1MzNaMAAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASPsDG11is+olzAoXZfqK1Nv/gB8Xyei6h4geb9YJtm5/10q5T5cEjpwwuzJWw98M0A9YFAVK88JtKDZ4R6OR9QowIwADAKBggqhkjOPQQDAgNJADBGAiEAoekPEwEtdApQYH3uYca8yBmjXibOVN4BZ+LeNkPoDC4CIQC4/RP3Uw9/z+8PJKxUTjpNyccfUQVHr4p+N6dhjhtL/A==")

	tests := []struct {
		name string
		l    *organizationIdentity
		args args
		want *lint.LintResult
	}{
		{
			name: "countryName is missed",
			l:    &organizationIdentity{},
			args: args{c: certWithoutSerialNumber},
			want: &lint.LintResult{Status: lint.Warn, Details: organizationIdentityCountryName},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &organizationIdentity{}
			if got := l.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("organizationIdentity.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
