package main

import (
	tui "github.com/marcusolsson/tui-go"
)

func main() {
	l := tui.NewList()
	l.SetFocused(true)
	l.AddItems("First row", "Second row", "Third row", "Fourth row", "Fifth row", "Sixth row")

	root := tui.NewVBox(l)

	// This is still a rough prototype and WILL change. Use at your own risk.
	p := tui.NewPalette()
	p.SetItem("list.item.selected", tui.PaletteItem{
		Bg: tui.ColorRed,
		Fg: tui.ColorWhite,
	})

	ui := tui.New(root)
	ui.SetPalette(p)
	ui.SetKeybinding(tui.KeyEsc, func() {
		ui.Quit()
	})
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
