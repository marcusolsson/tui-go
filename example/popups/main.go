// popups is a program demonstrating modal dialogs.
package main

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

func Theme() *tui.Theme {
	t := tui.NewTheme()
	t.SetStyle("reversed", tui.Style{
		Reverse: tui.DecorationOn,
	})
	return t
}

func main() {
	fm := &FocusManager{}

	list := NewList(fm)
	for _, v := range []struct{ name, value string }{
		{"raghuvanshy", "Gaurav Raghuvanshy"},
		{"cceckman", "Charles Eckman"},
		{"marcusolsson", "Marcus Olsson"},
	} {
		list.Commit(NewItem(list, v.name, v.value))
	}

	ui, err := tui.New(list)
	if err != nil {
		log.Fatal(err)
	}
	ui.SetTheme(Theme())
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	fm.Attach(ui)

	list.UI = ui

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
