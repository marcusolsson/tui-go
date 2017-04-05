package tui

type UI interface {
	SetPalette(p *Palette)
	SetKeybinding(k interface{}, fn func())
	Run() error
	Quit()
}

func New(root Widget) UI {
	return newTermboxUI(root)
}
