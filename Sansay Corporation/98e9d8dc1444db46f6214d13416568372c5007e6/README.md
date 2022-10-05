# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 98e9d8dc1444db46f6214d13416568372c5007e6
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: emailAddress=stir-shaken@convoso.com, CN=SHAKEN Convoso 758J, OU=Infrastructure, O=Convoso, ST=California, C=US, emailAddress=stir-shaken@convoso.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://stirshaken.s3.us-west-1.amazonaws.com/stirshaken.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDpTCCA0ygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkSCYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDkwNjA2MDAyNFoXDTIyMTAxMTA2MDAyNFowgZMxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRAwDgYDVQQKDAdDb252b3NvMRcwFQYDVQQLDA5JbmZyYXN0cnVjdHVyZTEcMBoGA1UEAwwTU0hBS0VOIENvbnZvc28gNzU4SjEmMCQGCSqGSIb3DQEJARYXc3Rpci1zaGFrZW5AY29udm9zby5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQr0l5ZxKiYcoEQtP8DiZgZ22gPBvqCK41AW85shlZGPWjj6HD%2Ffq0GCeb9vaOdVm4VW%2FoTE03pP2jgHWjNgSfco4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDc1OEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT9SYAMUht0k6KTdYQSwd5BnmMWWTCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpowRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIDbp%2FmjVnQH8IoiiqFz5Ayhh4Keau5UmFVyw1JsEZwRsAiBKo4BXGM8urLIKjegb%2B56HNddomZh54savJOOJftj7JA%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 758J' |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
