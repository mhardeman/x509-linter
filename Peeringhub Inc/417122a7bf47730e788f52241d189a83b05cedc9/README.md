# STIR/SHAKEN CA Ecosystem Compliance
## Peeringhub Inc

### Certificate 417122a7bf47730e788f52241d189a83b05cedc9
Tested At: 2022-10-05 17:17:15 +0000 UTC\
Subject: CN=Teleinx SHAKEN 744J, O=Teleinx LLC, L=Chicago, ST=IL, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US

Link: https://certificates.peeringhub.io/744J/744J.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDGTCCAr6gAwIBAgIQLj028KcpXmDFbWAuDEdOyDAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjA4MjYyMzMxMzBaFw0yMzA4MjYyMzMxMzBaMGAxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJJTDEQMA4GA1UEBwwHQ2hpY2FnbzEUMBIGA1UECgwLVGVsZWlueCBMTEMxHDAaBgNVBAMME1RlbGVpbnggU0hBS0VOIDc0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATG%2BFyfEfJrX4ic%2BCwvxuXBpGYz3Hg3jcdNONw3XyI3vENLsxFe%2FPA7sykjNKmSF2eH7p7XFF0JxZerFnFyIzz8o4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFCWCjv3puPEBWuzbdYxa%2FavObQ%2BCMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYENzQ0SjCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDSQAwRgIhALdOTyAl%2FGe2JejEyp1tiNSv81kSbKRcjn20MJbfOQr9AiEAnayQzCoGtdNRt0fX0%2BV%2Bl6AYhzEKt1mWPBasFLeH5m8%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
