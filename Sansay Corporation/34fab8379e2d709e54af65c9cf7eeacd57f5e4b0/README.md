# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 34fab8379e2d709e54af65c9cf7eeacd57f5e4b0
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: CN=SHAKEN  XCast Labs 689J, OU=XCast Labs, O=XCast Labs, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.xcastlabs.net/1663914600/xclsshaken.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDgTCCAyagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkQ0YwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDgyMzA2MzAwMFoXDTIyMDkyMzA2MzAwMFowbjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAoMClhDYXN0IExhYnMxEzARBgNVBAsMClhDYXN0IExhYnMxIDAeBgNVBAMMF1NIQUtFTiAgWENhc3QgTGFicyA2ODlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAECSaRM8oTiHWzZxnWX07VK9ClMrfygu659qCz3%2FNdIZetgYKmv2%2Fi%2B%2F3bRH4jjvo5INmhtTlMMWWnkQ0T%2Bn9R4qOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ2ODlKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUQu1fhldX0A%2ByhR0kbkEsERXho%2B4wgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6aMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEA6ypLYtzLKT3MXUbeHmKEAQNSMhc33FqC4cj5OobAuXoCIQDNPhvhm8MglmGFrfz04OOPYxqpbx07dWdjAVnnhU9jIw%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 689J' |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
