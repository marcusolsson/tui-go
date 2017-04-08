package tui

import "image"

type SizePolicy int

const (
	Preferred SizePolicy = iota
	Minimum
	Maximum
	Expanding
)

// Widget defines common operations on widgets.
type Widget interface {
	Draw(p *Painter)
	MinSizeHint() image.Point
	Size() image.Point
	SizeHint() image.Point
	SizePolicy() (SizePolicy, SizePolicy)
	Resize(size image.Point)
	OnEvent(ev Event)
}
