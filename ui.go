package tui

type UI interface {
	SetWidget(w Widget)
	SetTheme(p *Theme)
	SetKeybinding(seq string, fn func())
	SetFocusChain(ch FocusChain)
	Run() error
	Update(fn func())
	Quit()
}

func New(root Widget) UI {
	tcellui, _ := newTcellUI(root)
	return tcellui
}
