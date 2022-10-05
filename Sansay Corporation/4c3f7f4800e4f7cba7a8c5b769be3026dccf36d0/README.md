# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 4c3f7f4800e4f7cba7a8c5b769be3026dccf36d0
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: CN=SHAKEN Carrier One Inc. 705J, OU=Voice NOC, O=Carrier One Inc., ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Carrier_One_Inc._705J

View: [Click to view](https://understandingwebpki.com/?cert=MIIDiDCCAy2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkOGQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDUyNjIyMzMxMVoXDTIyMTAxMTIyMzMxMVowdTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGTAXBgNVBAoMEENhcnJpZXIgT25lIEluYy4xEjAQBgNVBAsMCVZvaWNlIE5PQzElMCMGA1UEAwwcU0hBS0VOIENhcnJpZXIgT25lIEluYy4gNzA1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABB424KM1Fpxx%2FW8QXNzarsuHIfoz1dq4yQZYg2%2FVKqjOnjgP8Ki2ySz0EnsRrod4mu82CD4vsUjdyvxwpjRzFN6jggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENzA1SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFH6eW3DcQiBGPpZUenPGyDP4EYEiMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BmjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAPsSMfZbnv4wMmYe3YgqnOJf%2FEnY49eL0mNBx85zYK9SAiEA7%2FNoNv3D9eRXXMyItFFB%2BUPMUgIWCA8gpfbnjn0TFvI%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 705J' |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
