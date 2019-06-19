package leetcode

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		zigzag string
		row    int
		line   string
	}{
		{
			zigzag: "PAYPALISHIRING",
			row:    3,
			line:   "PAHNAPLSIIGYIR",
		},
		{
			zigzag: "PAYPALISHIRING",
			row:    4,
			line:   "PINALSIGYAHRPI",
		},
		{
			zigzag: "",
			row:    4,
			line:   "",
		},
	}
	for i, tt := range tests {
		if got, want := convert(tt.zigzag, tt.row), tt.line; got != want {
			t.Errorf("case%d: %q != %q", i, got, want)
		}
	}
}
