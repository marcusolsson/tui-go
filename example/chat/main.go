package main

import (
	"image"

	"github.com/marcusolsson/tui-go"
)

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

	for _, m := range []string{"hi, what's up?", "not much"} {
		b := tui.NewHBox(
			tui.NewLabel("john:"), tui.NewLabel(m),
		)

		history.Append(b)
	}

	input := tui.NewEntry()
	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Minimum)

	chat := tui.NewVBox(history, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {
		b := tui.NewHBox(
			tui.NewLabel("john:"),
			tui.NewPadder(tui.NewLabel(e.Text()), image.Point{1, 0}),
		)

		history.Append(b)

		input.SetText("")
	})

	root := tui.NewHBox(sidebar, chat)
	root.SetSizePolicy(tui.Expanding, tui.Expanding)

	if err := tui.New(root).Run(); err != nil {
		panic(err)
	}
}
