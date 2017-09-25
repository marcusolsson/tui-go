package tui

// UI defines the operations needed by the underlying engine.
type UI interface {
	SetWidget(w Widget)
	SetTheme(p *Theme)
	SetKeybinding(seq string, fn func())
	SetFocusChain(ch FocusChain)
	Run() error
	Update(fn func())
	Quit()
}

// New returns a new UI with a root widget.
func New(root Widget) UI {
	tcellui, _ := newTcellUI(root)
	return tcellui
}
