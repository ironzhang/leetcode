package leetcode

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		x, s, y int
	}{
		{0, 1, 0},
		{1, 1, 1},
		{10, 1, 10},
		{-1, -1, 1},
		{-10, -1, 10},
	}
	for _, tt := range tests {
		s, y := abs(tt.x)
		if got, want := s, tt.s; got != want {
			t.Errorf("abs(%d)'sign: %d != %d", tt.x, got, want)
		}
		if got, want := y, tt.y; got != want {
			t.Errorf("abs(%d)'y: %d != %d", tt.x, got, want)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		x, y int
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{10, 1},
		{101, 101},
		{-21, -12},
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	}
	for _, tt := range tests {
		if got, want := reverse(tt.x), tt.y; got != want {
			t.Errorf("reverse(%d): %d != %d", tt.x, got, want)
		}
	}
}
