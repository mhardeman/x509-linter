# STIR/SHAKEN CA Ecosystem Compliance
## Sansay Corporation

### Certificate 0b5745860ea6ca082226e5e2890ee04bfce1a926
Tested At: 2022-10-05 17:17:11 +0000 UTC\
Subject: CN=SHAKEN VoIP Innovations 597F, OU=NOC, O=VoIP Innovations, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/VoIP_Innovations_597F

View: [Click to view](https://understandingwebpki.com/?cert=MIIDhjCCAyygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTMowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE0NDMyM1oXDTIyMTEwNDE0NDMyM1owdDELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEZMBcGA1UECgwQVm9JUCBJbm5vdmF0aW9uczEMMAoGA1UECwwDTk9DMSUwIwYDVQQDDBxTSEFLRU4gVm9JUCBJbm5vdmF0aW9ucyA1OTdGMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEZxb5rNCMqQTKIowEPxI78IPaRDMoC1gJoQtw%2FiJ%2BjTAuI2%2F4nDDGBz7ACAo3eIerjz3WANlpeFOQbeXW03U8jaOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ1OTdGMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU4p9sS%2FoiV03eqqex3c2k%2BTYL0%2FAwgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6aMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA5qrQ2PI9OTrZI3nCmkzwAC0LJHhebfyovBtz37XAMeYCIDSCUphqDYO4S7sA5I0%2BY7gsP7bOo6nhl5bMiBw8%2BhkY)


| Code | Type | Details |
|------|------|---------|
| e_shaken_sti_certificate_policies | error | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_shaken_cp1_3_subject_sn | error | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_shaken_cp1_3_ambiguous_identifier | error | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_shaken_sti_subject_cn | error | Common name shall contain the text string 'SHAKEN 597F' |
| w_shaken_sti_subject_rdn_unknown | warn | STI certificate shall not include RDNs that are not specified |

\* The percent of certificates per issuer is calculated against total certificates from all issuers\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer
\*\*\* Tests do not report on certificates with issues that predate the currently required ATIS 1000080 and Certificate Policy versions
