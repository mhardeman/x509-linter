# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate b3f7e958393f9f7d7b454eb7d9b772cc860082b9
Tested At: 2022-10-05 17:17:24 +0000 UTC\
Subject: CN=SHAKEN Phone.com\, Inc. 633J, OU=voipteam, O=Phone.com\, Inc., ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Phonedotcom_633J

View: [Click to view](https://understandingwebpki.com/?cert=MIIDiDCCAy2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTMcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE0NDIzM1oXDTIyMTEwNDE0NDIzM1owdTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxGDAWBgNVBAoMD1Bob25lLmNvbSwgSW5jLjERMA8GA1UECwwIdm9pcHRlYW0xJDAiBgNVBAMMG1NIQUtFTiBQaG9uZS5jb20sIEluYy4gNjMzSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJTLc%2BXG1A5e1KyUnjizH7mo5f%2F%2B26dXea1HU0AeWM5RZHqKQFPEVumydDsK204FUayisWfuxa8XQh2AlfbaMV2jggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENjMzSjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFNa4pN3vX8RPZO26g9gl0niQX1%2FdMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BmjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAIjRqDS95oiaY9QFGoUpEsgA%2BaNwZT%2FmeKoCP7xfozBcAiEA4x1eFDoVffRFIGJ2FXMb15gZAHDKSJJCu%2B5kMkhpS1Q%3D)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 633J' |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
