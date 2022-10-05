# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 70b7083d423ab62a1d155811379bbd7e85b1a654
Tested At: 2022-10-05 17:17:19 +0000 UTC\
Subject: emailAddress=kelsey@qualityvoicedata.com, CN=SHAKEN Quality Voice & Data Inc. 548J, OU=Quality Voice & Data, O=Quality Voice & Data Inc., ST=Nebraska, C=US, emailAddress=kelsey@qualityvoicedata.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/548J/order/137_548J_67

View: [Click to view](https://understandingwebpki.com/?cert=MIID0zCCA3igAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkQ3IwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDgyMTE5MDYxMVoXDTIyMDkyMDE5MDYxMVowgb8xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZWJyYXNrYTEiMCAGA1UECgwZUXVhbGl0eSBWb2ljZSAmIERhdGEgSW5jLjEdMBsGA1UECwwUUXVhbGl0eSBWb2ljZSAmIERhdGExLjAsBgNVBAMMJVNIQUtFTiBRdWFsaXR5IFZvaWNlICYgRGF0YSBJbmMuIDU0OEoxKjAoBgkqhkiG9w0BCQEWG2tlbHNleUBxdWFsaXR5dm9pY2VkYXRhLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHxIpvtIpathfVQA%2Bf9h5YTuYtpurXcE4Fo%2BWfOrs4qTvs97AEUcl9TF5yLWG%2BWt%2FOsXbViKVjuykX3%2F5qk0fMSjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENTQ4SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFMreULfINsK1cltcvHQlGBGCmgS5MIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BmjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAJ%2FODMdyz837%2Bzw82ykGHfzycHW2X3%2FBaYeJKuCbfq8eAiEAmzKGPD%2BUTe%2Bf9gO9pSNIxooUu9CewY7FQLmvtLocu1w%3D)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 548J' |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
