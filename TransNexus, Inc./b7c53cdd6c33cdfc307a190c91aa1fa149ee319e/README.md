# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate b7c53cdd6c33cdfc307a190c91aa1fa149ee319e
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: CN=SHAKEN 093K, OU=SHAKEN, O=Skywave Wireless Inc, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.transnexus.com/093K/26bed0c9-863c-4e31-9972-8ca30a9ea44f.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC9TCCApygAwIBAgIQfI40twOP5ZyBCqX4QXCMRjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MDExNTM5NTdaFw0yMjA5MjkxNTM5NTZaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRTa3l3YXZlIFdpcmVsZXNzIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDkzSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCS6iWS0Iobp2F7%2FVyMEW4DCvL555p64gvjYJckebaWkDdNppqfO0E2yZLiqiq4srRVBBUeJPmmC2LAvYAWSnEqjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUXfAnXpnDgJ9Zhv6JjYELIrfF%2F68wHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDkzSzAKBggqhkjOPQQDAgNHADBEAiABiKeh34jzed0Cn7mt4VZUGfMdAhI8OoY%2BTRyrGEQ2dQIgO5qj6rR5va%2FiCFYeY7kuk9iZIqqHxq4gySqaNfyr2gY%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
