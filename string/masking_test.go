package strings

import (
	"fmt"
	"testing"

	"github.com/herurahmat/pkg/strings"
)

type testcase struct {
	name     string
	output   string
	expected string
}

type testStruct struct {
	Foo string `json:"foo" mask:"any"`
	Bar string
}

func TestMask(t *testing.T) {
	authorization := "Bearer eyJ0eXAiOiJKV1QiLCJraWQiOiJ3VTNpZklJYUxPVUFSZVJCL0ZHNmVNMVAxUU09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiJmMThhY2NmNS04MWIyLTRlNWQtOWRhZS1mNmQ0MTQzN2ZlMGIiLCJjdHMiOiJPQVVUSDJfU1RBVEVMRVNTX0dSQU5UIiwiYXV0aF9sZXZlbCI6MCwiYXVkaXRUcmFja2luZ0lkIjoiYzFjNDQ0ZjAtOGFiNS00NWYzLTg3ZWEtNmVmMjQyNWRhOTgwLTE5NjA0OSIsInN1Ym5hbWUiOiJmMThhY2NmNS04MWIyLTRlNWQtOWRhZS1mNmQ0MTQzN2ZlMGIiLCJpc3MiOiJodHRwczovL2FtOjQ0My9hbS9vYXV0aDIvdHNlbC9ob21lbHRlL3dlYiIsInRva2VuTmFtZSI6ImFjY2Vzc190b2tlbiIsInRva2VuX3R5cGUiOiJCZWFyZXIiLCJhdXRoR3JhbnRJZCI6InIyNmY5WkRrQ1VxR1pjNjN5cWR6cm5yUElZRSIsIm5vbmNlIjoidHJ1ZSIsImF1ZCI6InJnOGFiYzQzMDMyOTE2MTlqMWt3cTQ5MTc3M2JtMDM0IiwibmJmIjoxNjg5NjYxOTQxLCJncmFudF90eXBlIjoiYXV0aG9yaXphdGlvbl9jb2RlIiwic2NvcGUiOlsib3BlbmlkIiwicHJvZmlsZSJdLCJhdXRoX3RpbWUiOjE2ODk2NjE5NDAsInJlYWxtIjoiL3RzZWwvaG9tZWx0ZS93ZWIiLCJleHAiOjE2ODk2NjU1NDEsImlhdCI6MTY4OTY2MTk0MSwiZXhwaXJlc19pbiI6MzYwMCwianRpIjoiQUpYd3kzU0lXQnYxLUwyQ2FxakpyVkMybGVnIn0.gRWQ6ecZaj1qBK6JnBblKSnOreRAuXaICOHnfyTyadZ3JqyaotudjV0CwdXlANLDaiODcYkMdGy0bMGR1Jes2pxfMb4ktYsKh8D7TXZmKmbcr-jPI-_z0EeeCXYCuwygxZJb1pB_lAC8zZ49xfUxZyXbkLMdgFjgpvyyofu0-hBgJbNcGQNam8XsLY5Vwm1NOaRmQ8wlkYRDk24FHTkvhrGD72srFyK8qBmev08rRwqZYTzq_F7uQdgeqehcTLLwoIyfnAT9I8DGwRaCwBuk9x7i1vvNxZdVWUf8bQE6s5kVDDXvKWS1s-lR4_dj_Gvp02trRvbpKDwk3X6282fEqQ"
	testcases := []testcase{
		{"mask with no type", strings.Mask("", authorization), authorization},
		{"mask any", strings.Mask(strings.Any, authorization), "******"},
		{"mask any with empty string", strings.Mask(strings.Any, ""), ""},
		{"mask left", strings.Mask(strings.Left, authorization), "******82fEqQ"},
		{"mask left with three letters", strings.Mask(strings.Left, "abc"), "******c"},
		{"mask left with empty string", strings.Mask(strings.Left, ""), ""},
		{"mask middle", strings.Mask(strings.Middle, authorization), "***iwibmJ***"},
		{"mask middle with three letters", strings.Mask(strings.Middle, "abc"), "***b***"},
		{"mask middle with empty string", strings.Mask(strings.Middle, ""), ""},
		{"mask right", strings.Mask(strings.Right, authorization), "Bearer******"},
		{"mask right with three letters", strings.Mask(strings.Right, "abc"), "a******"},
		{"mask right with empty string", strings.Mask(strings.Right, ""), ""},
	}

	for _, testcase := range testcases {
		if testcase.output != testcase.expected {
			t.Errorf(testcase.name, testcase.expected, testcase.output)
		}
	}
}

func TestMaskStruct(t *testing.T) {
	payload := testStruct{"bar", "foo"}
	strings.MaskStruct(&payload)

	testcases := []testcase{
		{"mask struct", fmt.Sprintf("%+v", payload), "{Foo:****** Bar:foo}"},
	}

	for _, testcase := range testcases {
		if testcase.output != testcase.expected {
			t.Errorf(testcase.name, testcase.expected, testcase.output)
		}
	}
}
