package tui

type UI interface {
	SetTheme(p *Theme)
	SetKeybinding(k interface{}, fn func())
	Run() error
	Quit()
}

func New(root Widget) UI {
	tcellui, _ := newTcellUI(root)
	return tcellui
}
