# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate adf6b3913ae512260a1102aabd66aa3d6183bcc0
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: CN=SHAKEN 345J, OU=SHAKEN, O=Ooma Inc, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.transnexus.com/345J/951b8406-ac81-4d2d-a6fd-78fa11b17cb8.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6zCCApCgAwIBAgIQbTVvJ8QksRwHdgOShAtYWTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTAwNDIzNDJaFw0yMjEwMTAwNDIzNDFaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhPb21hIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzQ1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPWYLm6fxXkvPXf6EvJ69csi9zyiAtlJ%2F67hTIFiqm0lYdyxSR4A%2BdXoctfID8YStZhkPnkw57SHO%2BC3BM90bFijggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUPn07IgYGk9pMHCwLKzmFMMsF73wwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzQ1SjAKBggqhkjOPQQDAgNJADBGAiEA34fTxWuRirfhBksnLzSc9ZZSjquA5CgMMWSleebLtQUCIQDrVHd1GTMu0ABty9NW%2B3MDBcOzDBr%2BQcT9LnN00EdJxQ%3D%3D)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
