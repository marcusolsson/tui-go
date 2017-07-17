package tui

import (
	"bytes"
	"image"
)

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
