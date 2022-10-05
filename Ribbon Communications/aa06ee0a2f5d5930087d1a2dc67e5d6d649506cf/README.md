# STIR/SHAKEN CA Ecosystem Compliance
## Ribbon Communications

### Certificate aa06ee0a2f5d5930087d1a2dc67e5d6d649506cf
Tested At: 2022-10-05 17:17:23 +0000 UTC\
Subject: CN=Nuwave Communications SHAKEN 620J, OU=Nuwave Communications, O=Nuwave Communications, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US

Link: https://prod001-cr.rbbnidhub.com/Jx6yR-yMgz/620J202204-2c7d5c55a20834b031681dbd3e2eb9f0

View: [Click to view](https://understandingwebpki.com/?cert=MIIDETCCAragAwIBAgIQLH1cVaIINLAxaB29Pi658DAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwNDI5MTc0MzAwWhcNMjMwNDI5MTc0MjU5WjB5MQswCQYDVQQGEwJVUzEeMBwGA1UECgwVTnV3YXZlIENvbW11bmljYXRpb25zMR4wHAYDVQQLDBVOdXdhdmUgQ29tbXVuaWNhdGlvbnMxKjAoBgNVBAMMIU51d2F2ZSBDb21tdW5pY2F0aW9ucyBTSEFLRU4gNjIwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKQcN14DpTdbPoHaOy6cVJljOZPyQ%2BgyX9kIcftnhDG18ZcaWvrgXq21a6A4Y3vyMPMBhNFH4oB4qergCJqbgdmjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUqbMilZVs0XNrwwIFe%2BAsB%2Fa6tHowGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MjBKMAoGCCqGSM49BAMCA0kAMEYCIQCiqUKrea18gAnjqSzfXDL3r3Dv%2BE31IMIKw8ctF0nttwIhALOBc%2BpGfwh08qg%2Bf8qmg5wCPmMKr%2BKaWCLysLzlVer6)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_extension_unknown | error | STI certificate shall not include extensions that are not specified |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
