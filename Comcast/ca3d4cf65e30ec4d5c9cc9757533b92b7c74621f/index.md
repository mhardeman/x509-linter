# SHAKEN COMPLIANCE
## Certificate

### Certificate ca3d4cf65e30ec4d5c9cc9757533b92b7c74621f
Tested At: 2022-10-05 08:12:03 +0000 UTC

Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US

Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/49241c1a-4e4a-4bf8-9a6a-3de6850c21ac.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIIF%2FeuJOkXRUgwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMDkxNjEwMDU0OVoXDTIyMTAxNjEwMDU0OVowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQsrDtPeLTBlP%2BOwvfTOO4ccKh8zYhhxAqaQPeVVAebStZgmYrrwyaFQD5UlHChwhZ4yu7TfRjyCJx3DgVpPQq9o4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNIADBFAiB5zY7QFYpge1YNhb8d2NuDuiBTRQXePaIXJ1bE3jhf9AIhALlgzdSkBINesd9FQoPllUXc6VOs1GuFzbwcFNA0fAA4)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_subject_key_identifier | error | STI certificates shall contain a Subject Key Identifier extension |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 7610' |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_ext_subject_key_identifier_missing_sub_cert | warn |  |

* The percent of certificates per issuer is calculated against total certificates from all issuers
** The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
