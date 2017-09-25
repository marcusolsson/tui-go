package tui

import "image"

var _ Widget = &Progress{}

// Progress is a widget to display a progress bar.
type Progress struct {
	WidgetBase

	current, max int
}

// NewProgress returns a new Progress.
func NewProgress(max int) *Progress {
	return &Progress{
		max: max,
	}
}

// Draw draws the progress bar.
func (p *Progress) Draw(painter *Painter) {
	painter.DrawRune(0, 0, '[')
	painter.DrawRune(p.Size().X-1, 0, ']')

	start := 1
	end := p.Size().X - 1
	curr := int((float64(p.current) / float64(p.max)) * float64(end-start))

	for i := start; i < curr; i++ {
		painter.DrawRune(i, 0, '=')
	}
	for i := curr + start; i < end; i++ {
		painter.DrawRune(i, 0, '-')
	}
	painter.DrawRune(curr, 0, '>')
}

// MinSizeHint returns the minimum size the widget is allowed to be.
func (p *Progress) MinSizeHint() image.Point {
	return image.Point{5, 1}
}

// SizeHint returns the recommended size for the progress bar.
func (p *Progress) SizeHint() image.Point {
	return image.Point{p.max, 1}
}

// SetCurrent sets the current progress.
func (p *Progress) SetCurrent(c int) {
	p.current = c
}

// SetMax sets the maximum progress.
func (p *Progress) SetMax(m int) {
	p.max = m
}
