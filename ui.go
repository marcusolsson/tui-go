package tui

type UI interface {
	SetTheme(p *Theme)
	SetKeybinding(seq string, fn func())
	SetFocusChain(ch FocusChain)
	Run() error
	Quit()
}

func New(root Widget) UI {
	tcellui, _ := newTcellUI(root)
	return tcellui
}
