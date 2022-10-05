# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus, Inc.

### Certificate cbef581071e3e24abfa76d3afbb9a5576560b271
Tested At: 2022-10-05 17:17:26 +0000 UTC\
Subject: CN=SHAKEN 738J, OU=SHAKEN, O=SkySwitch, C=US\
Issuer: CN=TransNexus\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\, Inc., C=US

Link: https://certificates.transnexus.com/738J/44bec728-126b-44d0-bb8e-353df6aa2735.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7DCCApGgAwIBAgIQab7j722NRIMCozRu5hpAszAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTIyMDE2MDVaFw0yMjA5MjEyMDE2MDRaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlTa3lTd2l0Y2gxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDczOEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQhoxag0dpvf%2BbdiEEpEsk2HAQYyQ6%2FnQtCkJzEjqxiwboH89WJVJ73t9cag8e%2BfU5%2Bb64%2BXz12lBEKoGjrrVxFo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFEt7qdTaNYoGFr5r%2BS%2BQgR1Q7qnMMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDczOEowCgYIKoZIzj0EAwIDSQAwRgIhAPgHJQxeS8K86zyYd6iYTtFyKri7YrCE2YZdbexnARQ8AiEAjpbJ1CwjzWCtlghexRUrGJc28zM2AmoihsvY5k4IHL0%3D)


| Code | Type | Details |
|------|------|---------|
| e_incorrect_ku_encoding | error | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
