package main

import "github.com/marcusolsson/tui-go"

func main() {
	authors := tui.NewList()
	authors.AddItems(
		"William Shakespeare",
		"Charles Dickens",
		"Jane Austen",
		"George Orwell",
	)
	authors.SetSelected(0)
	authors.SetRows(10)

	root := tui.NewHorizontalBox(authors)

	if err := tui.New(root).Run(); err != nil {
		panic(err)
	}
}
