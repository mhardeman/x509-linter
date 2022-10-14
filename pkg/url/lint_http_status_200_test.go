package url_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/peculiarventures/x509-linter/pkg/url"
)

func Test_httpStatus200_Execute(t *testing.T) {
	type args struct {
		data *url.LintUrlTestData
	}
	tests := []struct {
		name string
		args args
		want *url.LintUrlResult
	}{
		{
			name: "status 200",
			args: args{
				data: &url.LintUrlTestData{
					Response: &http.Response{
						StatusCode: 200,
					},
				},
			},
			want: &url.LintUrlResult{
				Status: url.Pass,
			},
		},
		{
			name: "status is not 200",
			args: args{
				data: &url.LintUrlTestData{
					Response: &http.Response{
						StatusCode: 500,
						Status:     "500 Internal",
					},
				},
			},
			want: &url.LintUrlResult{
				Status:  url.Error,
				Details: "HTTP response shall have StatusCode 200, but it is 500 Internal",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := url.NewHttpStatus200()
			if got := h.Execute(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpStatus200.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
