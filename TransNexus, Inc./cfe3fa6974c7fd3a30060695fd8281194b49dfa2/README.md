# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate cfe3fa6974c7fd3a30060695fd8281194b49dfa2
Tested At: 2022-10-05 17:17:27 +0000 UTC\
Subject: CN=SHAKEN 159H, OU=SHAKEN, O=Edge Communications, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.transnexus.com/159H/7871f4a2-5aa5-4adf-96c6-cc4ff0163c87.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC9TCCApugAwIBAgIQbM4HkZ1j7PBft6yONJ7KTzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTIwOTEzMzZaFw0yMjA5MTkwOTEzMzVaMFIxCzAJBgNVBAYTAlVTMRwwGgYDVQQKExNFZGdlIENvbW11bmljYXRpb25zMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNTlIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEE86oMTbpVIRtrVEEr6kh9A5XfeOOnJSu4brQhl4aaFc%2FDnyl5Y7tCKuAIFAYOEGk7lZONRcwFje%2BRksjoKuxgqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBR3z9iLg792IDtwvQG5aG4GdZaRKjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNTlIMAoGCCqGSM49BAMCA0gAMEUCIQDZ7BnyxbshwH7AD06uXI5wSL84QZPmUPQLLFtt237%2BtwIgFXsQ3P6zHwd7xs%2Bn3tNSxbf8iEdDg12D29iaY1BISzs%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
