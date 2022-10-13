package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_caCertificatePolicyCritical_CheckApplies(t *testing.T) {
	CheckApplies(t, "caCertificatePolicyCritical", atis1000080.NewCaCertificatePolicyCritical, []CheckAppliesVector{
		{
			name: "Leaf certificate",
			args: CheckAppliesArgs{
				c: CERT_LEAF,
			},
			want: false,
		},
		{
			name: "Intermediate certificate with CP ext",
			args: CheckAppliesArgs{
				c: CERT_CP_CA_WITH_EXT,
			},
			want: true,
		},
		{
			name: "Intermediate certificate without CP ext",
			args: CheckAppliesArgs{
				c: CERT_CP_CA_WITHOUT_EXT,
			},
			want: false,
		},
		{
			name: "Root certificate",
			args: CheckAppliesArgs{
				c: CERT_ROOT,
			},
			want: false,
		},
	})
}

var CERT_CP_CA_CRITICAL = ParseCert("MIIBUTCB+KADAgECAgEBMAoGCCqGSM49BAMCMBYxFDASBgNVBAMTC1NIQUtFTiBSb290MB4XDTIyMTAxMzA4NTIxMloXDTIyMTAxNDA4NTIxMlowHjEcMBoGA1UEAxMTU0hBS0VOIEludGVybWVkaWF0ZTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPp0f/XqHzp17YfdniGvl1fEhMtmc2SEHQ842Xu+Om8n6VnFue75wG3sHB0kA0F+NHq+y3P4Rp4l+tOjJI9h3FSjLzAtMA8GA1UdEwEB/wQFMAMBAf8wGgYDVR0gAQH/BBAwDjAMBgpghkgBhv8JAQEDMAoGCCqGSM49BAMCA0gAMEUCIFwdc+7SNf0gcCd9tlbPI6QBuUUYa2dWH5zO2ceP4ZgcAiEA66/4kFO1aBqo6ns45iZFDWx+lgAjlq/INcdDCJ3Ju3s=")

func Test_caCertificatePolicyCritical_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "extension is uncritical",
			args: args{
				c: CERT_CP_CA_WITH_EXT,
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "extension is critical",
			args: args{
				c: CERT_CP_CA_CRITICAL,
			},
			want: &lint.LintResult{
				Status:  lint.Notice,
				Details: "STI certificates should contain a CertificatePolicies extension marked uncritical",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := atis1000080.NewCaCertificatePolicyCritical()
			if got := c.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caCertificatePolicyCritical.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
