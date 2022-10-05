# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate 9df4720732ce5810ba9cf652a9086ee470a8a459
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/ae29b741-e66d-4105-87a1-7b45c5c2b9b0.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVzCCAf2gAwIBAgIIJAKhjaJQvzQwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMDkyMDEwMDU0OVoXDTIyMTAyMDEwMDU0OVowXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASbUoJJeK%2FL5n9CJwnL11s50KAWpkxaZIVPcIsjti%2FTCuozV37LcNDS%2FbJhy9COmdgFQV9G%2BUnL36rpzqHkMQEmo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNIADBFAiAprvqqmqHvyLzMk3tCMvVBTy2DnxlZyOoGt6yWLGYA2gIhAOZ8SBLoKazTKAbGV4Uuq%2FbJ7OV%2BotZ2jkh5QmfgQes%2B)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 7610' |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_key_identifier | error | STI certificates shall contain a Subject Key Identifier extension |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_ext_subject_key_identifier_missing_sub_cert | warn |  |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
