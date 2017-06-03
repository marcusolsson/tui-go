package tui

import (
	"image"

	"github.com/gdamore/tcell"
)

var Repaint = func() {}

var _ UI = &tcellUI{}

type tcellUI struct {
	Painter *Painter
	Root    Widget

	keybindings []*Keybinding

	quit chan struct{}

	screen tcell.Screen

	kbFocus *KbFocusController

	eventQueue chan Event
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
		Painter:     p,
		Root:        root,
		keybindings: make([]*Keybinding, 0),
		quit:        make(chan struct{}, 1),
		screen:      screen,
		kbFocus:     &KbFocusController{chain: DefaultFocusChain},
		eventQueue:  make(chan Event, 1),
	}, nil
}

func (ui *tcellUI) SetTheme(p *Theme) {
	ui.Painter.Theme = p
}

func (ui *tcellUI) SetFocusChain(chain FocusChain) {
	ui.kbFocus.chain = chain
}

func (ui *tcellUI) SetKeybinding(k interface{}, fn func()) {
	kb := new(Keybinding)

	switch key := k.(type) {
	case rune:
		kb.Rune = key
	case Key:
		kb.Key = key
	}
	kb.Handler = fn

	ui.keybindings = append(ui.keybindings, kb)
}

func (ui *tcellUI) Run() error {
	if err := ui.screen.Init(); err != nil {
		return err
	}

	if w := ui.kbFocus.chain.FocusDefault(); w != nil {
		w.SetFocused(true)
		ui.kbFocus.focusedWidget = w
	}

	ui.screen.SetStyle(tcell.StyleDefault)
	ui.screen.EnableMouse()
	ui.screen.Clear()

	Repaint = func() {
		ui.eventQueue <- PaintEvent{}
	}

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

	Repaint()

	for {
		select {
		case <-ui.quit:
			return nil
		case ev := <-ui.eventQueue:
			ui.handleEvent(ev)
		}
	}
}

func (ui *tcellUI) handleEvent(ev Event) {
	switch e := ev.(type) {
	case KeyEvent:
		for _, b := range ui.keybindings {
			if b.Match(e) {
				b.Handler()
			}
		}
		ui.kbFocus.OnKeyEvent(e)
		ui.Root.OnKeyEvent(e)
		ui.Painter.Repaint(ui.Root)
	case PaintEvent:
		ui.Painter.Repaint(ui.Root)
	}
}

func (ui *tcellUI) handleKeyEvent(tev *tcell.EventKey) {
	ui.eventQueue <- KeyEvent{
		Key:  convertTcellEventKey(tev.Key()),
		Rune: tev.Rune(),
	}
}

func (ui *tcellUI) handleMouseEvent(ev *tcell.EventMouse) {
	x, y := ev.Position()
	ui.eventQueue <- MouseEvent{Pos: image.Pt(x, y)}
}

func (ui *tcellUI) handleResizeEvent(ev *tcell.EventResize) {
	ui.eventQueue <- PaintEvent{}
}

// Quit signals to the UI to start shutting down.
func (ui *tcellUI) Quit() {
	ui.screen.Fini()
	ui.quit <- struct{}{}
}

var _ Surface = &tcellSurface{}

type tcellSurface struct {
	screen tcell.Screen
}

func (s *tcellSurface) SetCell(x, y int, ch rune, style Style) {
	st := tcell.StyleDefault.Normal().
		Foreground(convertColor(style.Fg, false)).
		Background(convertColor(style.Bg, false))

	s.screen.SetContent(x, y, ch, nil, st)
}

func (s *tcellSurface) SetCursor(x, y int) {
	s.screen.ShowCursor(x, y)
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

var tcellKeyMap = map[tcell.Key]Key{
	tcell.KeyEnter:      KeyEnter,
	tcell.KeyTab:        KeyTab,
	tcell.KeyBacktab:    KeyBacktab,
	tcell.KeyEsc:        KeyEsc,
	tcell.KeyBackspace:  KeyBackspace,
	tcell.KeyBackspace2: KeyBackspace2,
	tcell.KeyUp:         KeyArrowUp,
	tcell.KeyDown:       KeyArrowDown,
	tcell.KeyLeft:       KeyArrowLeft,
	tcell.KeyRight:      KeyArrowRight,
}

func convertTcellEventKey(key tcell.Key) Key {
	k, ok := tcellKeyMap[key]
	if !ok {
		return KeyUnknown
	}
	return k
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
		return tcell.ColorDefault
	}
}
