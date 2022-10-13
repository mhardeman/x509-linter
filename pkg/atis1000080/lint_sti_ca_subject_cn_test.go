package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_caSubjectCN_CheckApplies(t *testing.T) {
	CheckAppliesRootOrIntermediateCertificate(t, "caSubjectCN", atis1000080.NewCaSubjectCN)
}

func Test_caSubjectCN_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "CN=SHAKEN Intermediate",
			args: args{
				c: ParseCert("MIIBNTCB3KADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA5MTQ0NFoXDTIyMTAxNDA5MTQ0NFowHjEcMBoGA1UEAxMTU0hBS0VOIEludGVybWVkaWF0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLRPhWbwDnocnqmSDSQHmQ5wRIz9zLAdzPBBQHq0NQSCe5cCnQrMzq229XKhD6FpTZzEPV2RPYOL36wwGd/zO7OjEzARMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAMH6o/utJSJUp3o7A7O7kpnpF7ictXpxVYQrJAx0AzEtAiB2pSVGxdD9j4Or/fNjkdIwPjNBvhpP9walybCiSwMFKg=="),
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "CN=Intermediate",
			args: args{
				c: ParseCert("MIIBLjCB1aADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA5MTYwMVoXDTIyMTAxNDA5MTYwMVowFzEVMBMGA1UEAxMMSW50ZXJtZWRpYXRlMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfSkxKeEO2n5+MkBfEvACejFXRthcJgAs9OxsNJ58kXpLUhYQD4kLLDSwlsiYb2MAMWxZwSEHhoi+R2RSXG9qhKMTMBEwDwYDVR0TAQH/BAUwAwEB/zAKBggqhkjOPQQDAgNIADBFAiEAq0mAO7vIspVwG/S+Fe0sHTuMxDmEIse+cQUX85XgyR8CIDwEcf0aNvalZLaE8cQh3oJRmY2qGbX8bf0sI0XJqE/o"),
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The Common Name attribute shall include the text string \"SHAKEN\"",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCaSubjectCN()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caSubjectCN.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
