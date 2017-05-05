package tui

import (
	"image"

	"github.com/gdamore/tcell"
)

var _ UI = &tcellUI{}

type tcellUI struct {
	Painter *Painter
	Root    Widget

	keybindings []*Keybinding

	quit chan struct{}

	screen tcell.Screen

	kbFocus *KbFocusController
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
		kb.Ch = key
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
	ui.screen.Clear()

	ui.Painter.Repaint(ui.Root)

	eventCh := make(chan tcell.Event, 1)

	go func() {
		for {
			eventCh <- ui.screen.PollEvent()
		}
	}()

	for {
		select {
		case <-ui.quit:
			return nil
		case ev := <-eventCh:
			ui.notify(convertTcellEvent(ev))
			ui.Painter.Repaint(ui.Root)
		}
	}
}

// Quit signals to the UI to start shutting down.
func (ui *tcellUI) Quit() {
	ui.screen.Fini()
	ui.quit <- struct{}{}
}

func (ui *tcellUI) notify(ev Event) {
	for _, b := range ui.keybindings {
		if b.Match(ev) {
			b.Handler()
		}
	}
	ui.kbFocus.OnEvent(ev)
	ui.Root.OnEvent(ev)
}

func convertTcellEvent(tev tcell.Event) Event {
	switch tev := tev.(type) {
	case *tcell.EventKey:
		return Event{
			Type:      EventKey,
			Key:       convertTcellEventKey(tev.Key()),
			Ch:        tev.Rune(),
			Modifiers: ModMask(tev.Modifiers()),
		}
	default:
		return Event{}
	}
}

func convertTcellEventKey(key tcell.Key) Key {
	switch key {
	case tcell.KeyEnter:
		return KeyEnter
	case tcell.KeyTab:
		return KeyTab
	case tcell.KeyBacktab:
		return KeyBacktab
	case tcell.KeyEsc:
		return KeyEsc
	case tcell.KeyBackspace:
		return KeyBackspace
	case tcell.KeyBackspace2:
		return KeyBackspace2
	case tcell.KeyUp:
		return KeyArrowUp
	case tcell.KeyDown:
		return KeyArrowDown
	case tcell.KeyLeft:
		return KeyArrowLeft
	case tcell.KeyRight:
		return KeyArrowRight
	default:
		return KeyUnknown
	}
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
