package tui

import "image"

var _ Widget = &Spacer{}

// Spacer is a widget to fill out space.
type Spacer struct {
	WidgetBase
}

// NewSpacer returns a new Spacer.
func NewSpacer() *Spacer {
	return &Spacer{}
}

// MinSizeHint returns the minimum size the widget is allowed to be.
func (s *Spacer) MinSizeHint() image.Point {
	return image.ZP
}

// SizeHint returns the recommended size for the spacer.
func (s *Spacer) SizeHint() image.Point {
	return image.ZP
}

// SizePolicy returns the default layout behavior.
func (s *Spacer) SizePolicy() (SizePolicy, SizePolicy) {
	return Expanding, Expanding
}
