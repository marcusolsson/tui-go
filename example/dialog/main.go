package main

import (
	tui "github.com/marcusolsson/tui-go"
)

func main() {
	ui, err := tui.New(tui.NewVBox())
	if err != nil {
		panic(err)
	}

	d := tui.NewDialog()
	d.SetTitle("Really quit?")
	d.Append(tui.NewLabel("Are you sure you want to quit?"))

	d.OnFinished(func(result tui.DialogResult) {
		switch result {
		case tui.DialogAccepted:
			ui.Quit()
		case tui.DialogRejected:
			ui.HideDialog()
		}
	})

	ui.SetKeybinding("Ctrl+W", func() { ui.ShowDialog(d) })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
