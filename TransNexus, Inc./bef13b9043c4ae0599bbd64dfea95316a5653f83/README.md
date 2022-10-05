# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate bef13b9043c4ae0599bbd64dfea95316a5653f83
Tested At: 2022-10-05 17:08:39 +0000 UTC\
Subject: CN=SHAKEN 983J, OU=SHAKEN, O=ESI, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.clearip.com/59ebb7c1-25bd-4dfc-9794-fcb104b2f66a/37b8f7d06e9f6057267cbaa49b927ca7.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC5DCCAougAwIBAgIQTHyEJNP6UwnOTegWLmzbajAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTIyMDA3NTRaFw0yMjA5MTkyMDA3NTNaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNFU0kxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDk4M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQhQAkkRqBF7FXTqAnhgAJe9aZHBL9Y%2FV%2FCC2seeUQ4kSQ1FFQquyTc06s9Bnyt2vQWBTJuIO6v9nBgPq8vYM69o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFJgw7I%2F%2FtBWRBfgSuEetbBS3Yft9MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDk4M0owCgYIKoZIzj0EAwIDRwAwRAIgBhzz1Fu3YUEG9Vaz8nVk%2B0Jl7vZ5dESAnWHldAHKn4YCIEKMsBo%2FWkIKoBWY7gEqD16m4wSvfteXXer4LAXmA6yt)


| Code | Type | Details |
|------|------|---------|
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
