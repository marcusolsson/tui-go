package main

import (
	"github.com/marcusolsson/tui-go"
)

func main() {
	buffer := tui.NewTextEdit()
	buffer.SetSizePolicy(tui.Expanding, tui.Expanding)
	buffer.SetText("Hello, world!")
	buffer.SetFocused(true)

	status := tui.NewStatusBar("main.go")

	root := tui.NewVBox(buffer, status)

	ui := tui.New(root)
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
