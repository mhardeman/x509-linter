# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate 0579852fa1363af4515602dc413c7db91496d00d
Tested At: 2022-10-05 17:08:37 +0000 UTC\
Subject: CN=SHAKEN 738J, OU=SHAKEN, O=SIP.US LLC, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.clearip.com/fa6a5cd3-a6c2-4c98-9995-b87249365948/dd20f05f87b1e41476f21582bbafe1f3.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQdVM0TFS%2F0vfpec4r6dLh5zAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTEyMDE5NDJaFw0yMjA5MTgyMDE5NDFaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpTSVAuVVMgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzhKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEy0GPTrzgccdI%2ByIGAsal2ayMmMoDs6rMyyRHf4VbpU7vMMIRyijKZNWKteKK3q0Vh4KAmmwWbDD4XCJ8URJD0aOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTBXGEvON76EmQXh5Vuh0frBPHRGDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzhKMAoGCCqGSM49BAMCA0gAMEUCIQCr7gNNqoL17BnG574TwN3liy3GPsFi2TcF2BPFT46xLAIgRYMWI8Zvu7%2BxkolrAZzxYeN6saVCoOLPHYO3%2FpwN%2Fhc%3D)


| Code | Type | Details |
|------|------|---------|
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
