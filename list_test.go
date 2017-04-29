package tui

import (
	"image"
	"testing"

	"github.com/kr/pretty"
)

var listSizeTests = []struct {
	test        string
	setup       func() *List
	minSizeHint image.Point
	sizeHint    image.Point
	size        image.Point
}{
	{
		test: "Empty default",
		setup: func() *List {
			return NewList()
		},
		minSizeHint: image.Point{1, 1},
		sizeHint:    image.Point{0, 0},
		size:        image.Point{100, 100},
	},
	{
		test: "Empty with rows",
		setup: func() *List {
			l := NewList()
			l.AddItems("foo", "bar", "test")
			return l
		},
		minSizeHint: image.Point{1, 1},
		sizeHint:    image.Point{4, 3},
		size:        image.Point{100, 100},
	},
	{
		test: "Wide items",
		setup: func() *List {
			l := NewList()
			l.AddItems("あäa")
			return l
		},
		minSizeHint: image.Point{1, 1},
		sizeHint:    image.Point{4, 1},
		size:        image.Point{100, 100},
	},
}

func TestList_Size(t *testing.T) {
	for _, tt := range listSizeTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			t.Parallel()

			l := tt.setup()
			l.Resize(image.Point{100, 100})

			if got := l.MinSizeHint(); got != tt.minSizeHint {
				t.Errorf("l.MinSizeHint() = %s; want = %s", got, tt.minSizeHint)
			}
			if got := l.SizeHint(); got != tt.sizeHint {
				t.Errorf("l.SizeHint() = %s; want = %s", got, tt.sizeHint)
			}
			if got := l.Size(); got != tt.size {
				t.Errorf("l.Size() = %s; want = %s", got, tt.size)
			}
		})
	}
}

func TestList_Draw(t *testing.T) {
	surface := newTestSurface(10, 5)
	painter := NewPainter(surface, NewPalette())

	l := NewList()
	l.AddItems("foo", "bar")
	l.Resize(surface.size)
	l.Draw(painter)

	want := `
foo       
bar       
..........
..........
..........
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
	}
}
