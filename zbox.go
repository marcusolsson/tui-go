package tui

import (
	"image"
)

// A ZBox is a stack of Widgets that have the same X and Y dimensions, where
// Widgets on top are rendered over those on the bottom.
// It can be used to implement modal dialogs.
type ZBox struct {
	WidgetBase

	contents []Widget
}

func NewZBox(contents ...Widget) *ZBox {
	return &ZBox{
		contents: contents,
	}
}

func (z *ZBox) Draw(p *Painter) {
	for _, r := range z.contents {
		r.Draw(p)
	}
}

func (z *ZBox) Resize(size image.Point) {
	for _, w := range z.contents {
		w.Resize(size)
	}
}

// max is a helper to get the maximum image.Point out of the contents.
func (z *ZBox) max(f func(w Widget) image.Point) image.Point {
	r := image.ZP
	for _, w := range z.contents {
		hint := f(w)
		if hint.X > r.X {
			r.X = hint.X
		}
		if hint.Y > r.Y {
			r.Y = hint.Y
		}
	}
	return r
}

func (z *ZBox) SizeHint() image.Point {
	return z.max(func(w Widget) image.Point{
		return w.SizeHint()
	})
}

func (z *ZBox) MinSizeHint() image.Point {
	return z.max(func(w Widget) image.Point{
		return w.MinSizeHint()
	})
}

func (z *ZBox) Size() image.Point {
	return z.max(func(w Widget) image.Point{
		return w.Size()
	})
}

func (z *ZBox) SizePolicy() (SizePolicy, SizePolicy) {
	return Expanding, Expanding
}
