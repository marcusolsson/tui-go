package main

import (
	"log"

	tui "github.com/marcusolsson/tui-go"
	termbox "github.com/nsf/termbox-go"
)

func main() {
	t := tui.NewTable(0, 0)
	t.AppendRow(tui.NewLabel("First row"))
	t.AppendRow(tui.NewLabel("Second row"))

	// This is still a rough prototype and WILL change. Use at your own risk.
	p := tui.NewPalette()
	p.SetItem("table.cell.selected", tui.PaletteItem{
		Bg: tui.Color(termbox.ColorRed),
		Fg: tui.Color(termbox.ColorWhite),
	})

	ui := tui.New(t)
	ui.SetPalette(p)

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
