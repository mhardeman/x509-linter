# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate 4d38ef7f05f03f1eb0037fa1c987c3f254e47046
Tested At: 2022-10-05 17:17:16 +0000 UTC\
Subject: CN=SHAKEN 459J, OU=SHAKEN, O=Altaworx, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.clearip.com/d3d6dbfc-2914-49c7-8f47-d0aa5bd5d764/8cb0dce87d05b9515483c81f65663c96.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6jCCApCgAwIBAgIQWZCX3SXWD6ZYPm7j2pUYTDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcyMDE4NTlaFw0yMjA5MjQyMDE4NThaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhBbHRhd29yeDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNDU5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPJqt7z2CicUQ9FWwneHj5SLhwHJ%2FSAmPJmHMzVcp%2BzlZSQafQzcddENDkTshtwNxzvEAPAfaygHQJO5DC2rtNKjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUcuUllGK%2Fv9T%2BM%2BcC1agltFEt9aowHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENDU5SjAKBggqhkjOPQQDAgNIADBFAiEA%2B%2Bb0RXBTdrsPqVyEjo35xZdkY2MKCoVEbxiWbkh6qX8CIAJvi3JqiXRqGjOCT1DnAXwq144Kk8OFKmp8N%2F%2F8P7p8)


| Code | Type | Details |
|------|------|---------|
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
