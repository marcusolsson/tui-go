package tui_test

import (
	"image"
	"testing"
	"github.com/marcusolsson/tui-go"
)

var drawGridTests = []struct {
	test  string
	size  image.Point
	setup func() *tui.Grid
	want  string
}{
	{
		test: "Empty grid with border",
		size: image.Point{15, 5},
		setup: func() *tui.Grid {
			g := tui.NewGrid(0, 0)
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
		test: "tui.Grid with empty labels",
		size: image.Point{15, 5},
		setup: func() *tui.Grid {
			g := tui.NewGrid(2, 2)
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
		test: "tui.Grid with short labels",
		size: image.Point{19, 9},
		setup: func() *tui.Grid {
			g := tui.NewGrid(0, 0)
			g.SetBorder(true)
			l := tui.NewLabel("testing")
			l.SetSizePolicy(tui.Minimum, tui.Preferred)
			g.AppendRow(l, tui.NewLabel("test"))
			g.AppendRow(tui.NewLabel("foo"), tui.NewLabel("bar"))
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
		test: "tui.Grid with word wrap",
		size: image.Point{19, 5},
		setup: func() *tui.Grid {
			l := tui.NewLabel("this will wrap")
			l.SetWordWrap(true)
			l.SetSizePolicy(tui.Expanding, tui.Preferred)

			g := tui.NewGrid(0, 0)
			g.SetBorder(true)
			g.AppendRow(l, tui.NewLabel("test"))
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
		test: "tui.Grid with column stretch",
		size: image.Point{24, 3},
		setup: func() *tui.Grid {
			g := tui.NewGrid(3, 1)
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
		test: "tui.Grid with one undefined column stretch",
		size: image.Point{19, 3},
		setup: func() *tui.Grid {
			g := tui.NewGrid(3, 1)
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
		test: "tui.Grid with mixed column stretch",
		size: image.Point{34, 3},
		setup: func() *tui.Grid {
			g := tui.NewGrid(3, 1)
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
		test: "tui.Grid with single zero stretch column",
		size: image.Point{34, 3},
		setup: func() *tui.Grid {
			g := tui.NewGrid(0, 0)
			g.SetBorder(true)

			g.AppendRow(
				tui.NewLabel("foo"),
				tui.NewLabel("bar"),
				tui.NewLabel("test"),
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
		test: "tui.Grid with multiple zero stretch columns",
		size: image.Point{34, 3},
		setup: func() *tui.Grid {
			g := tui.NewGrid(0, 0)
			g.SetBorder(true)

			g.AppendRow(
				tui.NewLabel("foo"),
				tui.NewLabel("bar"),
				tui.NewLabel("baz"),
				tui.NewLabel("test"),
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
			painter := tui.NewPainter(surface, tui.NewTheme())
			painter.Repaint(tt.setup())

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}
