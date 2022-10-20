package url

import (
	"net/url"
	"strings"
)

type urlComponent struct{}

func init() {
	RegisterRule(&LintUrlRule{
		Code:        "e_atis_url_component",
		Description: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
		Source:      Atis1000074Source,
		Rule:        NewUrlComponent,
	})
}

func NewUrlComponent() LintUrlRuleInterface {
	return &urlComponent{}
}

// CheckApplies implements LintUrlRuleInterface
func (*urlComponent) CheckApplies(data *LintUrlTestData) bool {
	_, err := url.Parse(data.Url)
	return err != nil
}

// Execute implements lint.LintUrlRuleInterface
func (*urlComponent) Execute(data *LintUrlTestData) *LintUrlResult {
	url, _ := url.Parse(data.Url)
	if len(url.Fragment) > 0 || strings.HasSuffix(data.Url, "#") ||
		len(url.RawQuery) > 0 || strings.HasSuffix(data.Url, "?") ||
		url.User != nil {
		return &LintUrlResult{
			Status:  Error,
			Details: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
		}
	}

	return &LintUrlResult{
		Status: Pass,
	}
}
