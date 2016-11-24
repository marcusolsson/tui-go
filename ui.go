package tui

import termbox "github.com/nsf/termbox-go"

// UI holds the application state.
type UI struct {
	Painter *Painter
	Root    Widget

	shortcuts map[rune]func()
}

// New returns a new instance of an application.
func New(root Widget) *UI {
	s := NewTermboxSurface()
	p := NewPainter(s, DefaultPalette)

	return &UI{
		Painter:   p,
		Root:      root,
		shortcuts: make(map[rune]func()),
	}
}

func (ui *UI) SetPalette(p *Palette) {
	ui.Painter.palette = p
}

func (ui *UI) SetShortcut(ch rune, fn func()) {
	ui.shortcuts[ch] = fn
}

// Run starts the application and returns when application stops.
func (ui *UI) Run() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	ui.Painter.Repaint(ui.Root)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return nil
			}
			switch ev.Ch {
			case 'q':
				return nil
			}
			ui.notify(convertEvent(ev))
		case termbox.EventError:
			return ev.Err
		}

		ui.Painter.Repaint(ui.Root)
	}
}

func (ui *UI) notify(ev Event) {
	if fn, ok := ui.shortcuts[ev.Ch]; ok {
		fn()
	}
	ui.Root.OnEvent(ev)
}
