package url

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type LintUrlSource string

var (
	HttpSource LintUrlSource = "HTTP"
)

type LintUrlRule struct {
	Code        string
	Description string
	Source      LintUrlSource
	Rule        func() LintUrlRuleInterface
}

type LintUrlRuleInterface interface {
	CheckApplies(data *LintUrlTestData) bool
	Execute(data *LintUrlTestData) *LintUrlResult
}

type LintUrlRuleSet = map[string]*LintUrlRule

type LintUrlRegistry struct {
	Rules LintUrlRuleSet
}

func RegisterRule(rule *LintUrlRule) {
	if rule != nil {
		if registry.Rules[rule.Code] != nil {
			panic(fmt.Sprintf("Cannot append %s the URL linter. This code is already in use", rule.Code))
		}
		registry.Rules[rule.Code] = rule
	}
}

var registry LintUrlRegistry = LintUrlRegistry{
	Rules: LintUrlRuleSet{},
}

type LintUrlStatus = int

var (
	Pass   LintUrlStatus = 0
	Error  LintUrlStatus = 1
	Warn   LintUrlStatus = 2
	Notice LintUrlStatus = 3
)

type LintUrlResult struct {
	Status  LintUrlStatus
	Details string
}

type LintUrlResultSet struct {
	Timestamp time.Time
	Url       string
	Results   map[string]*LintUrlResult
}

func (t *LintUrlResultSet) Append(code string, item *LintUrlResult) {
	if item != nil {
		t.Results[code] = item
	}
}

type LintUrlTestData struct {
	Url      string
	Response *http.Response
	Error    error
	Body     []byte
}

func (t *LintUrlTestData) HasResponse() bool {
	return t.Response != nil
}

func LintUrl(url string) *LintUrlResultSet {
	result := &LintUrlResultSet{
		Timestamp: time.Now(),
		Url:       url,
		Results:   map[string]*LintUrlResult{},
	}

	// send GET request
	response, err := http.Get(url)
	testData := &LintUrlTestData{
		Url:      url,
		Response: response,
		Error:    err,
	}
	if response != nil {
		// read response body
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			testData.Error = err
		} else {
			testData.Body = body
		}
	}

	for code, v := range registry.Rules {
		rule := v.Rule()
		if rule.CheckApplies(testData) {
			result.Results[code] = v.Rule().Execute(testData)
		}
	}

	return result
}
