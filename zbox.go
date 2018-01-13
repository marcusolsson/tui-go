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
