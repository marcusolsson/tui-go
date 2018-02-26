package tui

import (
	"fmt"
	"image"
	"testing"
)

var drawScrollAreaTests = []struct {
	test  string
	size  image.Point
	setup func() *ScrollArea
	want  string
}{
	{
		test: "Empty scroll area",
		size: image.Point{10, 3},
		setup: func() *ScrollArea {
			b := NewVBox(
				NewLabel("foo"),
				NewLabel("bar"),
				NewLabel("test"),
			)
			a := NewScrollArea(b)
			return a
		},
		want: `
foo ......
bar ......
test......
`,
	},
	{
		test: "Vertical scroll top",
		size: image.Point{10, 2},
		setup: func() *ScrollArea {
			b := NewVBox(
				NewLabel("foo"),
				NewLabel("bar"),
				NewLabel("test"),
			)
			a := NewScrollArea(b)
			return a
		},
		want: `
foo ......
bar ......
`,
	},
	{
		test: "Vertical scroll bottom",
		size: image.Point{10, 2},
		setup: func() *ScrollArea {
			b := NewVBox(
				NewLabel("foo"),
				NewLabel("bar"),
				NewLabel("test"),
			)
			a := NewScrollArea(b)
			a.Scroll(0, 1)
			return a
		},
		want: `
bar ......
test......
`,
	},
	{
		test: "Horizontal scroll left",
		size: image.Point{10, 1},
		setup: func() *ScrollArea {
			b := NewVBox(
				NewLabel("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			)
			a := NewScrollArea(b)
			return a
		},
		want: `
Lorem ipsu
`,
	},
	{
		test: "Horizontal scroll right",
		size: image.Point{10, 1},
		setup: func() *ScrollArea {
			b := NewVBox(
				NewLabel("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			)
			a := NewScrollArea(b)
			a.Scroll(46, 0)
			return a
		},
		want: `
cing elit.
`,
	},
	{
		test: "Scroll horizontally and vertically",
		size: image.Point{10, 3},
		setup: func() *ScrollArea {
			b := NewVBox()
			for i := 0; i < 10; i++ {
				b.Append(NewLabel(fmt.Sprintf("row %d", i)))
			}
			a := NewScrollArea(b)
			a.Scroll(1, 2)
			return a
		},
		want: `
ow 2......
ow 3......
ow 4......
`,
	},
}

func TestScrollArea_Draw(t *testing.T) {
	for _, tt := range drawScrollAreaTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			var surface *TestSurface
			if tt.size.X == 0 && tt.size.Y == 0 {
				surface = NewTestSurface(10, 5)
			} else {
				surface = NewTestSurface(tt.size.X, tt.size.Y)
			}
			painter := NewPainter(surface, NewTheme())

			a := tt.setup()

			a.Resize(surface.size)
			a.Draw(painter)

			if diff := surfaceEquals(surface, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestScrollArea_ScrollToBottom(t *testing.T) {
	surface := NewTestSurface(10, 3)
	painter := NewPainter(surface, NewTheme())

	b := NewVBox()
	for i := 0; i < 10; i++ {
		b.Append(NewLabel(fmt.Sprintf("row %d", i)))
	}
	a := NewScrollArea(b)

	a.Resize(surface.size)

	a.ScrollToBottom()

	a.Draw(painter)

	want := `
row 7.....
row 8.....
row 9.....
`

	if diff := surfaceEquals(surface, want); diff != "" {
		t.Error(diff)
	}
}

func TestScrollArea_ScrollToBottom_MaintainHScroll(t *testing.T) {
	surface := NewTestSurface(10, 3)
	painter := NewPainter(surface, NewTheme())

	b := NewVBox()
	for i := 0; i < 10; i++ {
		b.Append(NewLabel(fmt.Sprintf("row %d", i)))
	}
	a := NewScrollArea(b)

	a.Resize(surface.size)

	a.Scroll(2, 0)
	a.ScrollToBottom()

	a.Draw(painter)

	want := `
w 7.......
w 8.......
w 9.......
`

	if diff := surfaceEquals(surface, want); diff != "" {
		t.Error(diff)
	}
}

func TestScrollArea_ScrollToTop(t *testing.T) {
	surface := NewTestSurface(10, 3)
	painter := NewPainter(surface, NewTheme())

	b := NewVBox()
	for i := 0; i < 10; i++ {
		b.Append(NewLabel(fmt.Sprintf("row %d", i)))
	}
	a := NewScrollArea(b)

	a.Resize(surface.size)

	a.Scroll(0, 2)
	a.ScrollToTop()

	a.Draw(painter)

	want := `
row 0.....
row 1.....
row 2.....
`

	if diff := surfaceEquals(surface, want); diff != "" {
		t.Error(diff)
	}
}

func TestScrollArea_ScrollToTop_MaintainHScroll(t *testing.T) {
	surface := NewTestSurface(10, 3)
	painter := NewPainter(surface, NewTheme())

	b := NewVBox()
	for i := 0; i < 10; i++ {
		b.Append(NewLabel(fmt.Sprintf("row %d", i)))
	}
	a := NewScrollArea(b)

	a.Resize(surface.size)

	a.Scroll(2, 5)
	a.ScrollToTop()

	a.Draw(painter)

	want := `
w 0.......
w 1.......
w 2.......
`

	if diff := surfaceEquals(surface, want); diff != "" {
		t.Error(diff)
	}
}

var drawNestedScrollAreaTests = []struct {
	test  string
	size  image.Point
	setup func() *Box
	want  string
}{
	{
		test: "Nested vertical scroll",
		size: image.Point{11, 12},
		setup: func() *Box {
			l := NewList()
			l.AddItems("foo", "bar", "test")

			b1 := NewVBox(l)
			b1.SetBorder(true)

			nested := NewVBox(NewLabel("foo"))
			nested.SetBorder(true)

			nested2 := NewVBox(nested)
			nested2.SetBorder(true)

			s := NewScrollArea(nested2)
			s.Scroll(0, 4)

			b2 := NewVBox(s)
			b2.SetBorder(true)

			b3 := NewVBox(b1, b2)
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
││       ││
││       ││
│└───────┘│
└─────────┘
`,
	},
	{
		test: "Nested horizontal scroll",
		size: image.Point{20, 9},
		setup: func() *Box {
			nested := NewVBox(NewLabel("foo"))
			nested.SetBorder(true)

			nested2 := NewVBox(nested)
			nested2.SetBorder(true)

			s := NewScrollArea(nested2)
			s.Scroll(-5, 0)

			b1 := NewVBox(s)
			b1.SetBorder(true)

			b2 := NewVBox(NewLabel("1234567"))
			b2.SetBorder(true)

			b3 := NewHBox(b1, b2)
			b3.SetBorder(true)

			return b3
		},
		want: `
┌──────────────────┐
│┌───────┐┌───────┐│
││     ┌─││1234567││
││     │┌││       ││
││     ││││       ││
││     │└││       ││
││     └─││       ││
│└───────┘└───────┘│
└──────────────────┘
`,
	},
}

func TestNestedScrollArea_Draw(t *testing.T) {
	for _, tt := range drawNestedScrollAreaTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			var surface *TestSurface
			if tt.size.X == 0 && tt.size.Y == 0 {
				surface = NewTestSurface(10, 5)
			} else {
				surface = NewTestSurface(tt.size.X, tt.size.Y)
			}

			painter := NewPainter(surface, NewTheme())
			painter.Repaint(tt.setup())

			if diff := surfaceEquals(surface, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
