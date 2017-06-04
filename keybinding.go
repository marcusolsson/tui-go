package tui

import (
	"strings"
)

type Keybinding struct {
	Sequence string
	Handler  func()
}

func (b *Keybinding) Match(ev KeyEvent) bool {
	return strings.ToLower(b.Sequence) == strings.ToLower(ev.Name())
}
