package leetcode

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		input  string
		output []string
	}{
		{
			input:  "23",
			output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			input:  "",
			output: []string{},
		},
	}
	for i, tt := range tests {
		if got, want := letterCombinations(tt.input), tt.output; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: letterCombinations: got %q, want %q", i, got, want)
		}
	}
}
