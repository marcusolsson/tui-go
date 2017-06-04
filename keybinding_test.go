package tui

import "testing"

func TestKeybinding_Match(t *testing.T) {
	for _, tt := range []struct {
		Binding Keybinding
		Event   KeyEvent
		Match   bool
	}{
		{Keybinding{Sequence: "a"}, KeyEvent{Key: KeyRune, Rune: 'a'}, true},
		{Keybinding{Sequence: "a"}, KeyEvent{Key: KeyRune, Rune: 'l'}, false},
		{Keybinding{Sequence: "Enter"}, KeyEvent{Key: KeyRune, Rune: 'l'}, false},
		{Keybinding{Sequence: "Enter"}, KeyEvent{Key: KeyEnter}, true},
		{Keybinding{Sequence: "Ctrl+Space"}, KeyEvent{Key: KeyCtrlSpace, Modifiers: ModCtrl}, true},
	} {
		tt := tt
		t.Run(tt.Binding.Sequence, func(t *testing.T) {
			if got := tt.Binding.Match(tt.Event); got != tt.Match {
				t.Errorf("got = %v; want = %v", got, tt.Match)
			}
		})
	}
}
