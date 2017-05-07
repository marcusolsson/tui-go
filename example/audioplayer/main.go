package main

import (
	"fmt"
	"time"

	"github.com/marcusolsson/tui-go"
)

type song struct {
	artist   string
	album    string
	track    string
	duration time.Duration
}

var songs = []song{
	{artist: "DJ Example", album: "Mixtape 3", track: "Testing stuff", duration: 110 * time.Second},
	{artist: "DJ Example", album: "Mixtape 3", track: "Breaking stuff", duration: 140 * time.Second},
	{artist: "DJ Example", album: "Initial commit", track: "****, not again ...", duration: 150 * time.Second},
}

func main() {
	var p player

	progress := tui.NewProgress(100)

	library := tui.NewTable(0, 0)
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

	status := tui.NewStatusBar("")
	status.SetPermanentText(`VOLUME 68%`)

	library.OnItemActivated(func(t *tui.Table) {
		p.play(songs[t.Selected()-1], func(curr, max int) {
			progress.SetCurrent(curr)
			progress.SetMax(max)

			status.SetText(fmt.Sprintf("%s / %s", time.Duration(curr)*time.Second, time.Duration(max)*time.Second))

			tui.Repaint()
		})
	})

	root := tui.NewVBox(
		library,
		tui.NewSpacer(),
		progress,
		status,
	)

	ui := tui.New(root)
	ui.SetKeybinding(tui.KeyEsc, func() { ui.Quit() })
	ui.SetKeybinding('q', func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}

type player struct {
	elapsed int
	total   int
	quit    chan struct{}
}

func (p *player) play(s song, callback func(current, max int)) {
	if p.quit != nil {
		p.quit <- struct{}{}
		<-p.quit
	}

	p.quit = make(chan struct{})
	p.total = int(s.duration.Seconds())
	p.elapsed = 0

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				if p.elapsed >= p.total {
					p.quit <- struct{}{}
				}

				callback(p.elapsed, p.total)
				p.elapsed++
			case <-p.quit:
				p.quit <- struct{}{}
				return
			}
		}
	}()
}
