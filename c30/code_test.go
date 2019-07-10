package leetcode

import (
	"reflect"
	"testing"
)

func TestMakeWordHashMap(t *testing.T) {
	tests := []struct {
		words   []string
		hashMap map[string]int
	}{
		{
			words: []string{"foo", "bar"},
			hashMap: map[string]int{
				"foo": 1,
				"bar": 1,
			},
		},
		{
			words: []string{"foo", "bar", "foo"},
			hashMap: map[string]int{
				"foo": 2,
				"bar": 1,
			},
		},
	}
	for i, tt := range tests {
		hashMap := makeWordHashMap(tt.words)
		if got, want := hashMap, tt.hashMap; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		s      string
		words  []string
		output []int
	}{
		{
			s:      "barfoothefoobarman",
			words:  []string{"foo", "bar"},
			output: []int{0, 9},
		},
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "word"},
			output: nil,
		},
		{
			s:      "",
			words:  []string{},
			output: nil,
		},
	}
	for i, tt := range tests {
		output := findSubstring(tt.s, tt.words)
		if got, want := output, tt.output; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}
