# SHAKEN COMPLIANCE
## Certificate

### Certificate 9ac0def10180cb22aad9e8885715c9409031697a
Tested At: 2022-10-05 08:40:22 +0000 UTC\
Subject: CN=cert.stir.t-mobile.com, OU=T-Mobile USA\, Inc, O=T-Mobile USA\, Inc., L=Bothell, ST=Washington, C=US\
Issuer: CN=TMOBILE-PROD-SUB-STIRSHAKEN-EC, O=TMOBILE-USA, C=US

Link: https://t-mobile-sticr.fosrvt.com/88a8e33055e725475530660e5d6c40d6adbe37ab7ae0ecc64b50205629548ae9.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIICyTCCAm%2BgAwIBAgIPAYH%2Bo0TdISlq5niitURDMAoGCCqGSM49BAMCMEwxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtUTU9CSUxFLVVTQTEnMCUGA1UEAxMeVE1PQklMRS1QUk9ELVNVQi1TVElSU0hBS0VOLUVDMB4XDTIyMDcxNDIxMDUyMVoXDTIzMDcxNDIxMzUyMVowgY4xCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdCb3RoZWxsMRswGQYDVQQKExJULU1vYmlsZSBVU0EsIEluYy4xGjAYBgNVBAsTEVQtTW9iaWxlIFVTQSwgSW5jMR8wHQYDVQQDExZjZXJ0LnN0aXIudC1tb2JpbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAER5y7ZAcpC5BgT4ElU782y9VvyzL5CSXR5dDg8QiLkbaMxuUJoFC6Yzdma2m4Jiz26UZhBA2AX1BDeGO29WDQZaOB8DCB7TAfBgNVHSMEGDAWgBQC2X9YhlA%2FeaB9iJtcgWc9pXwKeDAdBgNVHQ4EFgQUSnET9T3YyVih6PjLmbaCikNgEHswDgYDVR0PAQH%2FBAQDAgeAMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAhBgNVHREEGjAYghZjZXJ0LnN0aXIudC1tb2JpbGUuY29tMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENjUyOTAKBggqhkjOPQQDAgNIADBFAiBf03jL8qMfp5flDibYOien7k7E%2F9XlsSR0ReVxahdkWgIhAKvf9Zd8t%2BlIl6LdJ90f4MqUHtsHRQ%2FJVQ4Kty0IFM1h)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 6529' |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_basic_constraints | error | STI certificates shall contain a BasicConstraints extension marked critical |
| e_shaken_sti_extension_unknown | error | STI certificate shall not include extensions that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
