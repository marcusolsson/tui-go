package tui_test

import (
	"image"
	"testing"
	"github.com/marcusolsson/tui-go"
)

var drawBoxTests = []struct {
	test  string
	size  image.Point
	setup func() *tui.Box
	want  string
}{
	{
		test: "Empty horizontal box",
		setup: func() *tui.Box {
			b := tui.NewHBox()
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
		setup: func() *tui.Box {
			b := tui.NewHBox(tui.NewLabel("test"))
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
		setup: func() *tui.Box {
			b := tui.NewHBox(tui.NewLabel("test"), tui.NewLabel("foo"))
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
		setup: func() *tui.Box {
			b := tui.NewVBox()
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
		setup: func() *tui.Box {
			b := tui.NewVBox(tui.NewLabel("test"))
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
		setup: func() *tui.Box {
			b := tui.NewVBox(tui.NewLabel("test"), tui.NewLabel("foo"))
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
		setup: func() *tui.Box {
			nested := tui.NewVBox(tui.NewLabel("test"))
			nested.SetBorder(true)

			b := tui.NewHBox(tui.NewSpacer(), nested, tui.NewSpacer())
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
		setup: func() *tui.Box {
			first := tui.NewVBox(tui.NewLabel("test"))
			first.SetBorder(true)

			second := tui.NewVBox(tui.NewLabel("test"))
			second.SetBorder(true)

			third := tui.NewHBox(first, second)
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
		setup: func() *tui.Box {
			col0 := tui.NewVBox(tui.NewLabel("test"))
			col0.SetBorder(true)
			col1 := tui.NewVBox(tui.NewLabel("test"))
			col1.SetBorder(true)

			row0 := tui.NewHBox(col0, col1)
			row0.SetBorder(true)

			col2 := tui.NewVBox(tui.NewLabel("test"))
			col2.SetBorder(true)
			col3 := tui.NewVBox(tui.NewLabel("test"))
			col3.SetBorder(true)

			row1 := tui.NewHBox(col2, col3)
			row1.SetBorder(true)

			b := tui.NewVBox(row0, row1)
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
		setup: func() *tui.Box {
			edit0 := tui.NewLabel("")
			edit0.SetText("test\ntesting\nfoo\nbar")
			row0 := tui.NewVBox(edit0)
			row0.SetSizePolicy(tui.Preferred, tui.Maximum)
			row0.SetBorder(true)

			edit1 := tui.NewLabel("")
			edit1.SetText("test\ntesting\nfoo\nbar")
			row1 := tui.NewVBox(edit1)
			row1.SetSizePolicy(tui.Preferred, tui.Preferred)
			row1.SetBorder(true)

			edit2 := tui.NewLabel("")
			edit2.SetText("foo")
			row2 := tui.NewVBox(edit2)
			row2.SetSizePolicy(tui.Preferred, tui.Preferred)
			row2.SetBorder(true)

			b := tui.NewVBox(row0, row1, row2)
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
	{
		test: "tui.Box with title",
		setup: func() *tui.Box {
			b := tui.NewVBox(tui.NewLabel("test"))
			b.SetTitle("Title")
			b.SetBorder(true)
			return b
		},
		want: `
┌Title───┐
│test....│
│........│
│........│
└────────┘
`,
	},
	{
		test: "tui.Box with very long title",
		setup: func() *tui.Box {
			b := tui.NewVBox(tui.NewLabel("test"))
			b.SetTitle("Very long title")
			b.SetBorder(true)
			return b
		},
		want: `
┌Very lon┐
│test....│
│........│
│........│
└────────┘
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

			painter := tui.NewPainter(surface, tui.NewTheme())
			painter.Repaint(tt.setup())

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}

func TestBox_IsFocused(t *testing.T) {
	btn := tui.NewButton("Test box focus")
	box := tui.NewVBox(btn)
	want := false
	if box.IsFocused() != want {
		t.Errorf("got = \n%t\n\nwant = \n%t", box.IsFocused(), want)
	}
	btn.SetFocused(true)
	want = true
	if box.IsFocused() != want {
		t.Errorf("got = \n%t\n\nwant = \n%t", box.IsFocused(), want)
	}
}

var insertWidgetTests = []struct {
	test  string
	size  image.Point
	setup func() *tui.Box
	index int
	want  string
}{
	{
		test:  "Insert at beginning of box",
		index: 0,
		want: `
┌──────────────────┐
│Insertion.........│
│..................│
│Test 0............│
│..................│
│Test 1............│
│..................│
│Test 2............│
│..................│
└──────────────────┘
`,
	},
	{
		test:  "Insert in the middle",
		index: 1,
		want: `
┌──────────────────┐
│Test 0............│
│..................│
│Insertion.........│
│..................│
│Test 1............│
│..................│
│Test 2............│
│..................│
└──────────────────┘
`,
	},
	{
		test:  "Slice index out of range",
		index: 5,
		want: `
┌──────────────────┐
│Test 0............│
│..................│
│..................│
│Test 1............│
│..................│
│..................│
│Test 2............│
│..................│
└──────────────────┘
`,
	},
	{
		test:  "Append widget",
		index: 3,
		want: `
┌──────────────────┐
│Test 0............│
│..................│
│Test 1............│
│..................│
│Test 2............│
│..................│
│Insertion.........│
│..................│
└──────────────────┘
`,
	},
}

func TestBox_Insert(t *testing.T) {
	for _, tt := range insertWidgetTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			surface := newTestSurface(20, 10)
			painter := tui.NewPainter(surface, tui.NewTheme())

			label0 := tui.NewLabel("Test 0")
			label1 := tui.NewLabel("Test 1")
			label2 := tui.NewLabel("Test 2")

			b := tui.NewVBox(label0, label1, label2)

			insertLabel := tui.NewLabel("Insertion")
			b.Insert(tt.index, insertLabel)

			b.SetBorder(true)

			painter.Repaint(b)

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}

func TestBox_Prepend(t *testing.T) {
	want := `
┌──────────────────┐
│Prepend...........│
│..................│
│Test 0............│
│..................│
│Test 1............│
│..................│
│Test 2............│
│..................│
└──────────────────┘
`
	surface := newTestSurface(20, 10)
	painter := tui.NewPainter(surface, tui.NewTheme())

	label0 := tui.NewLabel("Test 0")
	label1 := tui.NewLabel("Test 1")
	label2 := tui.NewLabel("Test 2")

	b := tui.NewVBox(label0, label1, label2)

	label := tui.NewLabel("Prepend")
	b.Prepend(label)

	b.SetBorder(true)

	painter.Repaint(b)

	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestBox_Remove(t *testing.T) {
	want := `
┌──────────────────┐
│Test 0............│
│..................│
│Test 2............│
│..................│
└──────────────────┘
`
	surface := newTestSurface(20, 6)
	painter := tui.NewPainter(surface, tui.NewTheme())

	label0 := tui.NewLabel("Test 0")
	label1 := tui.NewLabel("Test 1")
	label2 := tui.NewLabel("Test 2")

	b := tui.NewVBox(label0, label1, label2)

	b.Remove(1)

	b.SetBorder(true)

	painter.Repaint(b)

	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}
