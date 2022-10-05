# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate c442c98ee4491a24afe84a5cd9f6cb796b0b13fd
Tested At: 2022-10-05 17:17:26 +0000 UTC\
Subject: CN=SHAKEN 148K, OU=SHAKEN, O=Orange County Rural Electric Membership Coop, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.transnexus.com/148K/340eb5b1-b461-4f5e-81f7-6da527199eb9.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDDTCCArSgAwIBAgIQZf2fHIdoflcF%2BAljLpMrDzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MDExNTMwNDlaFw0yMjA5MjkxNTMwNDhaMGsxCzAJBgNVBAYTAlVTMTUwMwYDVQQKEyxPcmFuZ2UgQ291bnR5IFJ1cmFsIEVsZWN0cmljIE1lbWJlcnNoaXAgQ29vcDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMTQ4SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPF%2BR0eiP%2Bw4JVRRz%2BmZJG6WqzNsmsnY8U5LD9xJRDp88tQ8txLfG%2FUnVIuLCmPz6kk5PC7QjXguZvW0kpSPORSjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUCab19HQtEStJySrqBcZMW74TGfwwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMTQ4SzAKBggqhkjOPQQDAgNHADBEAiBaEBLHM7JRTDcPV5yVFbpif0s9RtqAxVXtocvBmU%2ByvQIgPcPDZAGknZFnvwZdda2LX4KFkuyUIyHP7WCXPWtvE6c%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
