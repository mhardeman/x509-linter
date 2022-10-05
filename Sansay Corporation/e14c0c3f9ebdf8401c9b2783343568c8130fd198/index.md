# SHAKEN COMPLIANCE
## Certificate

### Certificate e14c0c3f9ebdf8401c9b2783343568c8130fd198
Tested At: 2022-10-05 08:12:04 +0000 UTC

Subject: CN=SHAKEN Broadband Dynamics LLC 583j, OU=Network Operations, O=Broadband Dynamics LLC, ST=Arizona, C=US

Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/583j_BROADBAND_DYNAMICS_STIR_SHAKEN.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDnTCCA0OgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkScAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDkxMzEzMjczM1oXDTIyMTAxNDEzMjczM1owgYoxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdBcml6b25hMR8wHQYDVQQKDBZCcm9hZGJhbmQgRHluYW1pY3MgTExDMRswGQYDVQQLDBJOZXR3b3JrIE9wZXJhdGlvbnMxKzApBgNVBAMMIlNIQUtFTiBCcm9hZGJhbmQgRHluYW1pY3MgTExDIDU4M2owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQa%2BkrvsTqc2Zvkhvf7Rq0FzWu85RTSP8drlJdy%2FFz%2FWI1pCkOHGvoK7xsPDhTZZB0avi892aK02iucqxhnSq2No4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDU4M2owFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSJgdEFQOWqTxHwuwLUlpA%2Bh1bm3zCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpowRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQC5n8wgM2QTUXdk7bM%2B1JOSDU%2ByVALwqqwu9URQZKYUsQIgLV5y00EuqUESqn4rHGcrmc8kn%2FsOyDQKUDYnSdd53N0%3D)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 583j' |

* The percent of certificates per issuer is calculated against total certificates from all issuers
** The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
