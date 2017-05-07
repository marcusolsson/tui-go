package tui

import "testing"

func TestKeybinding_Match(t *testing.T) {
	for _, tt := range []struct {
		Binding Keybinding
		Event   KeyEvent
		Match   bool
	}{
		{Keybinding{Rune: 'a'}, KeyEvent{Rune: 'a'}, true},
		{Keybinding{Rune: 'a'}, KeyEvent{Rune: 'l'}, false},
		{Keybinding{Key: KeyEnter}, KeyEvent{Rune: 'l'}, false},
		{Keybinding{Key: KeyEnter}, KeyEvent{Key: KeyEnter}, true},
	} {
		if got := tt.Binding.Match(tt.Event); got != tt.Match {
			t.Errorf("got = %v; want = %v", got, tt.Match)
		}
	}
}
