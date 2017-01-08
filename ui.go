package tui

import termbox "github.com/nsf/termbox-go"

// UI holds the application state.
type UI struct {
	Painter *Painter
	Root    Widget

	keybindings []*keybinding

	quit chan struct{}
}

// New returns a new instance of an application.
func New(root Widget) *UI {
	s := NewTermboxSurface()
	p := NewPainter(s, DefaultPalette)

	return &UI{
		Painter:     p,
		Root:        root,
		keybindings: make([]*keybinding, 0),
		quit:        make(chan struct{}, 1),
	}
}

func (ui *UI) SetPalette(p *Palette) {
	ui.Painter.palette = p
}

func (ui *UI) SetKeybinding(k interface{}, fn func()) {
	kb := new(keybinding)

	switch key := k.(type) {
	case rune:
		kb.ch = key
	case Key:
		kb.key = key
	}
	kb.handler = fn

	ui.keybindings = append(ui.keybindings, kb)
}

// Run starts the application and returns when application stops.
func (ui *UI) Run() error {
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
				ui.notify(convertEvent(ev))
			case termbox.EventError:
				return ev.Err
			}

			ui.Painter.Repaint(ui.Root)
		}
	}
}

// Quit signals to the UI to start shutting down.
func (ui *UI) Quit() {
	ui.quit <- struct{}{}
}

func (ui *UI) notify(ev Event) {
	for _, b := range ui.keybindings {
		if b.match(ev) {
			b.handler()
		}
	}
	ui.Root.OnEvent(ev)
}
