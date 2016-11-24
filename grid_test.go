package tui

import (
	"image"
	"testing"

	"github.com/kr/pretty"
)

var gridSizeTests = []struct {
	test     string
	setup    func() *Grid
	size     image.Point
	sizeHint image.Point
}{
	{
		test: "Empty grid",
		setup: func() *Grid {
			g := NewGrid(0, 0)
			g.SetBorder(true)
			return g
		},
		size:     image.Point{0, 0},
		sizeHint: image.Point{0, 0},
	},
	{
		test: "1x1",
		setup: func() *Grid {
			g := NewGrid(1, 1)
			g.SetBorder(true)
			return g
		},
		size:     image.Point{2, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "2x2",
		setup: func() *Grid {
			g := NewGrid(2, 2)
			g.SetBorder(true)
			return g
		},
		size:     image.Point{3, 3},
		sizeHint: image.Point{3, 3},
	},
	{
		test: "3x1",
		setup: func() *Grid {
			g := NewGrid(3, 1)
			g.SetBorder(true)
			return g
		},
		size:     image.Point{4, 2},
		sizeHint: image.Point{4, 2},
	},
	{
		test: "1x3",
		setup: func() *Grid {
			g := NewGrid(1, 3)
			g.SetBorder(true)
			return g
		},
		size:     image.Point{2, 4},
		sizeHint: image.Point{2, 4},
	},
	{
		test: "1x1 with label",
		setup: func() *Grid {
			g := NewGrid(1, 1)
			g.SetCell(image.Point{0, 0}, NewLabel("test"))
			g.SetBorder(true)
			return g
		},
		size:     image.Point{6, 3},
		sizeHint: image.Point{6, 3},
	},
	{
		test: "2x2 with labels",
		setup: func() *Grid {
			g := NewGrid(2, 2)
			g.SetCell(image.Point{0, 0}, NewLabel("test"))
			g.SetCell(image.Point{1, 0}, NewLabel("test"))
			g.SetCell(image.Point{0, 1}, NewLabel("test"))
			g.SetCell(image.Point{1, 0}, NewLabel("test"))
			g.SetBorder(true)
			return g
		},
		size:     image.Point{11, 5},
		sizeHint: image.Point{11, 5},
	},
	{
		test: "Stretch 2x2 with labels",
		setup: func() *Grid {
			g := NewGrid(2, 2)
			g.SetCell(image.Point{0, 0}, NewLabel("test"))
			g.SetCell(image.Point{1, 0}, NewLabel("test"))
			g.SetCell(image.Point{0, 1}, NewLabel("test"))
			g.SetCell(image.Point{1, 0}, NewLabel("test"))
			g.SetBorder(true)
			g.SetSizePolicy(Expanding, Minimum)
			return g
		},
		size:     image.Point{100, 5},
		sizeHint: image.Point{11, 5},
	},
}

func TestGrid_Size(t *testing.T) {
	for _, tt := range gridSizeTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			t.Parallel()

			b := tt.setup()
			b.Resize(image.Point{100, 100})

			if got := b.Size(); got != tt.size {
				t.Errorf("b.Size() = %s; want = %s", got, tt.size)
			}
			if got := b.SizeHint(); got != tt.sizeHint {
				t.Errorf("b.SizeHint() = %s; want = %s", got, tt.sizeHint)
			}
		})
	}
}

func TestGrid_NestedSize(t *testing.T) {
	g := NewGrid(2, 2)
	g.SetCell(image.Point{0, 0}, NewLabel("test"))
	g.SetCell(image.Point{1, 0}, NewLabel("test"))
	g.SetCell(image.Point{0, 1}, NewLabel("test"))
	g.SetCell(image.Point{1, 1}, NewLabel("test"))
	g.SetBorder(true)
	g.SetSizePolicy(Expanding, Expanding)

	b := NewVBox(g)
	b.SetBorder(true)
	b.SetSizePolicy(Expanding, Minimum)

	b.Resize(image.Point{100, 100})

	wantSize := image.Point{98, 5}
	if got := g.Size(); got != wantSize {
		t.Errorf("g.Size() = %s; want = %s", got, wantSize)
	}
	wantSizeHint := image.Point{11, 5}
	if got := g.SizeHint(); got != wantSizeHint {
		t.Errorf("g.SizeHint() = %s; want = %s", got, wantSizeHint)
	}
}

func TestGrid_Draw(t *testing.T) {
	surface := newTestSurface(15, 5)
	painter := NewPainter(surface, NewPalette())

	g := NewGrid(0, 0)
	g.AppendRow(NewLabel("testing"), NewLabel("test"))
	g.AppendRow(NewLabel("foo"), NewLabel("bar"))

	g.Resize(surface.size)
	g.Draw(painter)

	want := `testingtest....
foo....bar.....
...............
...............
...............
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
	}
}
