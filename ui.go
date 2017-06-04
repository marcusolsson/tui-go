package tui

type UI interface {
	SetTheme(p *Theme)
	SetKeybinding(k interface{}, m ModMask, fn func())
	SetFocusChain(ch FocusChain)
	Run() error
	Quit()
}

func New(root Widget) UI {
	tcellui, _ := newTcellUI(root)
	return tcellui
}
