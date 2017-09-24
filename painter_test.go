package tui

import (
	"bytes"
	"image"
	"testing"
)

func TestMask_Full(t *testing.T) {
	surface := newTestSurface(10, 10)

	p := NewPainter(surface, NewTheme())
	p.WithMask(image.Rect(0, 0, 10, 10), func(p *Painter) {
		p.WithMask(image.Rect(0, 0, 10, 10), func(p *Painter) {
			sz := p.surface.Size()
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, '█')
				}
			}
		})
	})

	want := `
██████████
██████████
██████████
██████████
██████████
██████████
██████████
██████████
██████████
██████████
`
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestMask_Inset(t *testing.T) {
	surface := newTestSurface(10, 10)

	p := NewPainter(surface, NewTheme())
	p.WithMask(image.Rect(0, 0, 10, 10), func(p *Painter) {
		p.WithMask(image.Rect(1, 1, 9, 9), func(p *Painter) {
			sz := p.surface.Size()
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, '█')
				}
			}
		})
	})

	want := `
..........
.████████.
.████████.
.████████.
.████████.
.████████.
.████████.
.████████.
.████████.
..........
`
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestMask_FirstCell(t *testing.T) {
	surface := newTestSurface(10, 10)

	p := NewPainter(surface, NewTheme())
	p.WithMask(image.Rect(0, 0, 10, 10), func(p *Painter) {
		p.WithMask(image.Rect(0, 0, 1, 1), func(p *Painter) {
			sz := p.surface.Size()
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, '█')
				}
			}
		})
	})

	want := `
█.........
..........
..........
..........
..........
..........
..........
..........
..........
..........
`
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestMask_LastCell(t *testing.T) {
	surface := newTestSurface(10, 10)

	p := NewPainter(surface, NewTheme())
	p.WithMask(image.Rect(0, 0, 10, 10), func(p *Painter) {
		p.WithMask(image.Rect(9, 9, 10, 10), func(p *Painter) {
			sz := p.surface.Size()
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, '█')
				}
			}
		})
	})

	want := `
..........
..........
..........
..........
..........
..........
..........
..........
..........
.........█
`
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestMask_MaskWithinEmptyMaskIsHidden(t *testing.T) {
	surface := newTestSurface(10, 10)

	p := NewPainter(surface, NewTheme())
	p.WithMask(image.Rect(0, 0, 0, 0), func(p *Painter) {
		p.WithMask(image.Rect(1, 1, 9, 9), func(p *Painter) {
			sz := p.surface.Size()
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, '█')
				}
			}
		})
	})

	want := `
..........
..........
..........
..........
..........
..........
..........
..........
..........
..........
`
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

type testCell struct {
	Rune  rune
	Style Style
}

type testSurface struct {
	cells   map[image.Point]testCell
	cursor  image.Point
	size    image.Point
	emptyCh rune
}

func newTestSurface(w, h int) *testSurface {
	return &testSurface{
		cells:   make(map[image.Point]testCell),
		size:    image.Point{w, h},
		emptyCh: '.',
	}
}

func (s *testSurface) SetCell(x, y int, ch rune, style Style) {
	s.cells[image.Point{x, y}] = testCell{
		Rune:  ch,
		Style: style,
	}
}

func (s *testSurface) SetCursor(x, y int) {
	s.cursor = image.Point{x, y}
}

func (s *testSurface) HideCursor() {
	s.cursor = image.Point{}
}

func (s *testSurface) Begin() {
	s.cells = make(map[image.Point]testCell)
}

func (s *testSurface) End() {
	// NOP
}

func (s *testSurface) Size() image.Point {
	return s.size
}

func (s *testSurface) String() string {
	var buf bytes.Buffer
	buf.WriteRune('\n')
	for j := 0; j < s.size.Y; j++ {
		for i := 0; i < s.size.X; i++ {
			if cell, ok := s.cells[image.Point{i, j}]; ok {
				buf.WriteRune(cell.Rune)
			} else {
				buf.WriteRune(s.emptyCh)
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}
