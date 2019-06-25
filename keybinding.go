package tui

import (
	"strings"
)

type keybinding struct {
	sequence string
	handler  func()
}

func (b *keybinding) match(ev KeyEvent) bool {
	return strings.EqualFold(b.sequence, ev.Name())
}
