package tui

import (
	"bytes"
	"fmt"
	"image"
	"strconv"
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

// NewTestSurface returns a new TestSurface.
func NewTestSurface(w, h int) *TestSurface {
	return &TestSurface{
		cells:   make(map[image.Point]testCell),
		size:    image.Point{w, h},
		emptyCh: '.',
	}
}

// SetCell sets the contents of the addressed cell.
func (s *TestSurface) SetCell(x, y int, ch rune, style Style) {
	s.cells[image.Point{x, y}] = testCell{
		Rune:  ch,
		Style: style,
	}
}

// SetCursor moves the Surface's cursor to the specified position.
func (s *TestSurface) SetCursor(x, y int) {
	s.cursor = image.Point{x, y}
}

// HideCursor removes the cursor from the display.
func (s *TestSurface) HideCursor() {
	s.cursor = image.Point{}
}

// Begin resets the state of the TestSurface, clearing all cells.
// It must be called before drawing the Surface.
func (s *TestSurface) Begin() {
	s.cells = make(map[image.Point]testCell)
}

// End indicates the surface has been painted on, and can be rendered.
// It's a no-op for TestSurface.
func (s *TestSurface) End() {
	// NOP
}

// Size returns the dimensions of the surface.
func (s *TestSurface) Size() image.Point {
	return s.size
}

// String returns the characters written to the TestSurface.
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

// FgColors renders the TestSurface's foreground colors, using the digits 0-7 for painted cells, and the empty character for unpainted cells.
func (s *TestSurface) FgColors() string {
	var buf bytes.Buffer
	buf.WriteRune('\n')
	for j := 0; j < s.size.Y; j++ {
		for i := 0; i < s.size.X; i++ {
			if cell, ok := s.cells[image.Point{i, j}]; ok {
				color := cell.Style.Fg
				buf.WriteRune('0' + rune(color))
			} else {
				buf.WriteRune(s.emptyCh)
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

// BgColors renders the TestSurface's background colors, using the digits 0-7 for painted cells, and the empty character for unpainted cells.
func (s *TestSurface) BgColors() string {
	var buf bytes.Buffer
	buf.WriteRune('\n')
	for j := 0; j < s.size.Y; j++ {
		for i := 0; i < s.size.X; i++ {
			if cell, ok := s.cells[image.Point{i, j}]; ok {
				color := cell.Style.Bg
				buf.WriteRune('0' + rune(color))
			} else {
				buf.WriteRune(s.emptyCh)
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

// Decorations renders the TestSurface's decorations (Reverse, Bold, Underline) using a bitmask:
//	Reverse: 1
//	Bold: 2
//	Underline: 4
func (s *TestSurface) Decorations() string {
	var buf bytes.Buffer
	buf.WriteRune('\n')
	for j := 0; j < s.size.Y; j++ {
		for i := 0; i < s.size.X; i++ {
			if cell, ok := s.cells[image.Point{i, j}]; ok {
				mask := int64(0)
				if cell.Style.Reverse == DecorationOn {
					mask |= 1
				}
				if cell.Style.Bold == DecorationOn {
					mask |= 2
				}
				if cell.Style.Underline == DecorationOn {
					mask |= 4
				}
				buf.WriteString(strconv.FormatInt(mask, 16))
			} else {
				buf.WriteRune(s.emptyCh)
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

func surfaceEquals(surface *TestSurface, want string) string {
	if surface.String() != want {
		return fmt.Sprintf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
	return ""
}
