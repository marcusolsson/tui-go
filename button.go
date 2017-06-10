package tui

import (
	"image"
	"strings"
)

var _ Widget = &Button{}

// Button is a widget that can be activated to perform some action, or to
// answer a question.
type Button struct {
	WidgetBase

	text string

	onActivated func(*Button)
}

// NewButton returns a new Button with the given text as the label.
func NewButton(text string) *Button {
	return &Button{
		text: text,
	}
}

// Draw draws the button.
func (b *Button) Draw(p *Painter) {
	style := "button"
	if b.IsFocused() {
		style += ".focused"
	}
	p.WithStyle(style, func(p *Painter) {
		lines := strings.Split(b.text, "\n")
		for i, line := range lines {
			p.FillRect(0, i, b.Size().X, 1)
			p.DrawText(0, i, line)
		}
	})
}

// SizeHint returns the recommended size hint for the button.
func (b *Button) SizeHint() image.Point {
	if len(b.text) == 0 {
		return b.MinSizeHint()
	}

	var size image.Point
	lines := strings.Split(b.text, "\n")
	for _, line := range lines {
		if w := stringWidth(line); w > size.X {
			size.X = w
		}
	}
	size.Y = len(lines)

	return size
}

// OnKeyEvent handles keys events.
func (b *Button) OnKeyEvent(ev KeyEvent) {
	if !b.IsFocused() {
		return
	}
	if ev.Key == KeyEnter && b.onActivated != nil {
		b.onActivated(b)
	}
}

// OnActivated allows a custom function to be run whenever the button is activated.
func (b *Button) OnActivated(fn func(b *Button)) {
	b.onActivated = fn
}
