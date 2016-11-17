package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

// Label is a widget to display read-only text.
type Label struct {
	text string

	size image.Point
}

// NewLabel returns a new Label.
func NewLabel(text string) *Label {
	return &Label{
		text: text,
	}
}

// Draw draws the label.
func (l *Label) Draw(p *Painter) {
	p.DrawText(0, 0, l.text)
}

// Size returns the size of the label.
func (l *Label) Size() image.Point {
	return l.size
}

// SizeHint returns the recommended size for the label.
func (l *Label) SizeHint() image.Point {
	return image.Point{len(l.text), 1}
}

// SizePolicy returns the default layout behavior.
func (l *Label) SizePolicy() (SizePolicy, SizePolicy) {
	return Minimum, Minimum
}

// Resize updates the size of the label.
func (l *Label) Resize(_ image.Point) {
	l.size = l.SizeHint()
}

func (l *Label) OnEvent(_ termbox.Event) {
}

func (l *Label) IsVisible() bool {
	return true
}
