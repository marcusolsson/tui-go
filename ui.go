package tui

// UI defines the operations needed by the underlying engine.
type UI interface {
	// Change the root widget of the UI.
	SetWidget(w Widget)
	SetTheme(p *Theme)
	SetKeybinding(seq string, fn func())
	ClearKeybindings()
	SetFocusChain(ch FocusChain)
	// Start the UI thread; wait for it to complete, either via error or Quit.
	Run() error
	// Schedule work in the UI thread and await its completion.
	// Note that calling Update from the UI thread will result in deadlock.
	Update(fn func())
	// Shut down the UI thread.
	Quit()
}

// New returns a new UI with a root widget.
func New(root Widget) UI {
	tcellui, _ := newTcellUI(root)
	return tcellui
}
