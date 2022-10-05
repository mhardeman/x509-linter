# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate 136f6839b2fa440965940ff380d9cde7b053f2d0
Tested At: 2022-10-05 17:17:11 +0000 UTC\
Subject: CN=Fusion Connect SHAKEN 2720, OU=Fusion Connect, O=TransNexus, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://prod001-cr.rbbnidhub.com/8ZJdVFg7gz/2

View: [Click to view](https://understandingwebpki.com/?cert=MIICoTCCAkegAwIBAgIQSFscNJOjGceOfFyuBLk%2BVzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMTAeFw0yMjAzMzAxMjU0MTVaFw0yMzAzMzAxMjU0MTRaMGAxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpUcmFuc05leHVzMRcwFQYDVQQLEw5GdXNpb24gQ29ubmVjdDEjMCEGA1UEAxMaRnVzaW9uIENvbm5lY3QgU0hBS0VOIDI3MjAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT63jwuwao3XgseWltjK2ryDbERqdg1YPJq5ERkdas77TTVUuYvquN%2FKmHQd5tGNxd0HFT74V7uRnF%2B0%2Faql%2B59o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQgC9%2Fhoul3TB6S5BONMObenIn0jjAfBgNVHSMEGDAWgBSUhjk%2F5PWSoJ%2F%2F3Cd1GppG8HnhYjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBYGCCsGAQUFBwEaBAowCKAGFgQyNzIwMAoGCCqGSM49BAMCA0gAMEUCIDNwXtIcpzboL8SV5qUwPjA3OiDBDuLVOeuxTqtMfwpQAiEAt7LcSEYFhKVzHAXNrJ9oJPpw8%2Fzrflp1JEJ6FudrYKU%3D)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
