# SHAKEN COMPLIANCE
## Certificate

### Certificate 5f7c718f5c9448ce21656eeea34dfc3630927a4b
Tested At: 2022-10-05 08:12:01 +0000 UTC

Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US

Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/d3f29d47-19de-4a66-9bc6-c4de17e27d2a.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVjCCAf2gAwIBAgIIEDlwoynsMAwwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMDkxOTEwMDU1MVoXDTIyMTAxOTEwMDU1MVowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASOmUyb%2FDxRGJ3BNH0aRr00iIfjmxdZxjUBOKL09Ebi6Za7ytF3MdN2YFLflmQ%2FqEqBKorouZTi3V6pSa34IC38o4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNHADBEAiAvvk6GylIlwvXFr7O6zB4DexHpfryY%2BL24Ua68HXLMuQIgD5iFoVN3WIT%2FBqzeyhh9bIQ6rWHqftZdOtfICO%2BqXtI%3D)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_ext_subject_key_identifier_missing_sub_cert | warn |  |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 7610' |
| e_shaken_sti_subject_key_identifier | error | STI certificates shall contain a Subject Key Identifier extension |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

* The percent of certificates per issuer is calculated against total certificates from all issuers
** The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
