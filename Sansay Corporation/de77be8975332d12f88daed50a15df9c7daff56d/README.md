# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate de77be8975332d12f88daed50a15df9c7daff56d
Tested At: 2022-10-05 17:17:28 +0000 UTC\
Subject: CN=SHAKEN Matrix 7379, OU=Operations, O=Matrix, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Matrix-7379

View: [Click to view](https://understandingwebpki.com/?cert=MIIDczCCAxigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTNQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE0NDk1NVoXDTIyMTEwNDE0NDk1NVowYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZNYXRyaXgxEzARBgNVBAsMCk9wZXJhdGlvbnMxGzAZBgNVBAMMElNIQUtFTiBNYXRyaXggNzM3OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABE4GmAGYBCOAjL6rb9BWCSH32ogQfdkjAeHFQqwI%2BV0kBB1pSP%2FpqKSG6LrIf2SEMB8QZu3GjwfBz4B5ZR1PzpCjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENzM3OTAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHDiBbFd6ojGs8VgKYPyXeDa5mjjMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BmjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAKlkqzKPMAbjE86uVuVsjU%2FSuNEKug%2BnReDJLGhJuLKnAiEAwBELN0oF4OoEwYwSDvrEq%2BFV%2FSuKd4J1ujbScZl2Gls%3D)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 7379' |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
