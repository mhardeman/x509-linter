# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate b7dff830688c5a12bbacbf5be73931bee2efdfc1
Tested At: 2022-10-05 17:17:25 +0000 UTC\
Subject: CN=SHAKEN Drop Inc 258K, OU=Drop, O=Drop Inc, ST=Illinois, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: http://sti.comsapi.com/258k/ca.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDcjCCAxmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkSdgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDkxMzIwNDk1OVoXDTIyMTAxMzIwNDk1OVowYTELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQKDAhEcm9wIEluYzENMAsGA1UECwwERHJvcDEdMBsGA1UEAwwUU0hBS0VOIERyb3AgSW5jIDI1OEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATOhFkDpXxZUkcO1A7tlPq22zAs8qdt3ksBa096kRLfKB0tr5bmzJLlU6etipasInle1oSQ%2F39yJYnT2kWR2rjpo4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDI1OEswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRbEKtMiW3MGNu6af26wikm1g6C8TCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpowRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIB04%2BUqPx6qESIvBxiusrfdoflYXmLbwbu0ysgiqLBQeAiBt89aFXdLW4l6cGo98RdETLMNK4pL%2B0XpRDsxNoEGEcw%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 258K' |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
