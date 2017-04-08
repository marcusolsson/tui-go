package main

import (
	"fmt"
	"time"

	"github.com/marcusolsson/tui-go"
)

type post struct {
	username string
	message  string
	time     string
}

var posts = []post{
	{username: "john", message: "hi, what's up?", time: "14:41"},
	{username: "jane", message: "not much", time: "14:43"},
}

func main() {
	channels := tui.NewVBox(
		tui.NewLabel("general"),
		tui.NewLabel("random"),
	)

	messages := tui.NewVBox(
		tui.NewLabel("slackbot"),
	)

	sidebar := tui.NewVBox(
		tui.NewLabel("CHANNELS"),
		channels,
		tui.NewLabel(""),
		tui.NewLabel("DIRECT MESSAGES"),
		messages,
	)
	sidebar.SetBorder(true)
	sidebar.SetSizePolicy(tui.Minimum, tui.Expanding)

	history := tui.NewVBox()
	history.SetBorder(true)
	history.SetSizePolicy(tui.Expanding, tui.Expanding)

	for _, m := range posts {
		b := tui.NewHBox(
			tui.NewLabel(m.time),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", m.username))),
			tui.NewLabel(m.message),
		)

		history.Append(b)
	}

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Minimum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Minimum)

	chat := tui.NewVBox(history, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {
		b := tui.NewHBox(
			tui.NewLabel(time.Now().Format("15:04")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", "john"))),
			tui.NewLabel(e.Text()),
		)

		history.Append(b)

		input.SetText("")
	})

	root := tui.NewHBox(sidebar, chat)
	root.SetSizePolicy(tui.Expanding, tui.Expanding)

	ui := tui.New(root)
	ui.SetKeybinding(tui.KeyEsc, func() {
		ui.Quit()
	})
	ui.SetKeybinding(tui.KeyArrowUp, func() {
		input.SetFocused(true)
	})
	ui.SetKeybinding(tui.KeyArrowDown, func() {
		input.SetFocused(false)
	})
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
