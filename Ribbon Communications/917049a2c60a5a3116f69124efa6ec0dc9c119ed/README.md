# STIR/SHAKEN CA Ecosystem Compliance
## Ribbon Communications

### Certificate 917049a2c60a5a3116f69124efa6ec0dc9c119ed
Tested At: 2022-10-05 17:17:21 +0000 UTC\
Subject: CN=Netfortris SHAKEN 8886, OU=VOIP, O=Netfortris, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US

Link: https://prod001-prod011-cr.rbbnidhub.com/t0CiOIjnRz/NFJune102022-26092b2abd3c6bbfc676d98623fc1b25

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6DCCAo%2BgAwIBAgIQJgkrKr08a7%2FGdtmGI%2FwbJTAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwNjEwMTkwMDA5WhcNMjMwNjEwMTkwMDA4WjBSMQswCQYDVQQGEwJVUzETMBEGA1UECgwKTmV0Zm9ydHJpczENMAsGA1UECwwEVk9JUDEfMB0GA1UEAwwWTmV0Zm9ydHJpcyBTSEFLRU4gODg4NjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHKjVX0y5iWU%2BPWcVma0eo8bOVa6szzU%2BEyknh5%2BvU%2FrxREY8K9gnIQZwzAQC%2Fobbdb9g467rPWSJnwYFydtfCCjggEjMIIBHzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUaaPiQZSfxkkc%2F0jrtzilcXSuLVUwGQYDVR0gBBIwEDAOBgpghkgBhv8JAQEBMAAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMEMGCCsGAQUFBwEBBDcwNTAzBggrBgEFBQcwAoYnaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9pY2EuY3J0MB8GA1UdIwQYMBaAFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMBYGCCsGAQUFBwEaBAowCKAGFgQ4ODg2MAoGCCqGSM49BAMCA0cAMEQCIAiL%2FtMGhx5ilLmO11px22GSd%2FRUvCx90DyXF7PHiMZzAiBR436HM6vmoXCoSf%2FSUpihGhcjv0LwC4Gaz4xQH7d4eQ%3D%3D)


| Code | Type | Details |
|------|------|---------|
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |
| e_shaken_sti_extension_unknown | error | STI certificate shall not include extensions that are not specified |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
