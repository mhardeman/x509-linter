package url

func init() {
	RegisterRule(&LintUrlRule{
		Code:        "e_bad_url",
		Source:      SystemSource,
		Description: "Any errors on GET request sending",
		Rule:        NewNothing,
	})
	RegisterRule(&LintUrlRule{
		Code:        "e_tls_transport",
		Source:      SystemSource,
		Description: "TLS problem on link loading",
		Rule:        NewNothing,
	})
}

type nothing struct{}

func NewNothing() LintUrlRuleInterface {
	return &nothing{}
}

// CheckApplies implements LintUrlRuleInterface
func (*nothing) CheckApplies(data *LintUrlTestData) bool {
	return false
}

// Execute implements lint.LintUrlRuleInterface
func (*nothing) Execute(data *LintUrlTestData) *LintUrlResult {
	return &LintUrlResult{
		Status: Pass,
	}
}
