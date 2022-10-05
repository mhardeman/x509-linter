# SHAKEN COMPLIANCE
## Certificate

### Certificate 1df53e63631fcf7fa54e2005c65a0e250074bc80
Tested At: 2022-10-05 08:12:00 +0000 UTC

Subject: CN=SHAKEN Momentum Telecom 1417, OU=NOC, O=Momentum Telecom, ST=Georgia, C=US

Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/MomentumTelecom_1417

View: [Click to view](https://understandingwebpki.com/?cert=MIIDgTCCAyegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkSbMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDkxMzEzMTA1M1oXDTIyMTAxMzEzMTA1M1owbzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExGTAXBgNVBAoMEE1vbWVudHVtIFRlbGVjb20xDDAKBgNVBAsMA05PQzElMCMGA1UEAwwcU0hBS0VOIE1vbWVudHVtIFRlbGVjb20gMTQxNzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABII7HUTol1f%2FqSE%2B%2Faw3uBpttsdkbcNLNe4GTT6QNc1p5xhfUjCm8EJ1xtaeemvakN4DOVzQkrxtJiYyjdmnnLSjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEMTQxNzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFPm%2B%2BVJ2uk1zdbhZtZQkLpnC7wxkMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BmjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgaqnVZbslwHZy1%2FvpSIH2pU5zaSRh2RcmDzfZjl9q6hICIQDYS8P4mI7iqFjhS2cHbZLVZcKjIt8j%2FGEQ9h9DuvEP4Q%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 1417' |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

* The percent of certificates per issuer is calculated against total certificates from all issuers
** The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
