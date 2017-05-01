package main

import (
	"github.com/marcusolsson/tui-go"
)

func main() {
	l := tui.NewList()
	l.SetFocused(true)
	l.AddItems("First row", "Second row", "Third row", "Fourth row", "Fifth row", "Sixth row")
	l.SetSelected(0)

	root := tui.NewVBox(l)

	// This is still a rough prototype and WILL change. Use at your own risk.
	t := tui.NewTheme()
	t.SetStyle("list.item", tui.Style{Bg: tui.ColorCyan, Fg: tui.ColorMagenta})
	t.SetStyle("list.item.selected", tui.Style{Bg: tui.ColorRed, Fg: tui.ColorWhite})

	ui := tui.New(root)
	ui.SetTheme(t)
	ui.SetKeybinding(tui.KeyEsc, func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
