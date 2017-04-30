package tui

import (
	"image"
	"testing"
)

var drawGridTests = []struct {
	test  string
	size  image.Point
	setup func() *Grid
	want  string
}{
	{
		test: "Empty grid with border",
		size: image.Point{15, 5},
		setup: func() *Grid {
			g := NewGrid(0, 0)
			g.SetBorder(true)
			return g
		},
		want: `
┌─────────────┐
│.............│
│.............│
│.............│
└─────────────┘
`,
	},
	{
		test: "Grid with empty labels",
		size: image.Point{15, 5},
		setup: func() *Grid {
			g := NewGrid(0, 0)
			g.SetBorder(true)
			g.AppendRow(NewLabel(""), NewLabel(""))
			g.AppendRow(NewLabel(""), NewLabel(""))
			return g
		},
		want: `
┌──────┬──────┐
│......│......│
├──────┼──────┤
│......│......│
└──────┴──────┘
`,
	},
	{
		test: "Grid with short labels",
		size: image.Point{19, 9},
		setup: func() *Grid {
			g := NewGrid(0, 0)
			g.SetBorder(true)
			l := NewLabel("testing")
			l.SetSizePolicy(Minimum, Preferred)
			g.AppendRow(l, NewLabel("test"))
			g.AppendRow(NewLabel("foo"), NewLabel("bar"))
			return g
		},
		want: `
┌────────┬────────┐
│testing.│test....│
│........│........│
│........│........│
├────────┼────────┤
│foo.....│bar.....│
│........│........│
│........│........│
└────────┴────────┘
`,
	},
	{
		test: "Grid with word wrap",
		size: image.Point{19, 5},
		setup: func() *Grid {
			l := NewLabel("this will wrap")
			l.SetWordWrap(true)
			l.SetSizePolicy(Expanding, Preferred)

			g := NewGrid(0, 0)
			g.SetBorder(true)
			g.AppendRow(l, NewLabel("test"))
			return g
		},
		want: `
┌────────┬────────┐
│this....│test....│
│will....│........│
│wrap....│........│
└────────┴────────┘
`,
	},
}

func TestGrid_Draw(t *testing.T) {
	for _, tt := range drawGridTests {
		surface := newTestSurface(tt.size.X, tt.size.Y)
		painter := NewPainter(surface, NewTheme())

		g := tt.setup()

		g.Resize(surface.size)
		g.Draw(painter)

		if surface.String() != tt.want {
			t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
		}
	}
}
