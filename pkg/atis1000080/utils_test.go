package atis1000080_test

import (
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/atis1000080"
	"github.com/zmap/zcrypto/encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
)

func TestFindTNAuthListExtension(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want *pkix.Extension
	}{
		{
			name: "TNAuthList extensions exists",
			args: args{
				c: TEST_CERT_CORRECT,
			},
			want: &pkix.Extension{
				Id:       asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 26},
				Critical: false,
				Value:    []byte{48, 8, 160, 6, 22, 4, 55, 48, 57, 74},
			},
		},
		{
			name: "TNAuthList extensions does not exist",
			args: args{
				c: TEST_CERT_VERSION_INCORRECT,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := atis1000080.FindTNAuthListExtension(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTNAuthListExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
