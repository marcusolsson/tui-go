package tui

import (
	"image"
)

var _ Widget = &ScrollArea{}

// ScrollArea is a widget to fill out space.
type ScrollArea struct {
	WidgetBase

	Widget Widget

	topLeft    image.Point
	autoscroll bool
}

// NewScrollArea returns a new ScrollArea.
func NewScrollArea(w Widget) *ScrollArea {
	return &ScrollArea{
		Widget: w,
	}
}

// MinSizeHint returns the minimum size the widget is allowed to be.
func (s *ScrollArea) MinSizeHint() image.Point {
	return image.ZP
}

// SizeHint returns the size hint of the underlying widget.
func (s *ScrollArea) SizeHint() image.Point {
	return image.Pt(15, 8)
}

// Scroll shifts the views over the content.
func (s *ScrollArea) Scroll(dx, dy int) {
	s.topLeft.X += dx
	s.topLeft.Y += dy
}

// ScrollToBottom ensures the bottom-most part of the scroll area is visible.
func (s *ScrollArea) ScrollToBottom() {
	s.topLeft.Y = s.Widget.SizeHint().Y - s.Size().Y
}

// ScrollToTop resets the vertical scroll position.
func (s *ScrollArea) ScrollToTop() {
	s.topLeft.Y = 0
}

// SetAutoscrollToBottom makes sure the content is scrolled to bottom on resize.
func (s *ScrollArea) SetAutoscrollToBottom(autoscroll bool) {
	s.autoscroll = autoscroll
}

// Draw draws the scroll area.
func (s *ScrollArea) Draw(p *Painter) {
	p.Translate(-s.topLeft.X, -s.topLeft.Y)
	defer p.Restore()

	off := image.Point{s.topLeft.X, s.topLeft.Y}
	p.WithMask(image.Rectangle{Min: off, Max: s.Size().Add(off)}, func(p *Painter) {
		s.Widget.Draw(p)
	})
}

// Resize resizes the scroll area and the underlying widget.
func (s *ScrollArea) Resize(size image.Point) {
	s.Widget.Resize(s.Widget.SizeHint())
	s.WidgetBase.Resize(size)

	if s.autoscroll {
		s.ScrollToBottom()
	}
}
