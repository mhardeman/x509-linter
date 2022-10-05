# SHAKEN COMPLIANCE
## Certificate

### Certificate 0bbfe21d475913d8501a0f85af92eca5d801b7d1
Tested At: 2022-10-05 08:12:00 +0000 UTC

Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US

Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/28ad7b45-d26e-44f9-9e94-956be1447094.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICWDCCAf2gAwIBAgIIRIQd8zU5seowCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMDkxNTEwMDU1MFoXDTIyMTAxNTEwMDU1MFowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR2rlWUg8Hfu41uMitktfNz9%2FH14w2r4XwcvIKFK%2FI5RDwzoYOs89XV8pR%2BS%2B8zZ0wCOLOxkMQeAFqJ2H%2Bf0yPdo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNJADBGAiEAut%2FNmnCyYeiOpHMWS43%2BrvS28zr9Sb4a5qLDLhzewekCIQCY4HXaoqsz47zv0Dmvb9Whq7exSuPPUOV01lfJ3SwY8w%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_subject_key_identifier | error | STI certificates shall contain a Subject Key Identifier extension |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_ext_subject_key_identifier_missing_sub_cert | warn |  |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 7610' |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

* The percent of certificates per issuer is calculated against total certificates from all issuers
** The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
