# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 9e8edc670c03d9a05eac0d0db763525491d3c473
Tested At: 2022-10-05 17:17:23 +0000 UTC\
Subject: CN=SHAKEN Noble Systems Communications LLC 187J, OU=NOC, O=Noble Systems Communications LLC, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/NobleSys_187J

View: [Click to view](https://understandingwebpki.com/?cert=MIIDozCCA0igAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTMwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE0NDM1OFoXDTIyMTEwNDE0NDM1OFowgY8xCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdHZW9yZ2lhMSkwJwYDVQQKDCBOb2JsZSBTeXN0ZW1zIENvbW11bmljYXRpb25zIExMQzEMMAoGA1UECwwDTk9DMTUwMwYDVQQDDCxTSEFLRU4gTm9ibGUgU3lzdGVtcyBDb21tdW5pY2F0aW9ucyBMTEMgMTg3SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABP%2FmGlJuj%2FstdPV1ajhIPkNsLxwZQGX6idMO%2Fp8KZhhsDFAajP4jO2H2H5v6dRG3QkxY459oHsT7Ovh3AZPOq46jggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEMTg3SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFE4hswnCdx2RDd9F2IjWs4xprzuEMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BmjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhANOXtkHTlWrG%2FvI6qY7Fzx6x%2B%2F7WT3auuPrC7h4uTzt7AiEAqAFPXSE3G4XNuzAS8Iq9uhAAUlvXmnSyUn%2BMZm959a8%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 187J' |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
