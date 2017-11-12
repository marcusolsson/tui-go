package tui_test

import (
	"image"
	"testing"
	"github.com/marcusolsson/tui-go"
)

var drawScrollAreaTests = []struct {
	test  string
	size  image.Point
	setup func() *tui.ScrollArea
	want  string
}{
	{
		test: "Empty scroll area",
		size: image.Point{10, 3},
		setup: func() *tui.ScrollArea {
			b := tui.NewVBox(
				tui.NewLabel("foo"),
				tui.NewLabel("bar"),
				tui.NewLabel("test"),
			)
			a := tui.NewScrollArea(b)
			return a
		},
		want: `
foo.......
bar.......
test......
`,
	},
	{
		test: "Vertical scroll top",
		size: image.Point{10, 2},
		setup: func() *tui.ScrollArea {
			b := tui.NewVBox(
				tui.NewLabel("foo"),
				tui.NewLabel("bar"),
				tui.NewLabel("test"),
			)
			a := tui.NewScrollArea(b)
			return a
		},
		want: `
foo.......
bar.......
`,
	},
	{
		test: "Vertical scroll bottom",
		size: image.Point{10, 2},
		setup: func() *tui.ScrollArea {
			b := tui.NewVBox(
				tui.NewLabel("foo"),
				tui.NewLabel("bar"),
				tui.NewLabel("test"),
			)
			a := tui.NewScrollArea(b)
			a.Scroll(0, 1)
			return a
		},
		want: `
bar.......
test......
`,
	},
	{
		test: "Horizontal scroll left",
		size: image.Point{10, 1},
		setup: func() *tui.ScrollArea {
			b := tui.NewVBox(
				tui.NewLabel("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			)
			a := tui.NewScrollArea(b)
			return a
		},
		want: `
Lorem ipsu
`,
	},
	{
		test: "Horizontal scroll right",
		size: image.Point{10, 1},
		setup: func() *tui.ScrollArea {
			b := tui.NewVBox(
				tui.NewLabel("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			)
			a := tui.NewScrollArea(b)
			a.Scroll(46, 0)
			return a
		},
		want: `
cing elit.
`,
	},
}

func TestScrollArea_Draw(t *testing.T) {
	for _, tt := range drawScrollAreaTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			var surface *testSurface
			if tt.size.X == 0 && tt.size.Y == 0 {
				surface = newTestSurface(10, 5)
			} else {
				surface = newTestSurface(tt.size.X, tt.size.Y)
			}
			painter := tui.NewPainter(surface, tui.NewTheme())

			a := tt.setup()

			a.Resize(surface.size)
			a.Draw(painter)

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}

var drawNestedScrollAreaTests = []struct {
	test  string
	size  image.Point
	setup func() *tui.Box
	want  string
}{
	{
		test: "Nested vertical scroll",
		size: image.Point{11, 12},
		setup: func() *tui.Box {
			l := tui.NewList()
			l.AddItems("foo", "bar", "test")

			b1 := tui.NewVBox(l)
			b1.SetBorder(true)

			nested := tui.NewVBox(tui.NewLabel("foo"))
			nested.SetBorder(true)

			nested2 := tui.NewVBox(nested)
			nested2.SetBorder(true)

			s := tui.NewScrollArea(nested2)
			s.Scroll(0, 4)

			b2 := tui.NewVBox(s)
			b2.SetBorder(true)

			b3 := tui.NewVBox(b1, b2)
			b3.SetBorder(true)

			return b3
		},
		want: `
┌─────────┐
│┌───────┐│
││foo    ││
││bar    ││
││test   ││
│└───────┘│
│┌───────┐│
││└─────┘││
││.......││
││.......││
│└───────┘│
└─────────┘
`,
	},
	{
		test: "Nested horizontal scroll",
		size: image.Point{20, 9},
		setup: func() *tui.Box {
			nested := tui.NewVBox(tui.NewLabel("foo"))
			nested.SetBorder(true)

			nested2 := tui.NewVBox(nested)
			nested2.SetBorder(true)

			s := tui.NewScrollArea(nested2)
			s.Scroll(-5, 0)

			b1 := tui.NewVBox(s)
			b1.SetBorder(true)

			b2 := tui.NewVBox(tui.NewLabel("1234567"))
			b2.SetBorder(true)

			b3 := tui.NewHBox(b1, b2)
			b3.SetBorder(true)

			return b3
		},
		want: `
┌──────────────────┐
│┌───────┐┌───────┐│
││.....┌─││1234567││
││.....│┌││.......││
││.....││││.......││
││.....│└││.......││
││.....└─││.......││
│└───────┘└───────┘│
└──────────────────┘
`,
	},
}

func TestNestedScrollArea_Draw(t *testing.T) {
	for _, tt := range drawNestedScrollAreaTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			var surface *testSurface
			if tt.size.X == 0 && tt.size.Y == 0 {
				surface = newTestSurface(10, 5)
			} else {
				surface = newTestSurface(tt.size.X, tt.size.Y)
			}

			painter := tui.NewPainter(surface, tui.NewTheme())
			painter.Repaint(tt.setup())

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}
