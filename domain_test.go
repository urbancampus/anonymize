package anonymize_test

import (
	"strings"
	"testing"

	"github.com/emmanuelay/anonymize"
)

func TestAnonymizeDomain(t *testing.T) {
	tests := []struct {
		got  string
		want string
	}{
		{
			got:  "www.somedomain.com",
			want: "www.s*********.com",
		},
		{
			got:  "f.com",
			want: "*.com",
		},
		{
			got:  "do.com",
			want: "d*.com",
		},
		{
			got:  "www.this-should-work.com",
			want: "www.t***-s*****-w***.com",
		},
		{
			got:  "www.this-should-work.com/some/path",
			want: "www.t***-s*****-w***.com/some/path",
		},
		{
			got:  "w**",
			want: "w**",
		},
		{
			got:  "words.com",
			want: "w****.com",
		},
		{
			got:  "domain.com",
			want: "d*****.com",
		},
	}

	for _, test := range tests {
		result := anonymize.Domain(test.got)
		if !strings.EqualFold(test.want, result) {
			t.Errorf("Incorrect result for '%v': got '%v', want '%v'", test.got, result, test.want)
		}
	}
}
