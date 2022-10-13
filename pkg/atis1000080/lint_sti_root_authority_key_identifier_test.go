package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

var CERT_AKI_ROOT_CORRECT = ParseCert("MIIBajCCARGgAwIBAgIBATAKBggqhkjOPQQDAjAWMRQwEgYDVQQDEwtTSEFLRU4gUm9vdDAeFw0yMjEwMTMwNzA4NDNaFw0yMjEwMTQwNzA4NDNaMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEIVgX3I8b30LQTW45mb+5IvKGrnAvtPgJZ8oQjVE/wyZZM9j0XdtKzMftpsJhnKo2ywkNIy6n8Tsn2xi5QEI4K6NQME4wDAYDVR0TBAUwAwEB/zAfBgNVHSMEGDAWgBR5F9vccud0rG8Ow/wXqINbCeeHGTAdBgNVHQ4EFgQUeRfb3HLndKxvDsP8F6iDWwnnhxkwCgYIKoZIzj0EAwIDRwAwRAIgIAgUNeXmLfKS2T2jBlxI8beHLx+884VK30MNr9nzrooCIB5ly9q9eZ+peVgrWCM5OsQsCkQ2sLx3ZsU7LJMqBX+m")
var CERT_AKI_ROOT_WITHOUT_EXT = ParseCert("MIIBSDCB8KADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA3MTMzN1oXDTIyMTAxNDA3MTMzN1owFjEUMBIGA1UEAxMLU0hBS0VOIFJvb3QwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATK23UcLYU4OPdogWQU109G3OWEPx1uDiOBnSgEfvHMl3qkYYnilcqLHnApSQSxvDaZ7zosjvaZ4FavdVda1waEoy8wLTAMBgNVHRMEBTADAQH/MB0GA1UdDgQWBBRa5eLlvrLmE9QHwJNpRsVEvUtKWjAKBggqhkjOPQQDAgNHADBEAiADHgjdAIIacQVFb1BTgyEF4WaybNYT2u94/vokJ5bBWwIgWQlnnt8cq1Sq6ZfKTxdtWtraCvkNXSgiS8Lr/Nii1kM=")
var CERT_AKI_ROOT_WITHOUT_SKI_EXT = ParseCert("MIIBSzCB8qADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA3MTkxNloXDTIyMTAxNDA3MTkxNlowFjEUMBIGA1UEAxMLU0hBS0VOIFJvb3QwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQZ2g7ZnEaGTkJh9W5CsP2Lub3FxE8Z9mzZvQgC9CGPd2qEUUtN3AiLKRhep3jLvTtfqgaFkPM2uxFtfuryhOTFozEwLzAMBgNVHRMEBTADAQH/MB8GA1UdIwQYMBaAFHlOQM0SMPXKvwhTiesQn5/OKZYqMAoGCCqGSM49BAMCA0gAMEUCICQsT+UdaeIIPtU4K/CQyuDp/jTH6Mbxj9GRccbDiMBVAiEA2WDC3QRErdVNkNg8zA83A5Oo93SP50Eck4qNcLYumGE=")
var CERT_AKI_ROOT_WITH_INCORRECT_SKI_EXT = ParseCert("MIIBazCCARGgAwIBAgIBATAKBggqhkjOPQQDAjAWMRQwEgYDVQQDEwtTSEFLRU4gUm9vdDAeFw0yMjEwMTMwNzIxMDZaFw0yMjEwMTQwNzIxMDZaMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEtq/J99t4dEmzxRyDD40Cvjjy7BTslXvrNt4um53xzebWEWjWYMFN6NByHZzYTiV3RCyCZ25hlN0ER+yrqlS1iKNQME4wDAYDVR0TBAUwAwEB/zAfBgNVHSMEGDAWgBSZqTB2rgoSkWgpWdAtjgx7nJ8kyTAdBgNVHQ4EFgQUfF35axGAUDdEsJoC5RZcfRPdnz0wCgYIKoZIzj0EAwIDSAAwRQIhAK8/6eBvw4RH0qMG8FxwEOJnHN5+EmswmBhzB8QWK7SOAiBb+LXxpKwNTcpjNuUb3FHAge/rqBvDdPMMyELmp07rug==")

func TestRootAuthorityKeyIdentifier_CheckApplies(t *testing.T) {
	CheckApplies(t, "rootAuthorityKeyIdentifier", atis1000080.NewRootAuthorityKeyIdentifier, []CheckAppliesVector{
		{
			name: "Leaf certtificate",
			args: CheckAppliesArgs{
				c: CERT_LEAF,
			},
			want: false,
		},
		{
			name: "Intermediat certtificate",
			args: CheckAppliesArgs{
				c: CERT_CA,
			},
			want: false,
		},
		{
			name: "Root certtificate without AKI ext",
			args: CheckAppliesArgs{
				c: CERT_ROOT,
			},
			want: false,
		},
		{
			name: "Root certtificate with AKI ext",
			args: CheckAppliesArgs{
				c: CERT_AKI_ROOT_CORRECT,
			},
			want: true,
		},
	})
}

func TestRootAuthorityKeyIdentifier_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "Root with correct AKI",
			args: args{
				c: CERT_AKI_ROOT_CORRECT,
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "Root without SKI ext",
			args: args{
				c: CERT_AKI_ROOT_WITHOUT_SKI_EXT,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI Root certificate shall contain a Subject Key Identifier extension",
			},
		},
		{
			name: "Root with incorrrect SKI ext",
			args: args{
				c: CERT_AKI_ROOT_WITH_INCORRECT_SKI_EXT,
			},
			want: &lint.LintResult{
				Status: lint.Error,
				Details: "Authority Key Identifier shall contain a keyIdentifier field with a value that matches " +
					"the Subject Key Identifier value of the same root certificate",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := atis1000080.NewRootAuthorityKeyIdentifier()
			if got := r.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rootAuthorityKeyIdentifier.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
