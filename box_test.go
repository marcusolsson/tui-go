package tui

import (
	"image"
	"testing"
)

var drawBoxTests = []struct {
	test  string
	size  image.Point
	setup func() *Box
	want  string
}{
	{
		test: "Empty horizontal box",
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
		test: "Horizontal box containing one widget",
		setup: func() *Box {
			b := NewHBox(NewLabel("test"))
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
		test: "Horizontal box containing multiple widgets",
		setup: func() *Box {
			b := NewHBox(NewLabel("test"), NewLabel("foo"))
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
		test: "Empty vertical box",
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
		test: "Vertical box containing one widget",
		setup: func() *Box {
			b := NewVBox(NewLabel("test"))
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
		test: "Vertical box containing multiple widgets",
		setup: func() *Box {
			b := NewVBox(NewLabel("test"), NewLabel("foo"))
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
	{
		test: "Horizontally centered box",
		size: image.Point{32, 5},
		setup: func() *Box {
			nested := NewVBox(NewLabel("test"))
			nested.SetBorder(true)

			b := NewHBox(NewSpacer(), nested, NewSpacer())
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
			first := NewVBox(NewLabel("test"))
			first.SetBorder(true)

			second := NewVBox(NewLabel("test"))
			second.SetBorder(true)

			third := NewHBox(first, second)
			third.SetBorder(true)

			return third
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
		test: "Maximum/Preferred/Preferred",
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

func TestBox_Draw(t *testing.T) {
	for _, tt := range drawBoxTests {
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
