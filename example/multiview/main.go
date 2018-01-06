package main

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

func main() {
	var currentView int

	views := []tui.Widget{
		tui.NewVBox(tui.NewLabel("Press right arrow to continue ...")),
		tui.NewVBox(tui.NewLabel("Almost there, one more time!")),
		tui.NewVBox(tui.NewLabel("Congratulations, you've finished the example!")),
	}

	root := tui.NewVBox(views[0])

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Left", func() {
		currentView = clamp(currentView-1, 0, len(views)-1)
		ui.SetWidget(views[currentView])
	})
	ui.SetKeybinding("Right", func() {
		currentView = clamp(currentView+1, 0, len(views)-1)
		ui.SetWidget(views[currentView])
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func clamp(n, min, max int) int {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}
