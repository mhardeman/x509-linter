# SHAKEN COMPLIANCE
## Certificate

### Certificate 29561a783e507f030a7fc1f4fb5900a9d2d0514c
Tested At: 2022-10-05 08:12:00 +0000 UTC

Subject: emailAddress=level5@lightspeedvoice.com, CN=SHAKEN Lightspeed Voice 557F, OU=NOC, O=Lightspeed Voice, ST=Florida, C=US, emailAddress=level5@lightspeedvoice.com

Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/557F/order/63_557F_83

View: [Click to view](https://understandingwebpki.com/?cert=MIIDrjCCA1OgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkQeEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDgxNjExNDM0MVoXDTIyMTAxNTExNDM0MVowgZoxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMRkwFwYDVQQKDBBMaWdodHNwZWVkIFZvaWNlMQwwCgYDVQQLDANOT0MxJTAjBgNVBAMMHFNIQUtFTiBMaWdodHNwZWVkIFZvaWNlIDU1N0YxKTAnBgkqhkiG9w0BCQEWGmxldmVsNUBsaWdodHNwZWVkdm9pY2UuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBXl3GCua%2BqMXp31Y8k4f0nN4gH7VMfG4sDv%2FiicinVgnyuVMVKGxE6ogbymRuI31eSsPawjrSItpvtyAgyUGcKOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ1NTdGMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUVyKFWLDI74e8KWf8JuXGw9Db0f0wgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6aMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAuzSIShvA6htC%2FW%2FpwX42Tc4yT%2BDC8U%2B4Hi%2BLes8Mx0UCIQDWo6xf1HX66BSM3Uv9AAeu%2BzjLirk%2FOkCpCruAF0SLSw%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 557F' |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

* The percent of certificates per issuer is calculated against total certificates from all issuers
** The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
