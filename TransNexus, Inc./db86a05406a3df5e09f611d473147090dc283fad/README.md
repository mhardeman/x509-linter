# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate db86a05406a3df5e09f611d473147090dc283fad
Tested At: 2022-10-05 17:08:39 +0000 UTC\
Subject: CN=SHAKEN 0172, OU=SHAKEN, O=CallTower, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.transnexus.com/0172/ada05b5e-0834-4c88-ba1f-4a8e4e907ba5.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6zCCApGgAwIBAgIQUf5UP5wfQ4hX0EgCSsyHUTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTEyMTMxMTZaFw0yMjA5MTgyMTMxMTVaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlDYWxsVG93ZXIxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDAxNzIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT%2F9zTXxOFv0bBs9QFq1lRZgJXOxpe1i5tHeVQGvyve0PRuWyXptC4Yq8EeenTduOBi%2BmNm84EjO6oFsnaJfVPYo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFFpvraZPHecqoLdJSmCdGvNOjoA8MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDAxNzIwCgYIKoZIzj0EAwIDSAAwRQIgXyg9IlrvYbBABepaluElGZDvCYk37fYXqxPFDmJvl30CIQCa0ylFLzydnShlJIDqBiCpD5Tw%2FSvoQLwLoEnHpG31rQ%3D%3D)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
