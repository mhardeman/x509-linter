# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 1cdf930f40715eef3d829d918e6fc7516c4af795
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: CN=SHAKEN Current Calls\, LLC 746J, OU=CurrentCalls, O=Current Calls\, LLC, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Current_Calls_LLC_746J

View: [Click to view](https://understandingwebpki.com/?cert=MIIDkTCCAzegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkPeUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDcyMjE1MzQ0N1oXDTIyMTAxMTE1MzQ0N1owfzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExGzAZBgNVBAoMEkN1cnJlbnQgQ2FsbHMsIExMQzEVMBMGA1UECwwMQ3VycmVudENhbGxzMScwJQYDVQQDDB5TSEFLRU4gQ3VycmVudCBDYWxscywgTExDIDc0NkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASCnQGHwC%2F7ZcJlA2%2F9NySkl3gZRbr3sReeCrF3k7%2B738HVJR34PYR1YI62ujddP2OrhGjFbg%2BrN8jkJAwgOBfjo4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDc0NkowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRL8zGGRUvjcZX4h8jAHMLlO%2BxTqzCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpowRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCxZHoh6briTy6EM%2BImr9X9TVkDujq%2F3k2%2BLSQ6smUE%2BwIgcZ%2F6rhyRqFXR8mOWqw8feW%2B1%2BgeaPgPqw%2FLGEwgK%2FhY%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 746J' |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
