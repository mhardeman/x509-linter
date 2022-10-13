package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func Test_pkiCaKeyUsage_CheckApplies(t *testing.T) {
	CheckApplies(t, "pkiCaKeyUsage", atis1000080.NewPkiCaKeyUsage, []CheckAppliesVector{
		{
			name: "Leaf certificate",
			args: CheckAppliesArgs{
				c: CERT_LEAF,
			},
			want: false,
		},
		{
			name: "CA with ext",
			args: CheckAppliesArgs{
				c: CERT_KU_CA_keyCertSign,
			},
			want: true,
		},
		{
			name: "CA without ext",
			args: CheckAppliesArgs{
				c: CERT_ROOT,
			},
			want: false,
		},
	})
}

func Test_pkiCaKeyUsage_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "keyCertSign flag only",
			args: args{
				c: CERT_KU_CA_keyCertSign,
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "without keyCertSign flag",
			args: args{
				c: CERT_KU_CA_WITHOUT_keyCertSign,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "For CA certificates, the Key Usage extension shall contain a single key usage value of keyCertSign",
			},
		},
		{
			name: "without keyCertSign and odd flag",
			args: args{
				c: CERT_KU_CA_keyCertSign_digitalSignature,
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "For CA certificates, the Key Usage extension shall contain a single key usage value of keyCertSign",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := atis1000080.NewPkiCaKeyUsage()
			if got := p.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pkiCaKeyUsage.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
