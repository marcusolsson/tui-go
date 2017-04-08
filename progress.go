package tui

import "image"

var _ Widget = &Progress{}

// Progress is a widget to display a progress bar.
type Progress struct {
	size image.Point

	current, max int

	sizePolicyX SizePolicy
	sizePolicyY SizePolicy
}

// NewProgress returns a new Progress.
func NewProgress(max int) *Progress {
	return &Progress{
		max: max,
	}
}

// Draw draws the progress bar.
func (p *Progress) Draw(painter *Painter) {
	hpol, _ := p.SizePolicy()

	width := p.max
	if hpol == Expanding {
		width = p.Size().X
	}

	painter.DrawRune(0, 0, '[')
	painter.DrawRune(width-1, 0, ']')

	start := 1
	end := width - 1
	curr := int((float64(p.current) / float64(p.max)) * float64(end-start))

	for i := start; i < curr; i++ {
		painter.DrawRune(i, 0, '=')
	}
	for i := curr + start; i < end; i++ {
		painter.DrawRune(i, 0, '-')
	}
	painter.DrawRune(curr, 0, '>')
}

// Size returns the size of the progress bar.
func (p *Progress) Size() image.Point {
	return p.size
}

// MinSize returns the minimum size the widget is allowed to be.
func (p *Progress) MinSize() image.Point {
	return image.Point{5, 1}
}

// SizeHint returns the recommended size for the progress bar.
func (p *Progress) SizeHint() image.Point {
	return image.Point{p.max, 1}
}

// SizePolicy returns the default layout behavior.
func (p *Progress) SizePolicy() (SizePolicy, SizePolicy) {
	return p.sizePolicyX, p.sizePolicyY
}

// Resize updates the size of the progress bar.
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

func (p *Progress) OnEvent(_ Event) {
}

func (p *Progress) SetSizePolicy(horizontal, vertical SizePolicy) {
	p.sizePolicyX = horizontal
	p.sizePolicyY = vertical
}

func (p *Progress) SetCurrent(c int) {
	p.current = c
}
