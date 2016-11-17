package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

type SizePolicy int

const (
	Minimum SizePolicy = iota
	Expanding
)

// Widget defines common operations on widgets.
type Widget interface {
	Draw(p *Painter)

	Size() image.Point
	SizeHint() image.Point
	SizePolicy() (SizePolicy, SizePolicy)
	Resize(contentSize image.Point)

	HandleEvent(ev termbox.Event)
	IsVisible() bool
}
