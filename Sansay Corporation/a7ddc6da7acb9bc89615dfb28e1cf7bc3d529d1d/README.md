# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate a7ddc6da7acb9bc89615dfb28e1cf7bc3d529d1d
Tested At: 2022-10-05 17:17:23 +0000 UTC\
Subject: CN=SHAKEN NETRIO LLC 020K, OU=NOC, O=NETRIO LLC, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/NETRIO_LLC_020K

View: [Click to view](https://understandingwebpki.com/?cert=MIIDczCCAxmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTMgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE0NDI0OFoXDTIyMTEwNDE0NDI0OFowYTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRMwEQYDVQQKDApORVRSSU8gTExDMQwwCgYDVQQLDANOT0MxHzAdBgNVBAMMFlNIQUtFTiBORVRSSU8gTExDIDAyMEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATIwd%2FKwmHggCNfmSkcdHSrEQBSTDC5pLX1W6gA2%2BBf3sHszti9%2BYZTMf8XgJbHRLMylFri3hkPLYm6sSHCB%2FnUo4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDAyMEswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSRn6etTBT5QFC4Df9L0KjakmHqpzCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpowRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQC5dFqSCLaF%2FYxPeQ7Tl%2FTRUD8Y3ULRf4qlcil9TFER7gIgd96gOgI2DF4jmhkqV471QavWIpvg3%2BV8mQ6UvlfbC8M%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 020K' |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
