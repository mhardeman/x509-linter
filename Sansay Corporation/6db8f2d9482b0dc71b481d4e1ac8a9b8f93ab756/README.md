# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 6db8f2d9482b0dc71b481d4e1ac8a9b8f93ab756
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: CN=SHAKEN Touchtone 683A, OU=NOC, O=Touchtone, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Touchtone_683A

View: [Click to view](https://understandingwebpki.com/?cert=MIIDdTCCAxygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTMMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE0NDAxNVoXDTIyMTEwNDE0NDAxNVowZDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxEjAQBgNVBAoMCVRvdWNodG9uZTEMMAoGA1UECwwDTk9DMR4wHAYDVQQDDBVTSEFLRU4gVG91Y2h0b25lIDY4M0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATwPwhwR259HmHeCt9A84ohVtuEKJYTrlFMI5HXCls08mFEPsyKfoW%2BySkQmaGCaCmCEHCD6rRfhou6X8dWmxB4o4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDY4M0EwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTh8%2FsJaKJ4cvdArhinM9tZbqvyGTCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpowRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIFZCurmFQ8NLJsUVS1yJtj5tG%2BHmtfoaN5GPv4wa10WzAiBlTPvk53xRat006cV17KI4%2BjtOAEw5Gn%2F55cY9cgUj9w%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 683A' |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
