package url

import (
	"strings"
)

type atisHeaderContentType struct{}

func init() {
	RegisterRule(&LintUrlRule{
		Code:        "w_atis_content_type",
		Description: "ATIS-1000080 separately indicates that the mime type should be application/pem-certificate-chain",
		Source:      Atis1000080Source,
		Rule:        NewAtisHeaderContentType,
	})
}

func NewAtisHeaderContentType() LintUrlRuleInterface {
	return &atisHeaderContentType{}
}

// CheckApplies implements LintUrlRuleInterface
func (*atisHeaderContentType) CheckApplies(data *LintUrlTestData) bool {
	return true
}

// Execute implements lint.LintUrlRuleInterface
func (*atisHeaderContentType) Execute(data *LintUrlTestData) *LintUrlResult {
	contentType := data.Response.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/pem-certificate-chain") {
		return &LintUrlResult{
			Status:  Warn,
			Details: "HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain",
		}
	}

	return &LintUrlResult{
		Status: Pass,
	}
}
