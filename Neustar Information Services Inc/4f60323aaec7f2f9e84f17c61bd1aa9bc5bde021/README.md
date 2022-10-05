# STIR/SHAKEN CA Ecosystem Compliance
## Neustar Information Services Inc

### Certificate 4f60323aaec7f2f9e84f17c61bd1aa9bc5bde021
Tested At: 2022-10-05 17:17:16 +0000 UTC\
Subject: CN=SHAKEN 553J, OU=VOIP, O=Whitesky Communications LLC, L=Northport, ST=AL, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

Link: https://certs.sip.wtsky.net/prod/shaken-6-2023.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIIDPDCCAuKgAwIBAgIURaI2P9Jsx7RD6C1kjmV5TWZsFJ8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYyMzE2MzYxOFoXDTIzMDYyMzE2MzYxOFoweTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkFMMRIwEAYDVQQHDAlOb3J0aHBvcnQxJDAiBgNVBAoMG1doaXRlc2t5IENvbW11bmljYXRpb25zIExMQzENMAsGA1UECwwEVk9JUDEUMBIGA1UEAwwLU0hBS0VOIDU1M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQZ5o0jeSry66%2BhV2UrYHIR4MCsjypnjQ3RiIQiDuTD92R0LkalQ%2BixgE1qYPglLbBHB%2BT5ae7LlQsG6fsd8S1So4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDU1M0owDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBRW12yTlqYszKS%2FeuugmMqnaHBD7DAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgU3TF0DWAEd7kuNXeexP0JO8bL3fDCb30lAQbLiUl7AECIQDSuknhR%2FXIZ9OVZVIUiYStu2phAQjA1U0mA8MIflv7XA%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_extension_unknown | error | STI certificate shall not include extensions that are not specified |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
