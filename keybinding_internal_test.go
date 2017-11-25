package tui

import (
	"testing"
)

func TestKeybinding_Match(t *testing.T) {
	for _, tt := range []struct {
		binding keybinding
		event   KeyEvent
		match   bool
	}{
		{keybinding{sequence: "a"}, KeyEvent{Key: KeyRune, Rune: 'a'}, true},
		{keybinding{sequence: "a"}, KeyEvent{Key: KeyRune, Rune: 'l'}, false},
		{keybinding{sequence: "Enter"}, KeyEvent{Key: KeyRune, Rune: 'l'}, false},
		{keybinding{sequence: "Enter"}, KeyEvent{Key: KeyEnter}, true},
		{keybinding{sequence: "Ctrl+Space"}, KeyEvent{Key: KeyCtrlSpace, Modifiers: ModCtrl}, true},
	} {
		tt := tt
		t.Run(tt.binding.sequence, func(t *testing.T) {
			if got := tt.binding.match(tt.event); got != tt.match {
				t.Errorf("got = %v; want = %v", got, tt.match)
			}
		})
	}
}
