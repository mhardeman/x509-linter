package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_caAuthorityKeyIdentifier_CheckApplies(t *testing.T) {
	CheckAppliesIntermediateCertificate(t, "caAuthorityKeyIdentifier", atis1000080.NewCaAuthorityKeyIdentifier)
}

var CERT_AKI_CA_WITH_EXT = ParseCert("MIIBdjCCARygAwIBAgIBATAKBggqhkjOPQQDAjAWMRQwEgYDVQQDEwtTSEFLRU4gUm9vdDAeFw0yMjEwMTMwODMyMDVaFw0yMjEwMTQwODMyMDVaMB4xHDAaBgNVBAMTE1NIQUtFTiBJbnRlcm1lZGlhdGUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATu//e00nULdqJ66ZWwvGJeCIcUfOjCTQWLl1Q6hHXFF6K0avuWNhcIQ3c6zQ1S4lVeU/Dvy6cTjVU1lyX3Braio1MwUTAPBgNVHRMBAf8EBTADAQH/MB8GA1UdIwQYMBaAFCVu0n1sUav6c/FRzrqjN0jYfQ5XMB0GA1UdDgQWBBRJpuqL5KEn+12rHvpElHRIeGG0qjAKBggqhkjOPQQDAgNIADBFAiEAj3HWkhm1ZnejC4tmDryav9KAWxajTnuithDBavvRdZACIB6tAsqS99yxucjcsYGiCwg2pbOyQG7PzASI3/y2aMqQ")
var CERT_AKI_CA_WITHOUT_EXT = ParseCert("MIIBVDCB+6ADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA4MzMxNVoXDTIyMTAxNDA4MzMxNVowHjEcMBoGA1UEAxMTU0hBS0VOIEludGVybWVkaWF0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJ6pmnQ94ljFtjTt7rJkpdRIixWoknTpdi5LFjETpYvCEVRuBWFeHIgBJIXRuz1LHMuJ+bO1/Lk90+MDQzvyk5SjMjAwMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFGNlctkm+tZi171qHPZLv2Bg/WFFMAoGCCqGSM49BAMCA0gAMEUCIEB00x/SFSi8tqSTmBvGPt7+uHI14B9Mo43Mewej93KFAiEA2XNkQ1euJmOGApnMvQQn2SMsRs5P12mFtaJ6B0QtK/c=")

func Test_caAuthorityKeyIdentifier_Execute(t *testing.T) {
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
			args: args{c: CERT_AKI_CA_WITH_EXT},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "extension is absent",
			args: args{c: CERT_AKI_CA_WITHOUT_EXT},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "STI certificates shall contain an Authority Key Identifier extension",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCaAuthorityKeyIdentifier()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caAuthorityKeyIdentifier.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
