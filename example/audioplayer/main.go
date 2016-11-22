package main

import "github.com/marcusolsson/tui-go"

type song struct {
	artist string
	album  string
	track  string
}

var songs = []song{
	{artist: "DJ Example", album: "Mixtape 3", track: "Testing stuff"},
	{artist: "DJ Example", album: "Mixtape 3", track: "Breaking stuff"},
	{artist: "DJ Example", album: "Initial commit", track: "****, not again ..."},
}

func main() {
	library := tui.NewTable(0, 0)
	library.SetSizePolicy(tui.Expanding, tui.Expanding)
	library.SetColumnStretch(0, 1)
	library.SetColumnStretch(1, 1)
	library.SetColumnStretch(2, 2)

	library.AppendRow(
		tui.NewLabel("ARTIST"),
		tui.NewLabel("ALBUM"),
		tui.NewLabel("TRACK"),
	)

	for _, s := range songs {
		library.AppendRow(
			tui.NewLabel(s.artist),
			tui.NewLabel(s.album),
			tui.NewLabel(s.track),
		)
	}

	progress := tui.NewProgress(100)
	progress.SetCurrent(30)

	controls := tui.NewHBox(progress)
	controls.SetSizePolicy(tui.Expanding, tui.Minimum)

	status := tui.NewStatusBar("05:12 - 07:46")
	status.SetPermanentText(`VOLUME 68%`)

	root := tui.NewVBox(library, controls, status)
	root.SetSizePolicy(tui.Expanding, tui.Expanding)

	if err := tui.New(root).Run(); err != nil {
		panic(err)
	}
}
