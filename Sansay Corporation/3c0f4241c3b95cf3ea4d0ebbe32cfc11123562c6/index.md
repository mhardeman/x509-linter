# SHAKEN COMPLIANCE
## Certificate

### Certificate 3c0f4241c3b95cf3ea4d0ebbe32cfc11123562c6
Tested At: 2022-10-05 08:12:01 +0000 UTC

Subject: CN=SHAKEN Xchange Telecom LLC 325B, OU=Bulk Solutions STI-AS, O=Xchange Telecom LLC, ST=New York, C=US

Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/325B_20211101.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDnDCCA0GgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkSZMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDkxMjE5NTQ0M1oXDTIzMDkxMjE5NTQ0M1owgYgxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEcMBoGA1UECgwTWGNoYW5nZSBUZWxlY29tIExMQzEeMBwGA1UECwwVQnVsayBTb2x1dGlvbnMgU1RJLUFTMSgwJgYDVQQDDB9TSEFLRU4gWGNoYW5nZSBUZWxlY29tIExMQyAzMjVCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0ufYApoQnxM2ww8zQ%2Be%2BEyvMooxbmK9ISQFoyTcGN6WoPxmvIXWBGBu1zv3aKBksXVrKOdAsE5j3ONV%2BWqorXqOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQzMjVCMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUxBZcjmcWiVQEmk3mP5tBZHdMiiEwgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6aMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAvn393Nkc3lr3Yujzg1mwHVTbWZEWrimPoOCbIdTw2xACIQCl4JrNw%2Fn0ZI%2BkXRMXBJbzloQZbWQRN%2BzDnV%2Bg6MxXBw%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 325B' |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

* The percent of certificates per issuer is calculated against total certificates from all issuers
** The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
