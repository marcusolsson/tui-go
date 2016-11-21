package tui

import (
	"image"
	"strings"

	termbox "github.com/nsf/termbox-go"
)

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

func withBrush(p *Painter, fg, bg termbox.Attribute, fn func(*Painter)) {
	p.SetBrush(fg, bg)
	fn(p)
	p.SetBrush(termbox.ColorDefault, termbox.ColorDefault)
}

// Draw draws the button.
func (b *Button) Draw(p *Painter) {
	s := b.Size()

	var fg, bg termbox.Attribute

	if b.focused {
		fg = termbox.ColorBlack
		bg = termbox.ColorWhite
	}

	withBrush(p, fg, bg, func(p *Painter) {
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

func (b *Button) OnEvent(ev termbox.Event) {
	if !b.focused {
		return
	}

	switch ev.Key {
	case termbox.KeyEnter:
		if b.onActivated != nil {
			b.onActivated(b)
		}
	}
}

func (b *Button) OnActivated(fn func(b *Button)) {
	b.onActivated = fn
}

func (b *Button) IsVisible() bool {
	return true
}

func (b *Button) SetFocused(f bool) {
	b.focused = f
}
