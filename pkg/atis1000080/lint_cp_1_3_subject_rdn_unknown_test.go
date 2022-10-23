package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_subjectRdnUnknown_CheckApplies(t *testing.T) {
	CheckAppliesLeafCertificate(t, "subjectRdnUnknown", atis1000080.NewSubjectRdnUnknown)
}

func Test_subjectRdnUnknown_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "RDN is correct", // CN=, C=, O=, SERIALNUMBER=
			args: args{
				c: ParseCert("MIIBSTCB76ADAgECAgEBMAoGCCqGSM49BAMCMCwxCTAHBgNVBAMTADEJMAcGA1UEBhMAMQkwBwYDVQQKEwAxCTAHBgNVBAUTADAeFw0yMjA5MjIxNTAwMTVaFw0yMjA5MjMxNTAwMTVaMCwxCTAHBgNVBAMTADEJMAcGA1UEBhMAMQkwBwYDVQQKEwAxCTAHBgNVBAUTADBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCm7akhSiSoun8fuwpusqdrFlNuwnIQKZznBtmFLqF/b55jXppASRV1emveMJOeV6ljPy8JfeM/o640cUUk1uNmjAjAAMAoGCCqGSM49BAMCA0kAMEYCIQDcvhxKJPS0jMIylWtLgNq1Jdz3P1wi0Sb8VspgC4WWdAIhAOBuyYB2rVdYe0ejiVkNwEuZvQOVJ/ol1kLa0Dl0GQ0z"),
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "odd name", // CN=, C=, O=, SERIALNUMBER= , E=
			args: args{
				c: ParseCert("MIIBajCCARGgAwIBAgIBATAKBggqhkjOPQQDAjA9MQkwBwYDVQQDEwAxCTAHBgNVBAYTADEJMAcGA1UEChMAMQkwBwYDVQQFEwAxDzANBgkqhkiG9w0BCQEWADAeFw0yMjA5MjIxNTIwMThaFw0yMjA5MjMxNTIwMThaMD0xCTAHBgNVBAMTADEJMAcGA1UEBhMAMQkwBwYDVQQKEwAxCTAHBgNVBAUTADEPMA0GCSqGSIb3DQEJARYAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEpSYq7tSPXviGrfyjCipqZg0WMw0ANqT1xPkiEW7hGBwaOttK2ho0iy0Tx9L7lKdDs6PsAnSLS0aDjI09kyGwX6MCMAAwCgYIKoZIzj0EAwIDRwAwRAIgR26G4dVBJa8IOnxTzSpTcN58bRwEGisbBDJvYsG7jcsCIHJSwLW10+sEs+iBayN8UjyFXoiI/BPwo+b3ZuEPlpUF"),
			},
			want: &lint.LintResult{
				Status:  lint.Warn,
				Details: "No unknown RDNs are allowed as they may introduce ambiguity",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := atis1000080.NewSubjectRdnUnknown()
			if got := s.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subjectRdnUnknown.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
