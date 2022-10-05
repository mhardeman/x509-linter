# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate 1197bdcdf3c64e9d9b82e07ee707d1ca57e056ef
Tested At: 2022-10-05 17:17:11 +0000 UTC\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/2f9a70b9-e914-441c-8b80-9c233ad370c6.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIIHlAzgr4UGG8wCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMDkxNDEwMDU0OVoXDTIyMTAxNDEwMDU0OVowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASKLdAMqHdmY4ws3mA1Q4FbI%2BwCJFFpSmAeVln2EuWPwPxbbXnJDyI5515Ma%2B1M46DUm6MIHUSKw2IZU0uK2Y0ho4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNIADBFAiBPms%2BZUanRQ4C3S69Z54mZfluwy%2BFePwtpinIgAY0EMQIhANMKToyUD2Y5QKUhG0kbM37SPfMOdZrcZnLtTIPIX5I6)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_subject_key_identifier | error | STI certificates shall contain a Subject Key Identifier extension |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 7610' |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| w_ext_subject_key_identifier_missing_sub_cert | warn |  |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
