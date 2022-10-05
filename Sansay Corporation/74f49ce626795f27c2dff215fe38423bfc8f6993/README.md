# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 74f49ce626795f27c2dff215fe38423bfc8f6993
Tested At: 2022-10-05 17:17:20 +0000 UTC\
Subject: CN=SHAKEN Broadband Dynamics LLC 583j, OU=Network Operations, O=Broadband Dynamics LLC, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/583j_BROADBAND_DYNAMICS_STIR_SHAKEN.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDnTCCA0OgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTNswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE1MTcxOFoXDTIyMTEwNTE1MTcxOFowgYoxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdBcml6b25hMR8wHQYDVQQKDBZCcm9hZGJhbmQgRHluYW1pY3MgTExDMRswGQYDVQQLDBJOZXR3b3JrIE9wZXJhdGlvbnMxKzApBgNVBAMMIlNIQUtFTiBCcm9hZGJhbmQgRHluYW1pY3MgTExDIDU4M2owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQa%2BkrvsTqc2Zvkhvf7Rq0FzWu85RTSP8drlJdy%2FFz%2FWI1pCkOHGvoK7xsPDhTZZB0avi892aK02iucqxhnSq2No4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDU4M2owFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSJgdEFQOWqTxHwuwLUlpA%2Bh1bm3zCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpwwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCutZNKXx5h%2Bc4L60zF7AHrycwO86kdzXrzPuIe3ylW2AIgCOXSmp2M91IYXFbhb%2By3wNxzq4EwgbRWlrUQOC7LUxU%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 583j' |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
