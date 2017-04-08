package tui

import (
	"image"
	"testing"

	"github.com/kr/pretty"
)

var listSizeTests = []struct {
	test     string
	setup    func() *List
	size     image.Point
	sizeHint image.Point
}{
	{
		test: "Empty default",
		setup: func() *List {
			return NewList()
		},
		size:     image.Point{0, 5},
		sizeHint: image.Point{0, 5},
	},
	{
		test: "Empty with rows",
		setup: func() *List {
			l := NewList()
			l.SetRows(3)
			return l
		},
		size:     image.Point{0, 3},
		sizeHint: image.Point{0, 3},
	},
	{
		test: "Wide items",
		setup: func() *List {
			l := NewList()
			l.SetRows(3)
			l.AddItems("あäa")
			return l
		},
		size:     image.Point{4, 3},
		sizeHint: image.Point{4, 3},
	},
}

func TestList_Size(t *testing.T) {
	for _, tt := range listSizeTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			t.Parallel()

			l := tt.setup()
			l.Resize(image.Point{100, 00})

			if got := l.Size(); got != tt.size {
				t.Errorf("l.Size() = %s; want = %s", got, tt.size)
			}
			if got := l.SizeHint(); got != tt.sizeHint {
				t.Errorf("l.SizeHint() = %s; want = %s", got, tt.sizeHint)
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
foo.......
bar.......
..........
..........
..........
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
	}
}
