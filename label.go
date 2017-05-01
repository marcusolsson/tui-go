package tui

import (
	"image"
	"strings"

	wordwrap "github.com/mitchellh/go-wordwrap"
)

var _ Widget = &Label{}

// Label is a widget to display read-only text.
type Label struct {
	WidgetBase

	text     string
	wordWrap bool
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

	if l.wordWrap {
		lines = strings.Split(wordwrap.WrapString(l.text, uint(l.Size().X)), "\n")
	}
	for i, line := range lines {
		p.DrawText(0, i, line)
	}
}

// MinSizeHint returns the minimum size the widget is allowed to be.
func (l *Label) MinSizeHint() image.Point {
	return image.Point{1, 1}
}

// SizeHint returns the recommended size for the label.
func (l *Label) SizeHint() image.Point {
	var max int
	lines := strings.Split(l.text, "\n")
	for _, line := range lines {
		if w := stringWidth(line); w > max {
			max = w
		}
	}
	return image.Point{max, l.heightForWidth(max)}
}

func (l *Label) heightForWidth(w int) int {
	return len(strings.Split(wordwrap.WrapString(l.text, uint(w)), "\n"))
}

// SetText sets the text content of the label.
func (l *Label) SetText(text string) {
	l.text = text
}

// SetWordWrap sets whether text content should be wrapped.
func (l *Label) SetWordWrap(enabled bool) {
	l.wordWrap = enabled
}
