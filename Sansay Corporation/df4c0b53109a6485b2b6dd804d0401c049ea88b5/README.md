# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate df4c0b53109a6485b2b6dd804d0401c049ea88b5
Tested At: 2022-10-05 17:08:39 +0000 UTC\
Subject: CN=SHAKEN Airespring 996H, OU=Airespring NOC, O=Airespring, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Airespring_996H

View: [Click to view](https://understandingwebpki.com/?cert=MIIDhDCCAymgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTM8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE0NDQ0NVoXDTIyMTEwNDE0NDQ0NVowcTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAoMCkFpcmVzcHJpbmcxFzAVBgNVBAsMDkFpcmVzcHJpbmcgTk9DMR8wHQYDVQQDDBZTSEFLRU4gQWlyZXNwcmluZyA5OTZIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEO%2FDtGCxkRKOrL%2FCO4UTvqQyw1p4MqYMIqk92awEPU6Zj%2BIbwKfhahHQHw%2BYo3o%2BOQupLUMCFOnweTLm%2B2rEXzKOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ5OTZIMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUZc%2Fh%2FQbbTbRxhOwFMA%2FNOqB3wBQwgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6aMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAyiThpstPUZr%2Bfy6yzxXArDvyAYt51wC%2FOM9a%2F%2Bxgs6kCIQC%2BvbtZSZ%2FhDNaZD0AZDc5THD2gQd1spXX33COckLbCsQ%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 996H' |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
