package tui

type Keybinding struct {
	Key     Key
	Ch      rune
	Handler func()
}

func (b *Keybinding) Match(ev Event) bool {
	return (b.Key == ev.Key) && (b.Ch == ev.Ch)
}
