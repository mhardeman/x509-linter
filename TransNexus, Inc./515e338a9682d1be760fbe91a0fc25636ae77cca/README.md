# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate 515e338a9682d1be760fbe91a0fc25636ae77cca
Tested At: 2022-10-05 17:17:16 +0000 UTC\
Subject: CN=SHAKEN 700H, OU=SHAKEN, O=Metro Fibernet LLC, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.clearip.com/22b6cee0-8559-4c73-8092-6eee861c4b49/b672bc1e98e6e21b33ea3e861055488c.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8zCCApqgAwIBAgIQQsZlfNX29payWhNoY4CazjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTQyMDE2MDJaFw0yMjA5MjEyMDE2MDFaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExJNZXRybyBGaWJlcm5ldCBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcwMEgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATAygPktTy2jSNmKihH4PqMDYmuaEAClxB%2Bz7Uv%2BK4EM6LzLAaKE1vZCeLNjt%2FqmUdV7SbXO5v3kZWDPtdHsisYo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFGKXIKfZPwIDmI0o4MSEBMTL3VaUMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcwMEgwCgYIKoZIzj0EAwIDRwAwRAIgP0uecsl7XcYJ14pFZ6lt2VChFzeQeYO4qrUOUMwCQ3QCIFkiQm%2Fmbd0M1S2eTIzIOMcp8ItRgQWEMTeq1ZjhFQUQ)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
