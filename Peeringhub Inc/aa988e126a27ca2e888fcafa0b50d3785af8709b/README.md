# STIR/SHAKEN CA Ecosystem Compliance
## Peeringhub Inc

### Certificate aa988e126a27ca2e888fcafa0b50d3785af8709b
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: CN=ATI SHAKEN 731J, O=ATI, L=Phoenix, ST=AZ, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US

Link: https://certificates.peeringhub.io/731J/731J.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDCzCCArKgAwIBAgIQNyvUUal5R9gT%2F%2BFIQ2BJhDAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjA3MjkxNjQxNTZaFw0yMzA3MjkxNjQxNTZaMFQxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJBWjEQMA4GA1UEBwwHUGhvZW5peDEMMAoGA1UECgwDQVRJMRgwFgYDVQQDDA9BVEkgU0hBS0VOIDczMUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASl2XWD4%2FIjAfLLR7jmPtKWLMtDt5k0t%2BCzY0gHfiT3cbzYAutw1mgokOtKAiSt0RbTLUSvnu0K38F9zlprNh2io4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFH0dqeL3vPPVM1XgzSotBdU0l2rOMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYENzMxSjCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDRwAwRAIge24YdW%2Bsat%2BHlz4hNCajkVnrBCVBRX4hRLzcbKOpEwUCIBcL%2BVrHrJ5I0W3xDig6bdC%2FuWkmHtTKMtSp%2FZ46Zemy)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
