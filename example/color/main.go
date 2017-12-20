package main

import (
	"github.com/marcusolsson/tui-go"
)

// StyledBox is a Box with an overriden Draw method.
// Embedding a Widget within another allows overriding of some behaviors.
type StyledBox struct {
	Style string
	*tui.Box
}
func (s *StyledBox) Draw(p *tui.Painter) {
	p.WithStyle(s.Style, func(p *tui.Painter) {
		s.Box.Draw(p)
	})
}

func main() {
	t := tui.NewTheme()
	normal := tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack}
	t.SetStyle("normal", normal)

	// A simple label.
	okay := tui.NewLabel("Everything is fine.")

	// A list with some items selected.
	l := tui.NewList()
	l.SetFocused(true)
	l.AddItems("First row", "Second row", "Third row", "Fourth row", "Fifth row", "Sixth row")
	l.SetSelected(0)

	t.SetStyle("list.item", tui.Style{Bg: tui.ColorCyan, Fg: tui.ColorMagenta})
	t.SetStyle("list.item.selected", tui.Style{Bg: tui.ColorRed, Fg: tui.ColorWhite})

		// The style name is appended to the widget name to support coloring of
	// individual labels.
	warning := tui.NewLabel("WARNING: This is a warning")
	warning.SetStyleName("warning")
	t.SetStyle("label.warning", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorYellow})

	fatal := tui.NewLabel("FATAL: Cats and dogs are now living together.")
	fatal.SetStyleName("fatal")
	t.SetStyle("label.fatal", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorRed})

	// Styles inherit properties of the parent widget by default;
	// setting a property overrides only that property.
	message1 := tui.NewLabel("This is an ")
	emphasis := tui.NewLabel("important")
	message2 := tui.NewLabel(" message from our sponsors.")
	message := &StyledBox{
		Style: "bsod",
		Box: tui.NewHBox(message1, emphasis, message2, tui.NewSpacer()),
	}

	emphasis.SetStyleName("emphasis")
	t.SetStyle("label.emphasis", tui.Style{Bold: tui.DecorationOn, Underline: tui.DecorationOn, Bg: tui.ColorRed})
	t.SetStyle("bsod", tui.Style{Bg: tui.ColorCyan, Fg: tui.ColorWhite})

	// Another unstyled label.
	okay2 := tui.NewLabel("Everything is still fine.")

	root := tui.NewVBox(okay, l, warning, fatal, message, okay2)

	ui := tui.New(root)
	ui.SetTheme(t)
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
