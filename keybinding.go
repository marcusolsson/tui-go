package tui

type Keybinding struct {
	Key     Key
	Rune    rune
	Mod     ModMask
	Handler func()
}

func (b *Keybinding) Match(ev KeyEvent) bool {
	if b.Key != KeyUnknown {
		return b.Rune == ev.Rune && b.Mod == ev.Mod
	}

	return b.Key == ev.Key && b.Rune == ev.Rune
}
