package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_caKeyUsageCrlSign_CheckApplies(t *testing.T) {
	CheckAppliesRootOrIntermediateCertificate(t, "caKeyUsageCrlSign", atis1000080.NewCaKeyUsageCrlSign)
}

func Test_caKeyUsageCrlSign_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "without crlSign",
			args: args{
				c: ParseCert("MIIBQjCB6aADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAyNDA2MzgwOFoXDTIyMTAyNTA2MzgwOFowHjEcMBoGA1UEAxMTU0hBS0VOIEludGVybWVkaWF0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDk5XTCtPyJZjE8JY3NMt0EQRXb70R1uCLtmWD4H/Zml8tCBDie8N2pWY1cBylRZyalr30r3qhMEi969qkEvqiWjIDAeMA8GA1UdEwEB/wQFMAMBAf8wCwYDVR0PBAQDAgIEMAoGCCqGSM49BAMCA0gAMEUCIQCnMbGghyv9p5vE5YvKwaB3E8hIVseeOvXFbEjLGqjT4gIgEeq0EaQHDvcvcSguI/QJ23HzNpEln8Upw9IIZ1VGv3g="),
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "with crlSign",
			args: args{
				c: ParseCert("MIIBQTCB6aADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAyNDA2Mzg1NloXDTIyMTAyNTA2Mzg1NlowHjEcMBoGA1UEAxMTU0hBS0VOIEludGVybWVkaWF0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJdLUxr9PzssFw3ReXVq48YFgXTtcDk+AQ+4Yn0VzgQWNXj2qFTcSuZ65OpzmCut6izJZZAQEBrUgaXOXwfxcaSjIDAeMA8GA1UdEwEB/wQFMAMBAf8wCwYDVR0PBAQDAgEGMAoGCCqGSM49BAMCA0cAMEQCIAuVNTdmVracqLMu3W0SmQ73Nt1caN4LLGSJJe6VV1jKAiAqbQcNQxjQf9loPPBo4gAKB24lj1UxeTOLoPECVNz6gw=="),
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCaKeyUsageCrlSign()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caKeyUsageCrlSign.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
