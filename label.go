package tui

import (
	"image"
	"strings"

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
	lines := strings.Split(l.text, "\n")
	for i, line := range lines {
		p.DrawText(0, i, line)
	}
}

// Size returns the size of the label.
func (l *Label) Size() image.Point {
	return l.size
}

// SizeHint returns the recommended size for the label.
func (l *Label) SizeHint() image.Point {
	var size image.Point

	lines := strings.Split(l.text, "\n")

	for _, line := range lines {
		if len(line) > size.X {
			size.X = len(line)
		}
	}
	size.Y = len(lines)

	return size
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

func (l *Label) SetText(text string) {
	l.text = text
}
