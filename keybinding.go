package tui

type Keybinding struct {
	Key  Key
	Rune rune

	Handler func()
}

func (b *Keybinding) Match(ev KeyEvent) bool {
	return b.Key == ev.Key && b.Rune == ev.Rune
}
