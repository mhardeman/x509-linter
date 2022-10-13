package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_extensionUnknown_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "odd extension",
			args: args{
				c: ParseCert("MIH6MIGgoAMCAQICAQEwCgYIKoZIzj0EAwIwADAeFw0yMjA5MjIxMzQ5MTBaFw0yMjA5MjMxMzQ5MTBaMAAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT9DSfUmIBZgG1um/05ZxOgUTPtk5l8G8Ifuc7SAlDPTWRtNzefj5lgjDC11PWB7zljh+oPaPc7uUJJX5CGNU7CowswCTAHBgNVHQEEADAKBggqhkjOPQQDAgNJADBGAiEAyCKabyG717Gur+YGeLCCMpjsDLioHScBoNWVDPAQpgICIQCtluJwlyFJ/1cgDXn4GKWFO1WRxFEZJQF8d3GRPHMKKw=="),
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall not include extensions that are not specified",
			},
		},
		{
			name: "correct extensions",
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
			e := atis1000080.NewExtensionUnknown()
			if got := e.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extensionUnknown.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extensionList_Contains(t *testing.T) {
	CheckAppliesLeafCertificate(t, "extensionList", atis1000080.NewExtensionUnknown)
}
