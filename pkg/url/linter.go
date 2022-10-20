package url

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type LintUrlSource string

var (
	SystemSource      LintUrlSource = "System"
	HttpSource        LintUrlSource = "HTTP"
	Atis1000080Source LintUrlSource = "ATIS-1000080"
	Atis1000074Source LintUrlSource = "ATIS-1000074"
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

func GetRuleByName(name string) *LintUrlRule {
	return registry.Rules[name]
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
	Timestamp   time.Time
	Url         string
	StatusCode  int
	Body        []byte
	Results     map[string]*LintUrlResult
	HasErrors   bool
	HasWarnings bool
	HasNotices  bool
}

func (t *LintUrlResultSet) Append(code string, item *LintUrlResult) {
	if item != nil {
		t.Results[code] = item

		if item.Status == Error {
			t.HasErrors = true
		}
		if item.Status == Warn {
			t.HasWarnings = true
		}
		if item.Status == Notice {
			t.HasNotices = true
		}
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
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{}
	response, err := http.Get(url)
	if err != nil {
		// Disable SSL verification
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		var err2 error
		response, err2 = http.Get(url)
		if err2 == nil {
			result.Append("e_tls_transport", &LintUrlResult{
				Status:  Error,
				Details: err.Error(),
			})
		} else {
			err = err2
		}
	}
	testData := &LintUrlTestData{
		Url:      url,
		Response: response,
		Error:    err,
	}

	if response != nil {
		// test redirect
		if response.StatusCode == 200 {
			req, _ := http.NewRequest("GET", url, nil)
			if req != nil {
				client := new(http.Client)
				client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
					result.Append("e_atis_redirect", &LintUrlResult{
						Status:  Error,
						Details: "The STI-VS shall not follow HTTP redirections",
					})

					return errors.New("the STI-VS shall not follow HTTP redirections")
				}
				client.Do(req)
			}
		}

		// read response body
		result.StatusCode = response.StatusCode
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			testData.Error = err
		} else {
			testData.Body = body
			result.Body = body
		}
	} else {
		result.Append("e_bad_url", &LintUrlResult{
			Status:  Error,
			Details: err.Error(),
		})

		return result
	}

	for code, v := range registry.Rules {
		if v.Source == SystemSource {
			// skip System tests
			continue
		}

		rule := v.Rule()
		if rule.CheckApplies(testData) {
			result.Append(code, v.Rule().Execute(testData))
		}
	}

	return result
}
