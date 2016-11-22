package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

// Progress is a widget to fill out space.
type Progress struct {
	size image.Point

	current, max int
}

// NewProgress returns a new Progress.
func NewProgress(max int) *Progress {
	return &Progress{
		max: max,
	}
}

// Draw draws the spacer.
func (p *Progress) Draw(painter *Painter) {
	for i := 0; i < p.current; i++ {
		painter.DrawRune(i, 0, '=')
	}
	for i := p.current; i < p.max; i++ {
		painter.DrawRune(i, 0, '-')
	}
}

// Size returns the size of the spacer.
func (p *Progress) Size() image.Point {
	return p.size
}

// SizeHint returns the recommended size for the spacer.
func (p *Progress) SizeHint() image.Point {
	return image.Point{p.max, 1}
}

// SizePolicy returns the default layout behavior.
func (p *Progress) SizePolicy() (SizePolicy, SizePolicy) {
	return Expanding, Minimum
}

// Resize updates the size of the spacer.
func (p *Progress) Resize(size image.Point) {
	hpol, vpol := p.SizePolicy()

	switch hpol {
	case Minimum:
		p.size.X = p.SizeHint().X
	case Expanding:
		p.size.X = size.X
	}

	switch vpol {
	case Minimum:
		p.size.Y = p.SizeHint().Y
	case Expanding:
		p.size.Y = size.Y
	}
}

func (p *Progress) OnEvent(_ termbox.Event) {
}

func (p *Progress) IsVisible() bool {
	return true
}

func (p *Progress) SetCurrent(c int) {
	p.current = c
}
