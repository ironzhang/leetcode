package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input []byte
		want  bool
	}{
		{input: []byte(""), want: true},
		{input: []byte("a"), want: true},
		{input: []byte("aa"), want: true},
		{input: []byte("aaa"), want: true},
		{input: []byte("aba"), want: true},
		{input: []byte("abba"), want: true},

		{input: []byte("ab"), want: false},
		{input: []byte("abc"), want: false},
		{input: []byte("abca"), want: false},
	}

	for i, test := range tests {
		if got, want := isPalindrome(test.input), test.want; got != want {
			t.Errorf("testcase(%d): %v != %v", i, got, want)
		}
	}
}

func TestLongestPalindromeBytes(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{input: []byte(""), want: []byte("")},
		{input: []byte("a"), want: []byte("a")},
		{input: []byte("aa"), want: []byte("aa")},
		{input: []byte("aaa"), want: []byte("aaa")},
		{input: []byte("aba"), want: []byte("aba")},
		{input: []byte("abba"), want: []byte("abba")},

		{input: []byte("ab"), want: []byte("a")},
		{input: []byte("abc"), want: []byte("a")},
		{input: []byte("abca"), want: []byte("a")},
		{input: []byte("abaa"), want: []byte("aba")},
		{input: []byte("abaaaa"), want: []byte("aaaa")},
	}

	for i, test := range tests {
		if got, want := longestPalindromeBytes_DFS(test.input), test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("testcase(%d): DFS(%s): %s != %s", i, test.input, got, want)
		}
		if got, want := longestPalindromeBytes_BFS(test.input), test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("testcase(%d): BFS(%s): %s != %s", i, test.input, got, want)
		}
	}
}

func TestTimeLimit(t *testing.T) {
	input := "babaddtattarrattatddetartrateedredividerb"
	output := longestPalindrome(input)
	fmt.Printf("%s\n", output)
}
