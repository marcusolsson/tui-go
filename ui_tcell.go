package tui

import (
	"image"

	"github.com/gdamore/tcell"
)

var _ UI = &tcellUI{}

type tcellUI struct {
	painter *Painter
	root    Widget

	keybindings []*keybinding

	quit chan struct{}

	screen tcell.Screen

	kbFocus *kbFocusController

	eventQueue chan event
}

func newTcellUI(root Widget) (*tcellUI, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	s := &tcellSurface{
		screen: screen,
	}
	p := NewPainter(s, DefaultTheme)

	return &tcellUI{
		painter:     p,
		root:        root,
		keybindings: make([]*keybinding, 0),
		quit:        make(chan struct{}, 1),
		screen:      screen,
		kbFocus:     &kbFocusController{chain: DefaultFocusChain},
		eventQueue:  make(chan event),
	}, nil
}

func (ui *tcellUI) Repaint() {
	ui.painter.Repaint(ui.root)
}

func (ui *tcellUI) SetWidget(w Widget) {
	ui.root = w
}

func (ui *tcellUI) SetTheme(t *Theme) {
	ui.painter.theme = t
}

func (ui *tcellUI) SetFocusChain(chain FocusChain) {
	if ui.kbFocus.focusedWidget != nil {
		ui.kbFocus.focusedWidget.SetFocused(false)
	}

	ui.kbFocus.chain = chain
	ui.kbFocus.focusedWidget = chain.FocusDefault()

	if ui.kbFocus.focusedWidget != nil {
		ui.kbFocus.focusedWidget.SetFocused(true)
	}
}

func (ui *tcellUI) SetKeybinding(seq string, fn func()) {
	ui.keybindings = append(ui.keybindings, &keybinding{
		sequence: seq,
		handler:  fn,
	})
}

// ClearKeybindings reinitialises ui.keybindings so as to revert to a
// clear/original state
func (ui *tcellUI) ClearKeybindings() {
	ui.keybindings = make([]*keybinding, 0)
}

func (ui *tcellUI) Run() error {
	if err := ui.screen.Init(); err != nil {
		return err
	}

	failed := true
	defer func() {
		if failed {
			ui.screen.Fini()
		}
	}()

	if w := ui.kbFocus.chain.FocusDefault(); w != nil {
		w.SetFocused(true)
		ui.kbFocus.focusedWidget = w
	}

	ui.screen.SetStyle(tcell.StyleDefault)
	ui.screen.Clear()

	go func() {
		for {
			switch ev := ui.screen.PollEvent().(type) {
			case *tcell.EventKey:
				ui.handleKeyEvent(ev)
			case *tcell.EventMouse:
				ui.handleMouseEvent(ev)
			case *tcell.EventResize:
				ui.handleResizeEvent(ev)
			}
		}
	}()

	for {
		select {
		case <-ui.quit:
			failed = false
			return nil
		case ev := <-ui.eventQueue:
			ui.handleEvent(ev)
		}
	}
}

func (ui *tcellUI) handleEvent(ev event) {
	switch e := ev.(type) {
	case KeyEvent:
		logger.Printf("Received key event: %s", e.Name())

		for _, b := range ui.keybindings {
			if b.match(e) {
				b.handler()
			}
		}
		ui.kbFocus.OnKeyEvent(e)
		ui.root.OnKeyEvent(e)
		ui.painter.Repaint(ui.root)
	case callbackEvent:
		// Gets stuck in a print loop when the logger is a widget.
		//logger.Printf("Received callback event")
		e.cbFn()
		ui.painter.Repaint(ui.root)
	case paintEvent:
		logger.Printf("Received paint event")
		ui.painter.Repaint(ui.root)
	}
}

func (ui *tcellUI) handleKeyEvent(tev *tcell.EventKey) {
	ui.eventQueue <- KeyEvent{
		Key:       Key(tev.Key()),
		Rune:      tev.Rune(),
		Modifiers: ModMask(tev.Modifiers()),
	}
}

func (ui *tcellUI) handleMouseEvent(ev *tcell.EventMouse) {
	x, y := ev.Position()
	ui.eventQueue <- MouseEvent{Pos: image.Pt(x, y)}
}

func (ui *tcellUI) handleResizeEvent(ev *tcell.EventResize) {
	ui.eventQueue <- paintEvent{}
}

// Quit signals to the UI to start shutting down.
func (ui *tcellUI) Quit() {
	logger.Printf("Quitting")
	ui.screen.Fini()
	ui.quit <- struct{}{}
}

// Schedule an update of the UI, running the given
// function in the UI goroutine.
//
// Use this to update the UI in response to external events,
// like a timer tick.
// This method should be used any time you call methods
// to change UI objects after the first call to `UI.Run()`.
//
// Changes invoked outside of either this callback or the
// other event handler callbacks may appear to work, but
// is likely a race condition.  (Run your program with
// `go run -race` or `go install -race` to detect this!)
//
// Calling Update from within an event handler, or from within an Update call,
// is an error, and will deadlock.
func (ui *tcellUI) Update(fn func()) {
	blk := make(chan struct{})
	ui.eventQueue <- callbackEvent{func() {
		fn()
		close(blk)
	}}
	<-blk
}

var _ Surface = &tcellSurface{}

type tcellSurface struct {
	screen tcell.Screen
}

func (s *tcellSurface) SetCell(x, y int, ch rune, style Style) {
	st := tcell.StyleDefault.Normal().
		Foreground(convertColor(style.Fg, false)).
		Background(convertColor(style.Bg, false)).
		Reverse(style.Reverse == DecorationOn).
		Bold(style.Bold == DecorationOn).
		Underline(style.Underline == DecorationOn)

	s.screen.SetContent(x, y, ch, nil, st)
}

func (s *tcellSurface) SetCursor(x, y int) {
	s.screen.ShowCursor(x, y)
}

func (s *tcellSurface) HideCursor() {
	s.screen.HideCursor()
}

func (s *tcellSurface) Begin() {
	s.screen.Clear()
}

func (s *tcellSurface) End() {
	s.screen.Show()
}

func (s *tcellSurface) Size() image.Point {
	w, h := s.screen.Size()
	return image.Point{w, h}
}

func convertColor(col Color, fg bool) tcell.Color {
	switch col {
	case ColorDefault:
		if fg {
			return tcell.ColorWhite
		}
		return tcell.ColorDefault
	case ColorBlack:
		return tcell.ColorBlack
	case ColorWhite:
		return tcell.ColorWhite
	case ColorRed:
		return tcell.ColorRed
	case ColorGreen:
		return tcell.ColorGreen
	case ColorBlue:
		return tcell.ColorBlue
	case ColorCyan:
		return tcell.ColorDarkCyan
	case ColorMagenta:
		return tcell.ColorDarkMagenta
	case ColorYellow:
		return tcell.ColorYellow
	default:
		if col > 0 {
			return tcell.Color(col)
		}
		return tcell.ColorDefault
	}
}
