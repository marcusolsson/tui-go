package tui

import (
	"image"
)

// Surface defines a surface that can be painted on.
type Surface interface {
	SetCell(x, y int, ch rune, s Style)
	SetCursor(x, y int)
	Begin()
	End()
	Size() image.Point
}

// Painter provides operations to paint on a surface.
type Painter struct {
	Theme *Theme

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
		Theme:   p,
		surface: s,
		style:   p.Style("normal"),
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

// Begin prepares the surface for painting.
func (p *Painter) Begin() {
	p.surface.Begin()
}

// End finalizes any painting that has been made.
func (p *Painter) End() {
	p.surface.End()
}

// Repaint clears the surface, draws the scene and flushes it.
func (p *Painter) Repaint(w Widget) {
	p.Begin()
	w.Resize(p.surface.Size())
	w.Draw(p)
	p.End()
}

func (p *Painter) DrawCursor(x, y int) {
	wp := p.mapLocalToWorld(image.Point{x, y})
	p.surface.SetCursor(wp.X, wp.Y)
}

// DrawRune paints a rune at the given coordinate.
func (p *Painter) DrawRune(x, y int, r rune) {
	// If a mask is set, only draw if the mask contains the coordinate.
	if p.mask != image.ZR {
		if (x < p.mask.Min.X) || (x > p.mask.Max.X) ||
			(y < p.mask.Min.Y) || (y > p.mask.Max.Y) {
			return
		}
	}
	wp := p.mapLocalToWorld(image.Point{x, y})
	p.surface.SetCell(wp.X, wp.Y, r, p.style)
}

// DrawText paints a string starting at the given coordinate.
func (p *Painter) DrawText(x, y int, text string) {
	for _, r := range text {
		p.DrawRune(x, y, r)
		x += runeWidth(r)
	}
}

func (p *Painter) DrawHorizontalLine(x1, x2, y int) {
	for x := x1; x < x2; x++ {
		p.DrawRune(x, y, '─')
	}
}

func (p *Painter) DrawVerticalLine(x, y1, y2 int) {
	for y := y1; y < y2; y++ {
		p.DrawRune(x, y, '│')
	}
}

// DrawRect paints a rectangle.
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

func (p *Painter) FillRect(x, y, w, h int) {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			p.DrawRune(i+x, j+y, ' ')
		}
	}
}

func (p *Painter) SetStyle(s Style) {
	p.style = s
}

func (p *Painter) RestoreStyle() {
	p.SetStyle(p.Theme.Style("normal"))
}

func (p *Painter) WithStyle(n string, fn func(*Painter)) {
	p.SetStyle(p.Theme.Style(n))
	fn(p)
	p.RestoreStyle()
}

func (p *Painter) WithMask(r image.Rectangle) *Painter {
	p.mask = r
	return p
}

func (p *Painter) mapLocalToWorld(point image.Point) image.Point {
	var offset image.Point
	for _, s := range p.transforms {
		offset = offset.Add(s)
	}
	return point.Add(offset)
}
