# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate 46b11c828d17bfe5512f6e508a682133a5dc1cc4
Tested At: 2022-10-05 17:17:15 +0000 UTC\
Subject: CN=SHAKEN 366G, OU=SHAKEN, O=USA Digital, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.clearip.com/e9497931-96b3-4ff8-8306-1b1273847a4d/badc360632b171c32d2845cc958684df.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7TCCApOgAwIBAgIQcza75lOGAFxMBm5iTrX4EjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTQyMDE5NDZaFw0yMjA5MjEyMDE5NDVaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtVU0EgRGlnaXRhbDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzY2RzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHGRlquiQUTc%2BvVLqzIgnFgaSEW9ukB3L47YOcUUL9NTy8%2BviWfQrMuZvcscOYVxZzgVJL5%2F3rLkt%2Fa9s2RtUh2jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUathQG7Jq0UGzsP9CfSFa9orQdRYwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzY2RzAKBggqhkjOPQQDAgNIADBFAiEAoh32KRfeqHDxtL4jVI5W3w3zsZerccxDDHmbwCh1SA0CIDf9bj9RSq2Drlq90hihaCNVPdhOGr6HcTq3eFWjH%2Fsz)


| Code | Type | Details |
|------|------|---------|
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
