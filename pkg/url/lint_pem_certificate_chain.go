package url

import (
	"encoding/pem"

	"github.com/zmap/zcrypto/x509"
)

type pemCertificateChain struct{}

const pemCertificateChain_details = "ATIS-1000080 separately indicates that the mime type should be application/pem-certificate-chain"

func init() {
	RegisterRule(&LintUrlRule{
		Code:        "w_pem_certificate_chain",
		Description: pemCertificateChain_details,
		Source:      HttpSource,
		Rule:        NewPemCertificateChain,
	})
}

func NewPemCertificateChain() LintUrlRuleInterface {
	return &pemCertificateChain{}
}

// CheckApplies implements LintUrlRuleInterface
func (*pemCertificateChain) CheckApplies(data *LintUrlTestData) bool {
	return true
}

// Execute implements lint.LintUrlRuleInterface
func (*pemCertificateChain) Execute(data *LintUrlTestData) *LintUrlResult {
	var block *pem.Block
	rest := data.Body

	first := true
	for {
		block, rest = pem.Decode(rest)
		if block == nil {
			if first {
				return &LintUrlResult{
					Status:  Warn,
					Details: "HTTP response body should be PEM certificate chain. Response body is not PEM encoded",
				}
			}
			break
		}
		first = false

		// try parse the certificate
		_, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return &LintUrlResult{
				Status:  Warn,
				Details: "HTTP response body should be PEM certificate chain. PEM is not a certificate",
			}
		}
	}

	return &LintUrlResult{
		Status: Pass,
	}
}
