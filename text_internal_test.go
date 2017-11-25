package tui

import "testing"

func TestRuneWidth(t *testing.T) {
	for n, tt := range []struct {
		r      rune
		result int
	}{
		{' ', 1},
		{'a', 1},
		{'あ', 2},
	} {
		if got, want := runeWidth(tt.r), tt.result; got != want {
			t.Errorf("[%d] runeWidth(%q) = %d, want %d", n, tt.r, got, want)
		}
	}
}

func TestStringWidth(t *testing.T) {
	for n, tt := range []struct {
		s      string
		result int
	}{
		{"", 0},
		{" ", 1},
		{"a", 1},
		{"a ", 2},
		{"abc", 3},
		{"あ", 2},
		{"あいう", 6},
		{"abcあいう123", 12},
	} {
		if got, want := stringWidth(tt.s), tt.result; got != want {
			t.Errorf("[%d] stringWidth(%q) = %d, want %d", n, tt.s, got, want)
		}
	}
}

func TestTrimRightLen(t *testing.T) {
	for n, tt := range []struct {
		s      string
		n      int
		result string
	}{
		{"", 0, ""},
		{"", 1, ""},
		{"", -1, ""},
		{" ", 1, ""},
		{"abc", -1, "abc"},
		{"abc", 0, "abc"},
		{"abc", 1, "ab"},
		{"abc ", 1, "abc"},
		{"abc", 2, "a"},
		{"abc", 3, ""},
		{"abc", 4, ""},
		{"あいう", -1, "あいう"},
		{"あいう", 0, "あいう"},
		{"あいう", 1, "あい"},
		{"あいう", 2, "あ"},
		{"あいう", 3, ""},
		{"あいう", 4, ""},
	} {
		if got, want := trimRightLen(tt.s, tt.n), tt.result; got != want {
			t.Errorf("[%d] trimRightLen(%q, %d) = %q, want %q", n, tt.s, tt.n, got, want)
		}
	}
}
