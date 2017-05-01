package tui

import (
	"image"
	"testing"
)

var drawVBoxTests = []struct {
	test  string
	setup func() *Box
	want  string
}{
	{
		test: "Empty box",
		setup: func() *Box {
			b := NewVBox()
			b.SetBorder(true)
			return b
		},
		want: `
┌────────┐
│........│
│........│
│........│
└────────┘
`,
	},
	{
		test: "Box containing one widget",
		setup: func() *Box {
			b := NewVBox(
				NewLabel("test"),
			)
			b.SetBorder(true)
			return b
		},
		want: `
┌────────┐
│test....│
│........│
│........│
└────────┘
`,
	},
	{
		test: "Box containing multiple widget",
		setup: func() *Box {
			b := NewVBox(
				NewLabel("test"),
				NewLabel("foo"),
			)
			b.SetBorder(true)
			return b
		},
		want: `
┌────────┐
│test....│
│........│
│foo.....│
└────────┘
`,
	},
}

func TestVBox_Draw(t *testing.T) {
	for _, tt := range drawVBoxTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			surface := newTestSurface(10, 5)
			painter := NewPainter(surface, NewTheme())

			b := tt.setup()

			b.Resize(surface.size)
			b.Draw(painter)

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}

var drawHBoxTests = []struct {
	test  string
	size  image.Point
	setup func() *Box
	want  string
}{
	{
		test: "Empty box",
		setup: func() *Box {
			b := NewHBox()
			b.SetBorder(true)
			return b
		},
		want: `
┌────────┐
│........│
│........│
│........│
└────────┘
`,
	},
	{
		test: "Box containing one widget",
		setup: func() *Box {
			b := NewHBox(
				NewLabel("test"),
			)
			b.SetBorder(true)
			return b
		},
		want: `
┌────────┐
│test....│
│........│
│........│
└────────┘
`,
	},
	{
		test: "Box containing multiple widgets",
		setup: func() *Box {
			b := NewHBox(
				NewLabel("test"),
				NewLabel("foo"),
			)
			b.SetBorder(true)
			return b
		},
		want: `
┌────────┐
│testfoo.│
│........│
│........│
└────────┘
`,
	},
	{
		test: "Nested boxes expands equally",
		size: image.Point{32, 5},
		setup: func() *Box {
			nested := NewVBox(
				NewLabel("test"),
			)
			nested.SetBorder(true)

			b := NewHBox(
				NewSpacer(),
				nested,
				NewSpacer(),
			)
			b.SetBorder(true)
			return b
		},
		want: `
┌──────────────────────────────┐
│............┌────┐............│
│............│test│............│
│............└────┘............│
└──────────────────────────────┘
`,
	},
	{
		test: "Two columns",
		size: image.Point{32, 10},
		setup: func() *Box {
			first := NewVBox(
				NewLabel("test"),
			)
			first.SetBorder(true)

			second := NewVBox(
				NewLabel("test"),
			)
			second.SetBorder(true)

			b := NewHBox(
				first,
				second,
			)
			b.SetBorder(true)
			return b
		},
		want: `
┌──────────────────────────────┐
│┌─────────────┐┌─────────────┐│
││test.........││test.........││
││.............││.............││
││.............││.............││
││.............││.............││
││.............││.............││
││.............││.............││
│└─────────────┘└─────────────┘│
└──────────────────────────────┘
`,
	},
	{
		test: "Two rows with two columns",
		size: image.Point{32, 22},
		setup: func() *Box {
			col0 := NewVBox(NewLabel("test"))
			col0.SetBorder(true)
			col1 := NewVBox(NewLabel("test"))
			col1.SetBorder(true)

			row0 := NewHBox(col0, col1)
			row0.SetBorder(true)

			col2 := NewVBox(NewLabel("test"))
			col2.SetBorder(true)
			col3 := NewVBox(NewLabel("test"))
			col3.SetBorder(true)

			row1 := NewHBox(col2, col3)
			row1.SetBorder(true)

			b := NewVBox(row0, row1)
			b.SetBorder(true)
			return b
		},
		want: `
┌──────────────────────────────┐
│┌────────────────────────────┐│
││┌────────────┐┌────────────┐││
│││test........││test........│││
│││............││............│││
│││............││............│││
│││............││............│││
│││............││............│││
│││............││............│││
││└────────────┘└────────────┘││
│└────────────────────────────┘│
│┌────────────────────────────┐│
││┌────────────┐┌────────────┐││
│││test........││test........│││
│││............││............│││
│││............││............│││
│││............││............│││
│││............││............│││
│││............││............│││
││└────────────┘└────────────┘││
│└────────────────────────────┘│
└──────────────────────────────┘
`,
	},
	{
		test: "Combined",
		size: image.Point{32, 24},
		setup: func() *Box {
			edit0 := NewLabel("")
			edit0.SetText("test\ntesting\nfoo\nbar")
			row0 := NewVBox(edit0)
			row0.SetSizePolicy(Preferred, Maximum)
			row0.SetBorder(true)

			edit1 := NewLabel("")
			edit1.SetText("test\ntesting\nfoo\nbar")
			row1 := NewVBox(edit1)
			row1.SetSizePolicy(Preferred, Preferred)
			row1.SetBorder(true)

			edit2 := NewLabel("")
			edit2.SetText("foo")
			row2 := NewVBox(edit2)
			row2.SetSizePolicy(Preferred, Preferred)
			row2.SetBorder(true)

			b := NewVBox(row0, row1, row2)
			b.SetBorder(true)
			return b
		},
		want: `
┌──────────────────────────────┐
│┌────────────────────────────┐│
││test........................││
││testing.....................││
││foo.........................││
││bar.........................││
│└────────────────────────────┘│
│┌────────────────────────────┐│
││test........................││
││testing.....................││
││foo.........................││
││bar.........................││
││............................││
││............................││
│└────────────────────────────┘│
│┌────────────────────────────┐│
││foo.........................││
││............................││
││............................││
││............................││
││............................││
││............................││
│└────────────────────────────┘│
└──────────────────────────────┘
`,
	},
}

func TestHBox_Draw(t *testing.T) {
	for _, tt := range drawHBoxTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			var surface *testSurface
			if tt.size.X == 0 && tt.size.Y == 0 {
				surface = newTestSurface(10, 5)
			} else {
				surface = newTestSurface(tt.size.X, tt.size.Y)
			}
			painter := NewPainter(surface, NewTheme())

			b := tt.setup()

			b.Resize(surface.size)
			b.Draw(painter)

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}

type dummyWidget struct {
	minSizeHint image.Point
	sizeHint    image.Point
	size        image.Point
	sizePolicyX SizePolicy
	sizePolicyY SizePolicy
}

func (w *dummyWidget) MinSizeHint() image.Point {
	return w.minSizeHint
}
func (w *dummyWidget) Size() image.Point {
	return w.size
}
func (w *dummyWidget) SizeHint() image.Point {
	return w.sizeHint
}
func (w *dummyWidget) SizePolicy() (SizePolicy, SizePolicy) {
	return w.sizePolicyX, w.sizePolicyY
}
func (w *dummyWidget) Resize(size image.Point) {
	w.size = size
}

// Not used
func (w *dummyWidget) Draw(p *Painter)  {}
func (w *dummyWidget) OnEvent(ev Event) {}
