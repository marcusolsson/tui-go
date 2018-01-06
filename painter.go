package tui

import (
	"image"
)

type surfaceCell struct {
	ch    rune
	style Style
}

type surfaceBuffer struct {
	cells map[image.Point]surfaceCell
}

func (s *surfaceBuffer) SetCell(x, y int, ch rune, style Style) {
	s.cells[image.Pt(x, y)] = surfaceCell{
		ch:    ch,
		style: style,
	}
}

// Surface defines a surface that can be painted on.
type Surface interface {
	SetCell(x, y int, ch rune, s Style)
	SetCursor(x, y int)
	HideCursor()
	Begin()
	End()
	Size() image.Point
}

// Painter provides operations to paint on a surface.
type Painter struct {
	theme *Theme

	buffer surfaceBuffer

	// Surface to paint on.
	surface Surface

	// Current brush.
	style Style

	// Transform stack
	transforms []image.Point

	mask image.Rectangle
}

// NewPainter returns a new instance of Painter.
func NewPainter(s Surface, p *Theme) *Painter {
	return &Painter{
		theme:   p,
		buffer:  surfaceBuffer{make(map[image.Point]surfaceCell)},
		surface: s,
		style:   p.Style("normal"),
		mask: image.Rectangle{
			Min: image.ZP,
			Max: s.Size(),
		},
	}
}

// Translate pushes a new translation transform to the stack.
func (p *Painter) Translate(x, y int) {
	p.transforms = append(p.transforms, image.Point{x, y})
}

// Restore pops the latest transform from the stack.
func (p *Painter) Restore() {
	if len(p.transforms) > 0 {
		p.transforms = p.transforms[:len(p.transforms)-1]
	}
}

// Flush writes buffers to the surface.
func (p *Painter) Flush() {
	p.surface.Begin()
	for k, v := range p.buffer.cells {
		p.surface.SetCell(k.X, k.Y, v.ch, v.style)
	}
	p.surface.End()
}

// DrawCursor draws the cursor at the given position.
func (p *Painter) DrawCursor(x, y int) {
	wp := p.mapLocalToWorld(image.Point{x, y})
	p.surface.SetCursor(wp.X, wp.Y)
}

// DrawRune paints a rune at the given coordinate.
func (p *Painter) DrawRune(x, y int, r rune) {
	wp := p.mapLocalToWorld(image.Point{x, y})
	if (p.mask.Min.X <= wp.X) && (wp.X < p.mask.Max.X) && (p.mask.Min.Y <= wp.Y) && (wp.Y < p.mask.Max.Y) {
		p.buffer.SetCell(wp.X, wp.Y, r, p.style)
	}
}

// DrawText paints a string starting at the given coordinate.
func (p *Painter) DrawText(x, y int, text string) {
	for _, r := range text {
		p.DrawRune(x, y, r)
		x += runeWidth(r)
	}
}

// DrawHorizontalLine paints a horizontal line using box characters.
func (p *Painter) DrawHorizontalLine(x1, x2, y int) {
	for x := x1; x < x2; x++ {
		p.DrawRune(x, y, '─')
	}
}

// DrawVerticalLine paints a vertical line using box characters.
func (p *Painter) DrawVerticalLine(x, y1, y2 int) {
	for y := y1; y < y2; y++ {
		p.DrawRune(x, y, '│')
	}
}

// DrawRect paints a rectangle using box characters.
func (p *Painter) DrawRect(x, y, w, h int) {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			m := i + x
			n := j + y

			switch {
			case i == 0 && j == 0:
				p.DrawRune(m, n, '┌')
			case i == w-1 && j == 0:
				p.DrawRune(m, n, '┐')
			case i == 0 && j == h-1:
				p.DrawRune(m, n, '└')
			case i == w-1 && j == h-1:
				p.DrawRune(m, n, '┘')
			case i == 0 || i == w-1:
				p.DrawRune(m, n, '│')
			case j == 0 || j == h-1:
				p.DrawRune(m, n, '─')
			}
		}
	}
}

// FillRect clears a rectangular area with whitespace.
func (p *Painter) FillRect(x, y, w, h int) {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			p.DrawRune(i+x, j+y, ' ')
		}
	}
}

// SetStyle sets the style used when painting.
func (p *Painter) SetStyle(s Style) {
	p.style = s
}

// WithStyle executes the provided function with the named Style applied on top of the current one.
func (p *Painter) WithStyle(n string, fn func(*Painter)) {
	prev := p.style
	new := prev.mergeIn(p.theme.Style(n))
	p.SetStyle(new)
	fn(p)
	p.SetStyle(prev)
}

// WithMask masks a painter to restrict painting within the given rectangle.
func (p *Painter) WithMask(r image.Rectangle, fn func(*Painter)) {
	tmp := p.mask
	defer func() { p.mask = tmp }()

	p.mask = p.mask.Intersect(image.Rectangle{
		Min: p.mapLocalToWorld(r.Min),
		Max: p.mapLocalToWorld(r.Max),
	})

	fn(p)
}

func (p *Painter) mapLocalToWorld(point image.Point) image.Point {
	var offset image.Point
	for _, s := range p.transforms {
		offset = offset.Add(s)
	}
	return point.Add(offset)
}
