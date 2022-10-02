package anonymize_test

import (
	"strings"
	"testing"

	"github.com/emmanuelay/anonymize"
)

func TestAnonymizeName(t *testing.T) {
	tests := []struct {
		got      string
		want     string
		anonRune rune
	}{
		{
			got:  "Emmanuel Ay",
			want: "E******* A*",
		},
		{
			got:      "Emmanuel Ay",
			want:     "E••••••• A•",
			anonRune: '•',
		},
		{
			got:  "Em Ay",
			want: "E* A*",
		},
		{
			got:  "E A",
			want: "E A",
		},
		{
			got:  "E",
			want: "E",
		},
		{
			got:  "   ",
			want: "",
		},
		{
			got:  "",
			want: "",
		},
		{
			got:  "Emmanuel von Ay",
			want: "E******* v** A*",
		},
		{
			got:  "Emmanuel McDonalds",
			want: "E******* M*D******",
		},
		{
			got:  "EMMANUEL MCDONALDS",
			want: "E******* M********",
		},
		{
			got:  "Luke Skywalker",
			want: "L*** S********",
		},
		{
			got:  "Luke SkyWalker",
			want: "L*** S**W*****",
		},
	}

	for _, test := range tests {
		var result string

		if test.anonRune > 0 {
			result = anonymize.NameWithCustomRune(test.got, test.anonRune)
		} else {
			result = anonymize.Name(test.got)
		}

		if !strings.EqualFold(test.want, result) {
			t.Errorf("Incorrect result for '%v': got '%v', want '%v'", test.got, result, test.want)
		}
	}
}
