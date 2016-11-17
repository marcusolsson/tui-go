package main

import "github.com/marcusolsson/tui-go"

func main() {
	table := tui.NewGrid(0, 0)

	table.AppendRow(tui.NewLabel("AUTHOR"), tui.NewLabel("BOOK"))
	table.AppendRow(tui.NewLabel("William Shakespeare"), tui.NewLabel("Macbeth"))
	table.AppendRow(tui.NewLabel("George Orwell"), tui.NewLabel("Animal Farm"))

	nested := tui.NewVerticalBox(
		tui.NewLabel("War of the Worlds"),
		tui.NewLabel("Time Machine"),
		tui.NewLabel("Invisible Man"),
	)
	nested.SetBorder(true)
	nested.SetSizePolicy(tui.Expanding, tui.Minimum)

	table.AppendRow(tui.NewLabel("H.G. Wells"), nested)
	table.AppendRow(tui.NewLabel("Jane Austen"), tui.NewLabel("Pride and Prejudice"))
	table.AppendRow(tui.NewLabel("William Shakespeare"), tui.NewLabel("Macbeth"))
	table.AppendRow(tui.NewLabel("Charles Dickens"))

	if err := tui.New(table).Run(); err != nil {
		panic(err)
	}
}
