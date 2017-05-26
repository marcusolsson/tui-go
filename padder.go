package tui

import "image"

var _ Widget = &Padder{}

// Padder is a widget to fill out space.
type Padder struct {
	widget Widget

	padding image.Point
}

// NewPadder returns a new Padder.
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

// MinSizeHint returns the minimum size the widget is allowed to be.
func (p *Padder) MinSizeHint() image.Point {
	return p.widget.MinSizeHint().Add(p.padding.Mul(2))
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

// OnEvent handles an event and propagates it the widget.
func (p *Padder) OnEvent(ev Event) {
	p.widget.OnEvent(ev)
}

// SetFocused set the focus on the widget.
func (p *Padder) SetFocused(f bool) {
	p.widget.SetFocused(f)
}

// IsFocused returns true if the widget is focused.
func (p *Padder) IsFocused() bool {
	return p.widget.IsFocused()
}
