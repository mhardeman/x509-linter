# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate 8671b256030780fce1ce8b26889a625fc2c4746b
Tested At: 2022-10-05 17:08:38 +0000 UTC\
Subject: CN=SHAKEN 366G, OU=SHAKEN, O=USA Digital, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.clearip.com/e9497931-96b3-4ff8-8306-1b1273847a4d/15d50fe5baa51bba2360cc46e2489e06.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7DCCApOgAwIBAgIQeIelNxPUEUsoIJrwgxATcTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTEyMDE5MzZaFw0yMjA5MTgyMDE5MzVaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtVU0EgRGlnaXRhbDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzY2RzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMToSg2EykVCOv6b0xUBaXPSULwohVpl6J8kmEziko6Tgp8a5Nqc9Pn8L0bquuD0vfnYow7I4L4rdaxSXwN3BA6jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUiZYlqxeEXs651MBUWREyzojVjHswHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzY2RzAKBggqhkjOPQQDAgNHADBEAiAbfIPC5%2BWyeFKguCECl%2BG%2BdpIsvYxUdUnh6SQODXTxygIgXUN8V%2BTF10lheW2oPYtZ1KVb42ysy4CL%2FiQ7dqGm%2F08%3D)


| Code | Type | Details |
|------|------|---------|
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
