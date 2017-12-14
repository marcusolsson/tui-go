package tui

import (
	"bytes"
	"image"
)

type testCell struct {
	Rune  rune
	Style Style
}

// A TestSurface implements the Surface interface with local buffers,
// and provides accessors to check the output of a draw operation on the Surface.
type TestSurface struct {
	cells   map[image.Point]testCell
	cursor  image.Point
	size    image.Point
	emptyCh rune
}

func NewTestSurface(w, h int) *TestSurface {
	return &TestSurface{
		cells:   make(map[image.Point]testCell),
		size:    image.Point{w, h},
		emptyCh: '.',
	}
}

func (s *TestSurface) SetCell(x, y int, ch rune, style Style) {
	s.cells[image.Point{x, y}] = testCell{
		Rune:  ch,
		Style: style,
	}
}

func (s *TestSurface) SetCursor(x, y int) {
	s.cursor = image.Point{x, y}
}

func (s *TestSurface) HideCursor() {
	s.cursor = image.Point{}
}

func (s *TestSurface) Begin() {
	s.cells = make(map[image.Point]testCell)
}

func (s *TestSurface) End() {
	// NOP
}

func (s *TestSurface) Size() image.Point {
	return s.size
}

// String writes the TestSurface's characters as a string.
func (s *TestSurface) String() string {
	var buf bytes.Buffer
	buf.WriteRune('\n')
	for j := 0; j < s.size.Y; j++ {
		for i := 0; i < s.size.X; i++ {
			if cell, ok := s.cells[image.Point{i, j}]; ok {
				buf.WriteRune(cell.Rune)
				if w := runeWidth(cell.Rune); w > 1 {
					i += w - 1
				}
			} else {
				buf.WriteRune(s.emptyCh)
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

// FgColors renders the TestSurface's foreground colors, using the digit 0-7 for painted cells.
func (s *TestSurface) FgColors() string {
	var buf bytes.Buffer
	buf.WriteRune('\n')
	for j := 0; j < s.size.Y; j++ {
		for i := 0; i < s.size.X; i++ {
			if cell, ok := s.cells[image.Point{i, j}]; ok {
				color := cell.Style.Fg
				if cell.Style.Reverse {
					color = cell.Style.Bg
				}
				buf.WriteRune('0' + rune(color))
			} else {
				buf.WriteRune(s.emptyCh)
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

// BgColors renders the TestSurface's background colors, using the digit 0-7 for painted cells.
func (s *TestSurface) BgColors() string {
	var buf bytes.Buffer
	buf.WriteRune('\n')
	for j := 0; j < s.size.Y; j++ {
		for i := 0; i < s.size.X; i++ {
			if cell, ok := s.cells[image.Point{i, j}]; ok {
				color := cell.Style.Bg
				if cell.Style.Reverse {
					color = cell.Style.Fg
				}
				buf.WriteRune('0' + rune(color))
			} else {
				buf.WriteRune(s.emptyCh)
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

