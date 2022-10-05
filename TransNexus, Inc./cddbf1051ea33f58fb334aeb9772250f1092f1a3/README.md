# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate cddbf1051ea33f58fb334aeb9772250f1092f1a3
Tested At: 2022-10-05 17:08:39 +0000 UTC\
Subject: CN=SHAKEN 177K, OU=SHAKEN, O=Cytracom\, LLC, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.clearip.com/8d118994-4994-4735-ac71-42c0bbb7848f/38b263d9e700daf7431b0dbfb4615808.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7zCCApWgAwIBAgIQeFakMifnjh%2BGxMBNmO3zYDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTYxNzM1NTBaFw0yMjA5MjMxNzM1NDlaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1DeXRyYWNvbSwgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNzdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOAaSJENa63oPzTIv%2FK9DuKeSQ8Obc3pb6I1a2QxJiQIUc1L1JMX4UW%2FHL4lgT0VJLi4SU3GBmDuMzvTj%2FBUC9qOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQeW1z2U0DmthLsGPTc8XfVNVGWQDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNzdLMAoGCCqGSM49BAMCA0gAMEUCIG0LcyBvve1NfnqjVFVbiIXc872LAKt6ChJ%2FiNWgYNhoAiEAmlIK%2FvWRvSaOJTZZdEgiePugwh21pSnlXzM3MrHFCFI%3D)


| Code | Type | Details |
|------|------|---------|
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
