# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 17599bafbe0c44b8ec89c36eea809196b551d725
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: CN=SHAKEN Magna5\, LLC 3849, OU=NOC, O=Magna5\, LLC, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Magna5_3849.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDdTCCAxugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTOMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE1MzIwM1oXDTIyMTEwNTE1MzIwM1owYzELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAtNYWduYTUsIExMQzEMMAoGA1UECwwDTk9DMSAwHgYDVQQDDBdTSEFLRU4gTWFnbmE1LCBMTEMgMzg0OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHSmEvk0ocbvzVCoSD1ktVCycyAFmG%2BqFYMhuhqOwFgUfujidBw3leL%2Fu8Np6stCzm3%2BzRIpi9AhU24BgtokQuGjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEMzg0OTAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFNQ1JL%2F6yQTbEEZT2L%2BuS97RFERAMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAMqr7xOQaQKZ7JHK%2BeWu2I7TZWVGHQo8ORh6%2FoncpBoCAiABHcBA2i3lqDhMl8lqeLQi%2FD4XirUsazPeuj9R05MmoA%3D%3D)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 3849' |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
