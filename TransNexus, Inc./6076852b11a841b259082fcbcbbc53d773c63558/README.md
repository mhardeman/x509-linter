# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate 6076852b11a841b259082fcbcbbc53d773c63558
Tested At: 2022-10-05 17:17:18 +0000 UTC\
Subject: CN=SHAKEN 9714, OU=SHAKEN, O=Grid4 Communications, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.clearip.com/51a37c7a-5af2-439d-94ce-677fa750ee2f/7f824bec656dcd7a0d69cbbf76421f73.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC9zCCApygAwIBAgIQYqfwkOc%2Bq4e1tBPrSvSa4zAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTQxODM1NTlaFw0yMjA5MjExODM1NThaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRHcmlkNCBDb21tdW5pY2F0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gOTcxNDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKaWNIeD5RZcsu%2BVB7CQRjVNJ2fPyBaj4U6lHKUL3aAqNLlMFrNcuVV1iGEd%2BF7DknB%2FeBLkgnTH7ilwkRFSBWejggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUF%2FI6PeLnsvt5snF%2FB%2Fk5f2PK8fowHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEOTcxNDAKBggqhkjOPQQDAgNJADBGAiEAqkceW0dMqLeAdtpQKaeNjCZkr%2FG1OsRmQ4GzLYae%2BNICIQDuhEip9HCpeP4YIC%2B7G4xt3F7sqnsad%2FUR%2FHZUG3T21A%3D%3D)


| Code | Type | Details |
|------|------|---------|
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
