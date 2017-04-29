package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

// Surface defines a surface that can be painted on.
type Surface interface {
	SetCell(x, y int, ch rune, fg, bg Color)
	SetCursor(x, y int)
	Begin()
	End()
	Size() image.Point
}

type termboxSurface struct{}

// NewTermboxSurface returns the default paint surface.
func NewTermboxSurface() Surface {
	return termboxSurface{}
}

func (s termboxSurface) SetCell(x, y int, ch rune, fg, bg Color) {
	termbox.SetCell(x, y, ch, termbox.Attribute(fg), termbox.Attribute(bg))
}

func (s termboxSurface) SetCursor(x, y int) {
	termbox.SetCursor(x, y)
}

func (s termboxSurface) Begin() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func (s termboxSurface) End() {
	termbox.Flush()
}

func (s termboxSurface) Size() image.Point {
	w, h := termbox.Size()
	return image.Point{w, h}
}

// Painter provides operations to paint on a surface.
type Painter struct {
	// Surface to paint on.
	surface Surface

	palette *Palette

	// Current brush.
	fg, bg Color

	// Transform stack
	transforms []image.Point

	mask image.Rectangle
}

// NewPainter returns a new instance of Painter.
func NewPainter(s Surface, p *Palette) *Painter {
	return &Painter{
		surface: s,
		palette: p,
		fg:      p.Item("normal").Fg,
		bg:      p.Item("normal").Bg,
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

// DrawText paints a string starting at the given coordinate.
func (p *Painter) DrawText(x, y int, text string) {
	for _, r := range text {
		p.DrawRune(x, y, r)
		x += runeWidth(r)
	}
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
	p.surface.SetCell(wp.X, wp.Y, r, p.fg, p.bg)
}

func (p *Painter) DrawHorizontalLine(x1, x2, y int) {
	for x := x1; x < x2; x++ {
		wp := p.mapLocalToWorld(image.Point{x, y})
		p.surface.SetCell(wp.X, wp.Y, '─', p.fg, p.bg)
	}
}

func (p *Painter) DrawVerticalLine(x, y1, y2 int) {
	for y := y1; y < y2; y++ {
		wp := p.mapLocalToWorld(image.Point{x, y})
		p.surface.SetCell(wp.X, wp.Y, '│', p.fg, p.bg)
	}
}

// DrawRect paints a rectangle.
func (p *Painter) DrawRect(x, y, w, h int) {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			wp := p.mapLocalToWorld(image.Point{i + x, j + y})
			switch {
			case i == 0 && j == 0:
				p.surface.SetCell(wp.X, wp.Y, '┌', p.fg, p.bg)
			case i == w-1 && j == 0:
				p.surface.SetCell(wp.X, wp.Y, '┐', p.fg, p.bg)
			case i == 0 && j == h-1:
				p.surface.SetCell(wp.X, wp.Y, '└', p.fg, p.bg)
			case i == w-1 && j == h-1:
				p.surface.SetCell(wp.X, wp.Y, '┘', p.fg, p.bg)
			case i == 0 || i == w-1:
				p.surface.SetCell(wp.X, wp.Y, '│', p.fg, p.bg)
			case j == 0 || j == h-1:
				p.surface.SetCell(wp.X, wp.Y, '─', p.fg, p.bg)
			}
		}
	}
}

func (p *Painter) FillRect(x, y, w, h int) {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			wp := p.mapLocalToWorld(image.Point{i + x, j + y})
			p.surface.SetCell(wp.X, wp.Y, ' ', p.fg, p.bg)
		}
	}
}

func (p *Painter) DrawCursor(x, y int) {
	wp := p.mapLocalToWorld(image.Point{x, y})
	p.surface.SetCursor(wp.X, wp.Y)
}

func (p *Painter) SetBrush(fg, bg Color) {
	p.fg = fg
	p.bg = bg
}

func (p *Painter) RestoreBrush() {
	p.SetBrush(p.palette.Item("normal").Fg, p.palette.Item("normal").Bg)
}

func (p *Painter) WithStyledBrush(n string, fn func(*Painter)) {
	p.SetBrush(p.palette.Item(n).Fg, p.palette.Item(n).Bg)
	fn(p)
	p.RestoreBrush()
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
