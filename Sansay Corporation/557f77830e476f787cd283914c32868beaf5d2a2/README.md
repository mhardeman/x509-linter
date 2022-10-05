# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 557f77830e476f787cd283914c32868beaf5d2a2
Tested At: 2022-10-05 17:17:17 +0000 UTC\
Subject: CN=SHAKEN IPitomy 652J, OU=Operations, O=IPitomy, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/IPitomy_652J

View: [Click to view](https://understandingwebpki.com/?cert=MIIDdzCCAxygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTMYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE0NDIxOVoXDTIyMTEwNDE0NDIxOVowZDELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExEDAOBgNVBAoMB0lQaXRvbXkxEzARBgNVBAsMCk9wZXJhdGlvbnMxHDAaBgNVBAMME1NIQUtFTiBJUGl0b215IDY1MkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATya7C0e2%2F7028Zf07DXg%2BtZHQ5IwPYAr92QpZf1GcdmUZcSKaB3Su%2BTsDXww8AHKPLLNUXB6u5V1709i9AgLo9o4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDY1MkowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT9zpH7wrK0MJUL5lc43BYJ3qPeAjCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpowRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQDIQb6Dw4UHNDHFdZpsfS2%2BffScO81y7Dgufi27CO7ugAIhALiK9yzD9FIuBmZwYv9M1Ed%2FgMuGeAZKZlpTP8VjO44%2B)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 652J' |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
