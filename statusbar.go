package tui

import (
	"image"
)

var _ Widget = &StatusBar{}

// StatusBar is a widget to display status information.
type StatusBar struct {
	WidgetBase

	text     string
	permText string
}

// NewStatusBar returns a new StatusBar.
func NewStatusBar(text string) *StatusBar {
	return &StatusBar{
		text:     text,
		permText: "",
	}
}

// Draw draws the status bar.
func (b *StatusBar) Draw(p *Painter) {
	p.WithStyle("statusbar", func(p *Painter) {
		p.FillRect(0, 0, b.Size().X, 1)
		p.DrawText(0, 0, b.text)
		p.DrawText(b.Size().X-stringWidth(b.permText), 0, b.permText)
	})
}

// SizeHint returns the recommended size for the status bar.
func (b *StatusBar) SizeHint() image.Point {
	return image.Point{10, 1}
}

// SizePolicy returns the default layout behavior.
func (b *StatusBar) SizePolicy() (SizePolicy, SizePolicy) {
	return Preferred, Maximum
}

// SetText sets the text content of the status bar.
func (b *StatusBar) SetText(text string) {
	b.text = text
}

// SetPermanentText sets the permanent text of the status bar.
func (b *StatusBar) SetPermanentText(text string) {
	b.permText = text
}
