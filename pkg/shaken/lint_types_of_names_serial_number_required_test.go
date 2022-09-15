package shaken

import (
	"encoding/base64"
	"reflect"
	"testing"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func parseCertificate(t *testing.T, name string, base64String string) *x509.Certificate {
	raw, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		t.Fatalf("Cannot decode the base64 string of the '%s' certificate. Error: %s", name, err.Error())
	}
	cert, err := x509.ParseCertificate(raw)
	if err != nil {
		t.Fatalf("Cannot parse the '%s' certificate. Error: %s", name, err.Error())
	}

	return cert
}

func Test_serialNumber_Execute(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}

	certWithoutSerialNumber := parseCertificate(t, "certWithoutSerialNumber", "MIHxMIGXoAMCAQICAQEwCgYIKoZIzj0EAwIwADAeFw0yMjA5MTUwOTA1MzNaFw0yMjA5MTYwOTA1MzNaMAAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASPsDG11is+olzAoXZfqK1Nv/gB8Xyei6h4geb9YJtm5/10q5T5cEjpwwuzJWw98M0A9YFAVK88JtKDZ4R6OR9QowIwADAKBggqhkjOPQQDAgNJADBGAiEAoekPEwEtdApQYH3uYca8yBmjXibOVN4BZ+LeNkPoDC4CIQC4/RP3Uw9/z+8PJKxUTjpNyccfUQVHr4p+N6dhjhtL/A==")
	certWithSerialNumber := parseCertificate(t, "certWithSerialNumber", "MIIBQjCB6aADAgECAgEBMAoGCCqGSM49BAMCMCkxEjAQBgNVBAMTCVNvbWUgbmFtZTETMBEGA1UEBRMKMTIzNDU2Nzg5MDAeFw0yMjA5MTUwODI1MDFaFw0yMjA5MTYwODI1MDFaMCkxEjAQBgNVBAMTCVNvbWUgbmFtZTETMBEGA1UEBRMKMTIzNDU2Nzg5MDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHQqN4uPwRz86StcsnL+gH/o3kH1IPjG1RI/HzngOEWCCOdZlnDv2FxB/GWHp4IcWM0MYfvOgy9Q0TmMloh6KCujAjAAMAoGCCqGSM49BAMCA0gAMEUCIQCjDlC681x1DMUxKtXquHtb+nFSTifql9s46LzrMM5qbQIgNywNI9rOKluHDmWb9ECiUmXNXNbsO5C6jx16iaQlXl0=")
	certWithMultipleSerialNumbers := parseCertificate(t, "certWithMultipleSerialNumbers", "MIIBbjCCAROgAwIBAgIBATAKBggqhkjOPQQDAjA+MSYwEAYDVQQDEwlTb21lIG5hbWUwEgYDVQQFEwsxMjM0NTY3ODkwMTEUMBIGA1UEBRMLMTIzNDU2Nzg5MDIwHhcNMjIwOTE1MDkwODQ5WhcNMjIwOTE2MDkwODQ5WjA+MSYwEAYDVQQDEwlTb21lIG5hbWUwEgYDVQQFEwsxMjM0NTY3ODkwMTEUMBIGA1UEBRMLMTIzNDU2Nzg5MDIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS3HV1r+pz5s1yGejse/dvxTCsRabTaA8mMB1Eugb5jGA4tvuonnZNczj2u9HEw77y5HiDTyVnSgB31MDkAqeDWowIwADAKBggqhkjOPQQDAgNJADBGAiEAoe3pl4SyzavbQsZ5NomdsCQ33TSBr2QbNlgLjGvA0poCIQDl3NRfVgScOB5QaF1CvLJeALWfn299VuaCD297J37gKA==")

	tests := []struct {
		name string
		l    *organizationIdentity
		args args
		want *lint.LintResult
	}{
		{
			name: "Certificate without the serial number attribute",
			l:    &organizationIdentity{},
			args: args{c: certWithoutSerialNumber},
			want: &lint.LintResult{Status: lint.Warn, Details: serialNumberShallBeIncluded},
		},
		{
			name: "Certificate with the serial number attribute",
			l:    &organizationIdentity{},
			args: args{c: certWithSerialNumber},
			want: &lint.LintResult{Status: lint.Pass},
		},
		{
			name: "Certificate with multiple serial number attributes",
			l:    &organizationIdentity{},
			args: args{c: certWithMultipleSerialNumbers},
			want: &lint.LintResult{Status: lint.Warn, Details: serialNumberShallBeIncludedOnlyOnce},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &organizationIdentity{}
			if got := l.Execute(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serialNumber.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
