package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

var CERT_CP_ROOT_WITHOUT_EXT = ParseCert("MIIBKzCB0aADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA3NTAzNloXDTIyMTAxNDA3NTAzNlowFjEUMBIGA1UEAxMLU0hBS0VOIFJvb3QwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARnAlMO64apJouEQ4hlrMXExjejCoAN4RnaRdHEtXlQg6weMwvIztBbsGtKaWBNZaITUFmmbgeKLLg7di6UnpoWoxAwDjAMBgNVHRMEBTADAQH/MAoGCCqGSM49BAMCA0kAMEYCIQDS+HbOC9mmNP25Kpqsqx3XJaSBDqs3qktyeHkQHon4dQIhAPZhF7XDD/7yHUoER02xC4Rh2UIWNChC2qgX66PeQFLS")
var CERT_CP_ROOT_WITH_EXT = ParseCert("MIIBRDCB6qADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA3NTAwNVoXDTIyMTAxNDA3NTAwNVowFjEUMBIGA1UEAxMLU0hBS0VOIFJvb3QwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARl6bD4xTJINGnAQ5F5vWd4aQl+6LEiMYq/r46NCuetIDSv8Z0mfZwKNQemn4vh6kq64fVo6XLpcrJDHhl02MTnoykwJzAMBgNVHRMEBTADAQH/MBcGA1UdIAQQMA4wDAYKYIZIAYb/CQEBAzAKBggqhkjOPQQDAgNJADBGAiEApxr29B82ls6L8pCpUhqoPR067tsowuK2S2HEBzxiuysCIQCrU8XMKfNldd7rnCcyNpq5ASLh6NZvHdypHeGjdxVBTw==")

func TestRootCertificatePolicies_CheckApplies(t *testing.T) {
	CheckAppliesRootCertificate(t, "rootCertificatePolicies", atis1000080.NewRootCertificatePolicies)
}

func TestRootCertificatePolicies_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "root with CP ext",
			args: args{
				c: CERT_CP_ROOT_WITH_EXT,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI Root certificates shall not contain a Certificate Policy extension",
			},
		},
		{
			name: "root without CP ext",
			args: args{
				c: CERT_CP_ROOT_WITHOUT_EXT,
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := atis1000080.NewRootCertificatePolicies()
			if got := r.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rootCertificatePolicies.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
