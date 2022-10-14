package url

type atisProtocol struct{}

func init() {
	RegisterRule(&LintUrlRule{
		Code:        "w_atis_protocol",
		Description: "The verifier should not dereference any protocol other than https or a port other than 443 or 8443",
		Source:      Atis1000080Source,
		Rule:        NewAtisProtocol,
	})
}

func NewAtisProtocol() LintUrlRuleInterface {
	return &atisProtocol{}
}

// CheckApplies implements LintUrlRuleInterface
func (*atisProtocol) CheckApplies(data *LintUrlTestData) bool {
	return true
}

// Execute implements lint.LintUrlRuleInterface
func (*atisProtocol) Execute(data *LintUrlTestData) *LintUrlResult {
	port := data.Response.Request.URL.Port()
	if !(data.Response.Request.URL.Scheme == "https" &&
		(port == "" || port == "443" || port == "8443")) {
		return &LintUrlResult{
			Status:  Warn,
			Details: "The verifier should not dereference any protocol other than https or a port other than 443 or 8443",
		}
	}

	return &LintUrlResult{
		Status: Pass,
	}
}
