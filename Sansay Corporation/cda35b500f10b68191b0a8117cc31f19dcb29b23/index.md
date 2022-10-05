# SHAKEN COMPLIANCE
## Certificate

### Certificate cda35b500f10b68191b0a8117cc31f19dcb29b23
Tested At: 2022-10-05 08:40:23 +0000 UTC\
Subject: CN=SHAKEN Consolidated Communications 5113, OU=NOC, O=Consolidated Communications, ST=New Hampshire, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Consolidated_5113.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDnTCCA0SgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkScEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDkxMzEzMzUxMVoXDTIyMTAxNDEzMzUxMVowgYsxCzAJBgNVBAYTAlVTMRYwFAYDVQQIDA1OZXcgSGFtcHNoaXJlMSQwIgYDVQQKDBtDb25zb2xpZGF0ZWQgQ29tbXVuaWNhdGlvbnMxDDAKBgNVBAsMA05PQzEwMC4GA1UEAwwnU0hBS0VOIENvbnNvbGlkYXRlZCBDb21tdW5pY2F0aW9ucyA1MTEzMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEIk3QPdu6bR%2BCslRuPfG5Xa3JWRiGOg1SH71pp9phvaR0Q%2Bot4r%2B62VqKUStE5bzD80dVHyhtP7o3X5%2F7tI7rTaOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ1MTEzMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUWWvtJS74qwN3g%2FEc0ZmIHpxzySwwgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6aMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiADH3mum22Ypbx2V%2FODTKvHdiJ051Tcj1JNBoyaCYxOPQIgci%2FbZgZMvJfoAqPySbuJnsgmws%2F9uzScH6pFpwOxG%2Bw%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 5113' |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
