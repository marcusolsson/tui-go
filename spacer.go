package tui

import "image"

var _ Widget = &Spacer{}

// Spacer is a widget to fill out space.
type Spacer struct {
	size image.Point
}

// NewSpacer returns a new Spacer.
func NewSpacer() *Spacer {
	return &Spacer{}
}

// Draw draws the spacer.
func (s *Spacer) Draw(p *Painter) {
}

// Size returns the size of the spacer.
func (s *Spacer) Size() image.Point {
	return s.size
}

// MinSize returns the minimum size the widget is allowed to be.
func (s *Spacer) MinSize() image.Point {
	return image.Point{}
}

// SizeHint returns the recommended size for the spacer.
func (s *Spacer) SizeHint() image.Point {
	return image.Point{}
}

// SizePolicy returns the default layout behavior.
func (s *Spacer) SizePolicy() (SizePolicy, SizePolicy) {
	return Expanding, Expanding
}

// Resize updates the size of the spacer.
func (s *Spacer) Resize(size image.Point) {
	s.size = size
}

func (s *Spacer) OnEvent(_ Event) {
}
