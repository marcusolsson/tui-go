package tui

import (
	"strings"
)

type keybinding struct {
	sequence string
	handler  func()
}

func (b *keybinding) match(ev KeyEvent) bool {
	return strings.ToLower(b.sequence) == strings.ToLower(ev.Name())
}
