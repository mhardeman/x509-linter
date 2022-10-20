package url

type redirect struct{}

func init() {
	RegisterRule(&LintUrlRule{
		Code:        "e_atis_redirect",
		Description: "The STI-VS shall not follow HTTP redirections",
		Source:      Atis1000074Source,
		Rule:        NewRedirect,
	})
}

func NewRedirect() LintUrlRuleInterface {
	return &redirect{}
}

// CheckApplies implements LintUrlRuleInterface
func (*redirect) CheckApplies(data *LintUrlTestData) bool {
	return false
}

// Execute implements lint.LintUrlRuleInterface
func (*redirect) Execute(data *LintUrlTestData) *LintUrlResult {
	// The real test see in LintUrl method
	return &LintUrlResult{
		Status: Pass,
	}
}
