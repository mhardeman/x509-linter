# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 8fb71997a13c7eff35c8cd8ce93595e8c9e45924
Tested At: 2022-10-05 17:17:21 +0000 UTC\
Subject: CN=SHAKEN Asia Pacific Network 988J, OU=APN, O=Asia Pacific Network, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Asia_Pacific_Network_988J

View: [Click to view](https://understandingwebpki.com/?cert=MIIDhjCCAy2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTEEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMDAwMDAwMFoXDTIzMTAxMDAwMDAwMFowdTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMR0wGwYDVQQKDBRBc2lhIFBhY2lmaWMgTmV0d29yazEMMAoGA1UECwwDQVBOMSkwJwYDVQQDDCBTSEFLRU4gQXNpYSBQYWNpZmljIE5ldHdvcmsgOTg4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGprIYV6mubDNzZG2RPbgooSHOFxVCgV8IjCNPqJmSD47y2kcMzY2zCWJ40QKNL%2B4YYSUYw7WblAWG9N%2B%2FAMtq2jggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEOTg4SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFNol0GQg%2FmmWcZdO0zomI8Pog0xRMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BmjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgAPZMCkIxr5gmCKfXhnFX8rO4M6lrIGsPrJD%2BlaIDN1oCIG0AYoYwRtSpfIbV3kq8RIV06d5QvGa5HKOLEtSJRC75)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 988J' |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
