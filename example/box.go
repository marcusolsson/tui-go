package main

import "github.com/marcusolsson/tui-go"

func main() {
	root := tui.NewHorizontalBox(tui.NewLabel("Marcus"))

	if err := tui.New(root).Run(); err != nil {
		panic(err)
	}
}
