package tui

import "image"

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
	Resize(size image.Point)

	OnEvent(ev Event)
}
