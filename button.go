package tui

import (
	"image"
	"strings"
)

var _ Widget = &Button{}

// Button is a widget that can be activated to perform some action, or to
// answer a question. It displays a label that can be activated.
type Button struct {
	text string

	focused bool
	size    image.Point

	onActivated func(*Button)

	sizePolicyX SizePolicy
	sizePolicyY SizePolicy
}

// NewButton returns a new Button with the specified label.
func NewButton(text string) *Button {
	return &Button{
		text: text,
	}
}

// Draw draws the button.
func (b *Button) Draw(p *Painter) {
	s := b.Size()

	style := "button"
	if b.focused {
		style += ".focused"
	}

	p.WithStyledBrush(style, func(p *Painter) {
		lines := strings.Split(b.text, "\n")
		for i, line := range lines {
			p.FillRect(0, i, s.X, 1)
			p.DrawText(0, i, line)
		}
	})
}

// Size returns the size of the button.
func (b *Button) Size() image.Point {
	return b.size
}

// MinSizeHint returns the minimum size the widget is allowed to be.
func (b *Button) MinSizeHint() image.Point {
	return image.Point{1, 1}
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

// SizePolicy returns the default layout behavior.
func (b *Button) SizePolicy() (SizePolicy, SizePolicy) {
	return b.sizePolicyX, b.sizePolicyY
}

// Resize updates the size of the button.
func (b *Button) Resize(size image.Point) {
	b.size = size
}

// OnEvent handles terminal events.
func (b *Button) OnEvent(ev Event) {
	if !b.focused {
		return
	}

	if ev.Type != EventKey {
		return
	}

	switch ev.Key {
	case KeyEnter:
		if b.onActivated != nil {
			b.onActivated(b)
		}
	}
}

// OnActivated sets a function to be run whenever the button is activated.
func (b *Button) OnActivated(fn func(b *Button)) {
	b.onActivated = fn
}

// SetFocused focuses this button.
func (b *Button) SetFocused(f bool) {
	b.focused = f
}

// SetSizePolicy sets the size policy for each axis.
func (b *Button) SetSizePolicy(horizontal, vertical SizePolicy) {
	b.sizePolicyX = horizontal
	b.sizePolicyY = vertical
}
