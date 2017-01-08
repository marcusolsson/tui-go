package tui

type keybinding struct {
	key     Key
	ch      rune
	handler func()
}

func (b *keybinding) match(ev Event) bool {
	return (b.key == ev.Key) && (b.ch == ev.Ch)
}
