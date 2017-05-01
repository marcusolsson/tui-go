package tui

import (
	"image"
	"strings"
)

var _ Widget = &Button{}

// Button is a widget that can be activated to perform some action, or to
// answer a question. It displays a label that can be activated.
type Button struct {
	WidgetBase

	text string

	onActivated func(*Button)
}

// NewButton returns a new Button with the specified label.
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

// SizeHint returns the recommended size for the button.
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

// OnEvent handles terminal events.
func (b *Button) OnEvent(ev Event) {
	if !b.IsFocused() || ev.Type != EventKey {
		return
	}
	if ev.Key == KeyEnter && b.onActivated != nil {
		b.onActivated(b)
	}
}

// OnActivated sets a function to be run whenever the button is activated.
func (b *Button) OnActivated(fn func(b *Button)) {
	b.onActivated = fn
}
