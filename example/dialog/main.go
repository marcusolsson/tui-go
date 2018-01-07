package main

import (
	"log"

	tui "github.com/marcusolsson/tui-go"
)

func main() {
	left := tui.NewVBox(tui.NewLabel("Left column"))
	left.SetBorder(true)

	right := tui.NewVBox(tui.NewLabel("Right column"))
	right.SetBorder(true)

	root := tui.NewHBox(left, right)

	ui, err := tui.New(root)
	if err != nil {
		panic(err)
	}

	d := tui.NewDialog()
	d.SetTitle("Really quit?")

	var (
		ok      = tui.NewButton("OK")
		padding = tui.NewLabel(" ")
		cancel  = tui.NewButton("Cancel")
	)

	btngrp := tui.NewHBox(tui.NewSpacer(), ok, padding, cancel)

	d.Append(
		tui.NewLabel("Are you sure you want to quit?"),
		tui.NewLabel(""),
		btngrp,
	)

	d.OnFinished(func(result tui.DialogResult) {
		switch result {
		case tui.DialogAccepted:
			ui.Quit()
		case tui.DialogRejected:
			ui.HideDialog()
		}
	})

	ui.SetKeybinding("Ctrl+X", func() {
		var chain tui.SimpleFocusChain
		chain.Set(ok, cancel)
		ui.SetFocusChain(&chain)

		ui.ShowDialog(d)
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
