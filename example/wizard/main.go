package main

import (
	"image"

	"github.com/marcusolsson/tui-go"
)

var logo = `     _____ __ ____  ___   ______________  
    / ___// //_/\ \/ / | / / ____/_  __/  
    \__ \/ ,<    \  /  |/ / __/   / /     
   ___/ / /| |   / / /|  / /___  / /      
  /____/_/ |_|  /_/_/ |_/_____/ /_/     `

func main() {
	user := tui.NewEntry()
	user.SetFocused(true)

	password := tui.NewEntry()

	form := tui.NewGrid(0, 0)
	form.SetSizePolicy(tui.Expanding, tui.Minimum)

	form.AppendRow(tui.NewLabel("User"), tui.NewLabel("Password"))
	form.AppendRow(user, password)

	login := tui.NewButton("[Login]")
	login.SetFocused(true)

	register := tui.NewButton("[Register]")

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(login, image.Point{1, 0}),
		tui.NewPadder(register, image.Point{1, 0}),
	)
	buttons.SetSizePolicy(tui.Expanding, tui.Minimum)

	window := tui.NewVBox(
		tui.NewPadder(tui.NewLabel(logo), image.Point{10, 1}),
		tui.NewPadder(tui.NewLabel("Welcome to Skynet! Login or register."), image.Point{12, 0}),
		tui.NewPadder(form, image.Point{1, 1}),
		buttons,
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	wrapper.SetSizePolicy(tui.Minimum, tui.Expanding)

	status := tui.NewStatusBar("Ready.")
	login.OnActivated(func(b *tui.Button) {
		status.SetText("Logged in.")
	})

	content := tui.NewHBox(
		tui.NewSpacer(),
		wrapper,
		tui.NewSpacer(),
	)
	content.SetSizePolicy(tui.Expanding, tui.Expanding)

	root := tui.NewVBox(content, status)
	root.SetSizePolicy(tui.Expanding, tui.Expanding)

	if err := tui.New(root).Run(); err != nil {
		panic(err)
	}
}
