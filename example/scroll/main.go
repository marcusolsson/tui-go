package main

import (
	"log"

	tui "github.com/marcusolsson/tui-go"
)

var lorem = `Lorem ipsum dolor sit amet.`

func main() {
	root := tui.NewHBox()

	l := tui.NewLabel(lorem)

	s := tui.NewScrollArea(l)

	scrollBox := tui.NewVBox(s)
	scrollBox.SetBorder(true)

	root.Append(tui.NewVBox(tui.NewSpacer()))
	root.Append(scrollBox)
	root.Append(tui.NewVBox(tui.NewSpacer()))

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Up", func() { s.Scroll(0, -1) })
	ui.SetKeybinding("Down", func() { s.Scroll(0, 1) })
	ui.SetKeybinding("Left", func() { s.Scroll(-1, 0) })
	ui.SetKeybinding("Right", func() { s.Scroll(1, 0) })
	ui.SetKeybinding("a", func() { s.SetAutoscrollToBottom(true) })
	ui.SetKeybinding("t", func() { s.ScrollToTop() })
	ui.SetKeybinding("b", func() { s.ScrollToBottom() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
