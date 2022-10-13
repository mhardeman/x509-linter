package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_caCertificatePolicies_CheckApplies(t *testing.T) {
	CheckAppliesIntermediateCertificate(t, "caCertificatePolicies", atis1000080.NewCaCertificatePolicies)
}

var CERT_CP_CA_WITH_EXT = ParseCert("MIIBTTCB9aADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA4MzczNFoXDTIyMTAxNDA4MzczNFowHjEcMBoGA1UEAxMTU0hBS0VOIEludGVybWVkaWF0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDgSSnyGnbYD8vtd+4VUol30AbM//E6vPby22vvkYO8lGtgIQgxBI9L2FFE8z0MA5KGFRC0seW7Gm266TyuB86mjLDAqMA8GA1UdEwEB/wQFMAMBAf8wFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMAoGCCqGSM49BAMCA0cAMEQCIBzMEZ/nj2Y/xlVsEf/aZPE+CxIJfqM9IizFDn/XLev2AiAajNysINQFhD7wzozoYg2jhZNepWRcXq3wnONlmqfA5Q==")
var CERT_CP_CA_WITHOUT_EXT = ParseCert("MIIBNjCB3KADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA4Mzk0OFoXDTIyMTAxNDA4Mzk0OFowHjEcMBoGA1UEAxMTU0hBS0VOIEludGVybWVkaWF0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJlcEGSecOhd2cqfcS/YgFDFELn7kvim8jTPnxBRdP1WzWnE8BY/OF2NcLdtAVjkTc4n/hmzLmovA6Hb5+OBVWOjEzARMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSQAwRgIhALdKPfuYca1kNe8nN8MpagaMwMnBMM4SphUI26O70gY+AiEAy8MVQKYX27/CbHkSa7ZorRCJCfiFfWeNyFoCay0wg+o=")
var CERT_CP_CA_MULTIPLE_POLICIES = ParseCert("MIIBXTCCAQOgAwIBAgIBATAKBggqhkjOPQQDAjAWMRQwEgYDVQQDEwtTSEFLRU4gUm9vdDAeFw0yMjEwMTMwODQxNDRaFw0yMjEwMTQwODQxNDRaMB4xHDAaBgNVBAMTE1NIQUtFTiBJbnRlcm1lZGlhdGUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATVrHQKox3RjSxWP85mr9L1sXgW6RV2aNQGZgAbO4TRQ3iOchpxy2COlwUqkaV0mci6DvpYwYV8MlQno/O2Av8CozowODAPBgNVHRMBAf8EBTADAQH/MCUGA1UdIAQeMBwwDAYKYIZIAYb/CQEBAzAMBgpghkgBhv8JAQEBMAoGCCqGSM49BAMCA0gAMEUCIQDwkB5+2WqMZB1ZaN+pLfWNMCcN4O//EgaGHh9c34vWpwIgGJm4Sn+P/dtL9mKpwlyMoADaifmdMoFw3r1X66V3SrY=")

func Test_caCertificatePolicies_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "extension is correct",
			args: args{c: CERT_CP_CA_WITH_EXT},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "extension is absent",
			args: args{c: CERT_CP_CA_WITHOUT_EXT},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy",
			},
		},
		{
			name: "multiple policies",
			args: args{c: CERT_CP_CA_MULTIPLE_POLICIES},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCaCertificatePolicies()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caCertificatePolicies.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
