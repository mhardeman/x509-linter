# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate e3454582df5a1a64e5c7a3d7b99d6f65ac952e1d
Tested At: 2022-10-05 17:08:39 +0000 UTC\
Subject: CN=SHAKEN Apeiron Systems 012J, OU=NOC, O=Apeiron Systems, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Apeiron_012J

View: [Click to view](https://understandingwebpki.com/?cert=MIIDgjCCAyigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTMQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE0NDExNVoXDTIyMTEwNDE0NDExNVowcDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExGDAWBgNVBAoMD0FwZWlyb24gU3lzdGVtczEMMAoGA1UECwwDTk9DMSQwIgYDVQQDDBtTSEFLRU4gQXBlaXJvbiBTeXN0ZW1zIDAxMkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATvM7voHFBJOHVADYo1gKHK2wU0lgJXrz%2FJPCTX0AKFzXtzFfZF0xEGkCL70IlycCO4vTtcloIGOIJTSN8ZSPxko4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDAxMkowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBR8cFoye3ncdoDrRITGSbjMkTRaXjCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpowRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIDMq05u%2FmxTkcvQ87CK05IwZn7uqSiXsm1aSRllYmJ1RAiEA7EUWXZR6EhCw5fZUc4jTEc8U6FjOvgibeN%2FVQ8eyXqw%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 012J' |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
