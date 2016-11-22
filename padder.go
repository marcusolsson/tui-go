package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

// Padder is a widget to fill out space.
type Padder struct {
	widget Widget

	padding image.Point
}

// Padder returns a new Padder.
func NewPadder(x, y int, w Widget) *Padder {
	return &Padder{
		widget:  w,
		padding: image.Point{x, y},
	}
}

// Draw draws the padded widget.
func (p *Padder) Draw(painter *Painter) {
	painter.Translate(p.padding.X, p.padding.Y)
	defer painter.Restore()

	p.widget.Draw(painter)
}

// Size returns the size of the padded widget.
func (p *Padder) Size() image.Point {
	return p.widget.Size().Add(p.padding.Mul(2))
}

// SizeHint returns the recommended size for the padded widget.
func (p *Padder) SizeHint() image.Point {
	return p.widget.SizeHint().Add(p.padding.Mul(2))
}

// SizePolicy returns the default layout behavior.
func (p *Padder) SizePolicy() (SizePolicy, SizePolicy) {
	return p.widget.SizePolicy()
}

// Resize updates the size of the padded widget.
func (p *Padder) Resize(size image.Point) {
	p.widget.Resize(size.Sub(p.padding.Mul(2)))
}

func (p *Padder) OnEvent(ev termbox.Event) {
	p.widget.OnEvent(ev)
}
