package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

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

func (s *Spacer) OnEvent(_ termbox.Event) {
}

func (s *Spacer) IsVisible() bool {
	return true
}
