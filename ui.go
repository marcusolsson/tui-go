package tui

// UI defines the operations needed by the underlying engine.
type UI interface {
	// SetWidget sets the root widget of the UI.
	SetWidget(w Widget)
	// SetTheme sets the current theme of the UI.
	SetTheme(p *Theme)
	// SetKeybinding sets the callback for when a key sequence is pressed.
	SetKeybinding(seq string, fn func())
	// ClearKeybindings removes all previous set keybindings.
	ClearKeybindings()
	// SetFocusChain sets a chain of widgets that determines focus order.
	SetFocusChain(ch FocusChain)
	// Run starts the UI goroutine and blocks either Quit was called or an error occurred.
	Run() error
	// Update schedules work in the UI thread and await its completion.
	// Note that calling Update from the UI thread will result in deadlock.
	Update(fn func())
	// Quit shuts down the UI goroutine.
	Quit()

	ShowDialog(d *Dialog)
	HideDialog()
}

// New returns a new UI with a root widget.
func New(root Widget) (UI, error) {
	return newTcellUI(root)
}
