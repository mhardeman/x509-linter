package url_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/url"
)

func TestLintUrl_(t *testing.T) {
	res := url.LintUrl("https://app.connexcs.com/api/stir-shaken/cert/41.crt")

	fmt.Println(res)

	want := &url.LintUrlResult{
		Status:  url.Error,
		Details: "HTTP response shall have StatusCode 200, but it is 503 Service Unavailable",
	}
	if test := res.Results["e_http_status_200"]; !reflect.DeepEqual(test, want) {
		t.Errorf("httpStatus200.Execute() = %v, want %v", test, want)
	}
}
