package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

type termboxUI struct {
	Painter *Painter
	Root    Widget

	keybindings []*Keybinding

	quit chan struct{}
}

func newTermboxUI(root Widget) *termboxUI {
	s := &termboxSurface{}
	p := NewPainter(s, DefaultPalette)

	return &termboxUI{
		Painter:     p,
		Root:        root,
		keybindings: make([]*Keybinding, 0),
		quit:        make(chan struct{}, 1),
	}
}

func (ui *termboxUI) SetPalette(p *Palette) {
	ui.Painter.Palette = p
}

func (ui *termboxUI) SetKeybinding(k interface{}, fn func()) {
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

// Run starts the application and returns when application stops.
func (ui *termboxUI) Run() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	ui.Painter.Repaint(ui.Root)

	eventCh := make(chan termbox.Event, 1)

	go func() {
		for {
			eventCh <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case <-ui.quit:
			return nil
		case ev := <-eventCh:
			switch ev.Type {
			case termbox.EventKey:
				ui.notify(convertTermboxEvent(ev))
			case termbox.EventError:
				return ev.Err
			}

			ui.Painter.Repaint(ui.Root)
		}
	}
}

// Quit signals to the UI to start shutting down.
func (ui *termboxUI) Quit() {
	ui.quit <- struct{}{}
}

func (ui *termboxUI) notify(ev Event) {
	for _, b := range ui.keybindings {
		if b.Match(ev) {
			b.Handler()
		}
	}
	ui.Root.OnEvent(ev)
}

func convertTermboxEvent(tev termbox.Event) Event {
	return Event{
		Type: convertTermboxEventType(tev.Type),
		Key:  convertTermboxEventKey(tev.Key),
		Ch:   tev.Ch,
	}
}

func convertTermboxEventType(t termbox.EventType) EventType {
	switch t {
	case termbox.EventKey:
		return EventKey
	case termbox.EventResize:
		return EventResize
	case termbox.EventMouse:
		return EventMouse
	case termbox.EventError:
		return EventError
	case termbox.EventInterrupt:
		return EventInterrupt
	case termbox.EventRaw:
		return EventRaw
	case termbox.EventNone:
		return EventNone
	default:
		return EventUnknown
	}
}

func convertTermboxEventKey(key termbox.Key) Key {
	switch key {
	case termbox.KeyEnter:
		return KeyEnter
	case termbox.KeySpace:
		return KeySpace
	case termbox.KeyTab:
		return KeyTab
	case termbox.KeyEsc:
		return KeyEsc
	case termbox.KeyBackspace:
		return KeyBackspace
	case termbox.KeyBackspace2:
		return KeyBackspace2
	case termbox.KeyArrowUp:
		return KeyArrowUp
	case termbox.KeyArrowDown:
		return KeyArrowDown
	case termbox.KeyArrowLeft:
		return KeyArrowLeft
	case termbox.KeyArrowRight:
		return KeyArrowRight
	default:
		return KeyUnknown
	}
}

type termboxSurface struct{}

func (s termboxSurface) SetCell(x, y int, ch rune, fg, bg Color) {
	termbox.SetCell(x, y, ch, convertTermboxColor(fg), convertTermboxColor(bg))
}

func (s termboxSurface) SetCursor(x, y int) {
	termbox.SetCursor(x, y)
}

func (s termboxSurface) Begin() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func (s termboxSurface) End() {
	termbox.Flush()
}

func (s termboxSurface) Size() image.Point {
	w, h := termbox.Size()
	return image.Point{w, h}
}

func convertTermboxColor(col Color) termbox.Attribute {
	switch col {
	case ColorDefault:
		return termbox.ColorDefault
	case ColorBlack:
		return termbox.ColorBlack
	case ColorWhite:
		return termbox.ColorWhite
	case ColorBlue:
		return termbox.ColorBlue
	case ColorRed:
		return termbox.ColorRed
	default:
		return termbox.ColorDefault
	}
}
