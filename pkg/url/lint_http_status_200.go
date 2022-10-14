package url

import "fmt"

type httpStatus200 struct{}

const httpStatus200_details = "HTTP response shall have StatusCode 200"

func init() {
	RegisterRule(&LintUrlRule{
		Code:        "e_http_status_200",
		Description: httpStatus200_details,
		Source:      HttpSource,
		Rule:        NewHttpStatus200,
	})
}

func NewHttpStatus200() LintUrlRuleInterface {
	return &httpStatus200{}
}

// CheckApplies implements LintUrlRuleInterface
func (*httpStatus200) CheckApplies(data *LintUrlTestData) bool {
	return data.HasResponse()
}

// Execute implements lint.LintUrlRuleInterface
func (*httpStatus200) Execute(data *LintUrlTestData) *LintUrlResult {
	if data.Response.StatusCode != 200 {
		return &LintUrlResult{
			Status:  Error,
			Details: fmt.Sprintf("%s, but it is %s", httpStatus200_details, data.Response.Status),
		}
	}

	return &LintUrlResult{
		Status: Pass,
	}
}
