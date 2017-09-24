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
			g := NewGrid(2, 2)
			g.SetBorder(true)
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
	{
		test: "Grid with column stretch",
		size: image.Point{24, 3},
		setup: func() *Grid {
			g := NewGrid(3, 1)
			g.SetBorder(true)

			g.SetColumnStretch(0, 1)
			g.SetColumnStretch(1, 2)
			g.SetColumnStretch(2, 1)

			return g
		},
		want: `
┌─────┬──────────┬─────┐
│.....│..........│.....│
└─────┴──────────┴─────┘
`,
	},
	{
		test: "Grid with one undefined column stretch",
		size: image.Point{19, 3},
		setup: func() *Grid {
			g := NewGrid(3, 1)
			g.SetBorder(true)

			// First column stretch defaults to 0
			//g.SetColumnStretch(0, 0)

			g.SetColumnStretch(1, 2)
			g.SetColumnStretch(2, 1)

			return g
		},
		want: `
┌┬──────────┬─────┐
││..........│.....│
└┴──────────┴─────┘
`,
	},
	{
		test: "Grid with mixed column stretch",
		size: image.Point{34, 3},
		setup: func() *Grid {
			g := NewGrid(3, 1)
			g.SetBorder(true)

			g.SetColumnStretch(0, 3)
			g.SetColumnStretch(1, 2)
			g.SetColumnStretch(2, 1)

			return g
		},
		want: `
┌───────────────┬──────────┬─────┐
│...............│..........│.....│
└───────────────┴──────────┴─────┘
`,
	},
	{
		test: "Grid with single zero stretch column",
		size: image.Point{34, 3},
		setup: func() *Grid {
			g := NewGrid(0, 0)
			g.SetBorder(true)

			g.AppendRow(
				NewLabel("foo"),
				NewLabel("bar"),
				NewLabel("test"),
			)

			g.SetColumnStretch(0, 1)
			g.SetColumnStretch(1, 2)
			g.SetColumnStretch(2, 0)

			return g
		},
		want: `
┌──────────┬───────────────────┬─┐
│foo.......│bar................│t│
└──────────┴───────────────────┴─┘
`,
	},
	{
		test: "Grid with multiple zero stretch columns",
		size: image.Point{34, 3},
		setup: func() *Grid {
			g := NewGrid(0, 0)
			g.SetBorder(true)

			g.AppendRow(
				NewLabel("foo"),
				NewLabel("bar"),
				NewLabel("baz"),
				NewLabel("test"),
			)

			g.SetColumnStretch(0, 0)
			g.SetColumnStretch(1, 1)
			g.SetColumnStretch(2, 2)
			g.SetColumnStretch(3, 0)

			return g
		},
		want: `
┌─┬─────────┬──────────────────┬─┐
│f│bar......│baz...............│t│
└─┴─────────┴──────────────────┴─┘
`,
	},
}

func TestGrid_Draw(t *testing.T) {
	for _, tt := range drawGridTests {
		t.Run(tt.test, func(t *testing.T) {
			surface := newTestSurface(tt.size.X, tt.size.Y)
			painter := NewPainter(surface, NewTheme())
			painter.Repaint(tt.setup())

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}
