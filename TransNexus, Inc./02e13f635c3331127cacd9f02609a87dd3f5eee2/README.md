# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate 02e13f635c3331127cacd9f02609a87dd3f5eee2
Tested At: 2022-10-05 17:17:10 +0000 UTC\
Subject: CN=SHAKEN 297K, OU=SHAKEN, O=Clarity Voice, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.clearip.com/859299ca-4c86-4dfd-9785-6c758bee1b37/fc801558e11d85f018c7783bb92429e8.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7zCCApWgAwIBAgIQasR6AsAYWxLIsx1XsQGFRjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcxOTUyMzZaFw0yMjA5MjQxOTUyMzVaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1DbGFyaXR5IFZvaWNlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAyOTdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKrHK6vXqArQLTcfSpbpjc31Tv8TEK6c1BOuya0LE%2B3QUO0ENMAas5tTXrWhwo%2FvRvETjMXiEyjSUM0jMIaHvLqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBRqYO9OAeLdmmh1PA7JEFsBCOEPZDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQyOTdLMAoGCCqGSM49BAMCA0gAMEUCICkaPyUh15Bx%2BkLRAJ41uvOYX%2BmjErSDnLcktwmrqt1LAiEAsRHatp4NCW1aW5tyKdnq1WTXcAvF38J4mXUPQBTYZ4E%3D)


| Code | Type | Details |
|------|------|---------|
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
