package propertybasedtests

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Desc   string
		Arabic int
		Want   string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}

	for _, test := range cases {
		t.Run(test.Desc, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
