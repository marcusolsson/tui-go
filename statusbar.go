package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

var _ Widget = &StatusBar{}

// StatusBar is a widget to display status information.
type StatusBar struct {
	size image.Point

	text     string
	permText string

	fg, bg termbox.Attribute
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
	s := b.Size()

	p.WithStyle("statusbar", func(p *Painter) {
		p.FillRect(0, 0, s.X, 1)
		p.DrawText(0, 0, b.text)
		p.DrawText(s.X-stringWidth(b.permText), 0, b.permText)
	})
}

// Size returns the size of the status bar.
func (b *StatusBar) Size() image.Point {
	return b.size
}

// MinSizeHint returns the minimum size the widget is allowed to be.
func (b *StatusBar) MinSizeHint() image.Point {
	return image.Point{1, 1}
}

// SizeHint returns the recommended size for the status bar.
func (b *StatusBar) SizeHint() image.Point {
	return image.Point{10, 1}
}

// SizePolicy returns the default layout behavior.
func (b *StatusBar) SizePolicy() (SizePolicy, SizePolicy) {
	return Expanding, Minimum
}

// Resize updates the size of the status bar.
func (b *StatusBar) Resize(size image.Point) {
	b.size = size
}

func (b *StatusBar) OnEvent(_ Event) {
}

func (b *StatusBar) SetBrush(fg, bg termbox.Attribute) {
	b.fg = fg
	b.bg = bg
}

func (b *StatusBar) SetText(text string) {
	b.text = text
}

func (b *StatusBar) SetPermanentText(text string) {
	b.permText = text
}
