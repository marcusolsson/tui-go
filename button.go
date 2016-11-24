package tui

import (
	"image"
	"strings"
)

var _ Widget = &Button{}

// Button is a widget to fill out space.
type Button struct {
	text string

	focused bool
	size    image.Point

	onActivated func(*Button)
}

// NewButton returns a new Button.
func NewButton(text string) *Button {
	return &Button{
		text: text,
	}
}

func withItemBrush(p *Painter, n string, fn func(*Painter)) {
	p.SetBrush(p.palette.Item(n).Fg, p.palette.Item(n).Bg)
	fn(p)
	p.RestoreBrush()
}

func withBrush(p *Painter, fg, bg Color, fn func(*Painter)) {
	p.SetBrush(fg, bg)
	fn(p)
	p.RestoreBrush()
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

// SizeHint returns the recommended size for the button.
func (b *Button) SizeHint() image.Point {
	var size image.Point

	lines := strings.Split(b.text, "\n")

	for _, line := range lines {
		if len(line) > size.X {
			size.X = len(line)
		}
	}
	size.Y = len(lines)

	return size
}

// SizePolicy returns the default layout behavior.
func (b *Button) SizePolicy() (SizePolicy, SizePolicy) {
	return Minimum, Minimum
}

// Resize updates the size of the button.
func (b *Button) Resize(size image.Point) {
	b.size = b.SizeHint()
}

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

func (b *Button) OnActivated(fn func(b *Button)) {
	b.onActivated = fn
}

func (b *Button) SetFocused(f bool) {
	b.focused = f
}
