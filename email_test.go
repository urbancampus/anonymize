package anonymize_test

import (
	"strings"
	"testing"

	"github.com/emmanuelay/anonymize"
)

func TestAnonymizeEmail(t *testing.T) {
	tests := []struct {
		got  string
		want string
	}{
		{
			got:  "emmanuel@somedomain.com",
			want: "e*******@s*********.com",
		},
		{
			got:  "this.is.incorrect@domain@domain.com",
			want: "t***.i*.i********@d*****@d*****.com",
		},
		{
			got:  "e@f.com",
			want: "*@*.com",
		},
		{
			got:  "@do.com",
			want: "@d*.com",
		},
		{
			got:  "@.",
			want: "@.",
		},
		{
			got:  "@",
			want: "@",
		},
		{
			got:  ".",
			want: ".",
		},
		{
			got:  "a@domain.com",
			want: "*@d*****.com",
		},
		{
			got:  "al@domain.com",
			want: "a*@d*****.com",
		},
		{
			got:  "emmanuel.my.last.name@some.domain.sub.domain.com",
			want: "e*******.m*.l***.n***@s***.d*****.s**.d*****.com",
		},
		{
			got:  "regular.name@domain-something.com",
			want: "r******.n***@d*****-s********.com",
		},
		{
			got:  "regular.name@alfa-beta-gamma-delta.com",
			want: "r******.n***@a***-b***-g****-d****.com",
		},
		{
			got:  "Luke.SkyWalker@McDonalds.com",
			want: "L***.S**W*****@M*D******.com",
		},
	}

	for _, test := range tests {
		result := anonymize.Email(test.got)
		if !strings.EqualFold(test.want, result) {
			t.Errorf("Incorrect result for '%v': got '%v', want '%v'", test.got, result, test.want)
		}
	}
}
