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
│        │
│        │
│        │
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
│test    │
│        │
│        │
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
│testfoo │
│        │
│        │
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
│        │
│        │
│        │
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
│test    │
│        │
│        │
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
│test    │
│        │
│foo     │
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
│            ┌────┐            │
│            │test│            │
│            └────┘            │
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
││test         ││test         ││
││             ││             ││
││             ││             ││
││             ││             ││
││             ││             ││
││             ││             ││
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
│││test        ││test        │││
│││            ││            │││
│││            ││            │││
│││            ││            │││
│││            ││            │││
│││            ││            │││
││└────────────┘└────────────┘││
│└────────────────────────────┘│
│┌────────────────────────────┐│
││┌────────────┐┌────────────┐││
│││test        ││test        │││
│││            ││            │││
│││            ││            │││
│││            ││            │││
│││            ││            │││
│││            ││            │││
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
││test                        ││
││testing                     ││
││foo                         ││
││bar                         ││
│└────────────────────────────┘│
│┌────────────────────────────┐│
││test                        ││
││testing                     ││
││foo                         ││
││bar                         ││
││                            ││
││                            ││
│└────────────────────────────┘│
│┌────────────────────────────┐│
││foo                         ││
││                            ││
││                            ││
││                            ││
││                            ││
││                            ││
│└────────────────────────────┘│
└──────────────────────────────┘
`,
	},
	{
		test: "Box with title",
		setup: func() *Box {
			b := NewVBox(NewLabel("test"))
			b.SetTitle("Title")
			b.SetBorder(true)
			return b
		},
		want: `
┌Title───┐
│test    │
│        │
│        │
└────────┘
`,
	},
	{
		test: "Box with very long title",
		setup: func() *Box {
			b := NewVBox(NewLabel("test"))
			b.SetTitle("Very long title")
			b.SetBorder(true)
			return b
		},
		want: `
┌Very lon┐
│test    │
│        │
│        │
└────────┘
`,
	},
}

func TestBox_Draw(t *testing.T) {
	for _, tt := range drawBoxTests {
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

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}

var styleBoxTests = []struct {
	test            string
	setup           func() *Box
	theme           func() *Theme
	wantContents    string
	wantColors      string
	wantDecorations string
}{
	{
		test: "Red horizontal box",
		setup: func() *Box {
			b := NewHBox()
			b.SetBorder(true)
			return b
		},
		theme: func() *Theme {
			t := NewTheme()
			t.SetStyle("box", Style{
				Fg: Color(3),
			})
			return t
		},
		wantContents: `
┌────────┐
│        │
│        │
│        │
└────────┘
`,
		wantColors: `
3333333333
3333333333
3333333333
3333333333
3333333333
`,
	},
	{
		test: "red box, styled and unstyled labels",
		setup: func() *Box {
			unstyled := NewLabel("unstyled")
			styled := NewLabel("styled")
			styled.SetStyleName("blue")
			box := NewVBox(unstyled, styled)
			box.SetBorder(true)
			return box
		},
		theme: func() *Theme {
			t := NewTheme()
			t.SetStyle("box", Style{
				Fg:   Color(3),
				Bold: DecorationOn,
			})
			t.SetStyle("label.blue", Style{
				Fg: Color(4),
			})
			return t
		},
		wantContents: `
┌────────┐
│unstyled│
│        │
│styled  │
└────────┘
`,
		wantColors: `
3333333333
3333333333
3333333333
3444444333
3333333333
`,
		wantDecorations: `
2222222222
2222222222
2222222222
2222222222
2222222222
`,
	},
	{
		test: "Styled box, labels inherit",
		setup: func() *Box {
			return NewVBox(
				NewLabel("label 1"),
				NewLabel("label 2"),
			)
		},
		theme: func() *Theme {
			t := NewTheme()
			t.SetStyle("box", Style{
				Fg:   Color(3),
				Bold: DecorationOn,
			})
			return t
		},
		wantContents: `
label 1   
          
          
label 2   
          
`,
		wantColors: `
3333333333
3333333333
3333333333
3333333333
3333333333
`,
		wantDecorations: `
2222222222
2222222222
2222222222
2222222222
2222222222
`,
	},
	{
		test: "blue box, bold border",
		setup: func() *Box {
			r := NewVBox(
				NewLabel("label 1"),
				NewLabel("label 2"),
			)
			r.SetBorder(true)
			return r
		},
		theme: func() *Theme {
			t := NewTheme()
			t.SetStyle("box", Style{
				Fg: Color(3),
			})
			t.SetStyle("box.border", Style{
				Bold: DecorationOn,
			})
			return t
		},
		wantContents: `
┌────────┐
│label 1 │
│        │
│label 2 │
└────────┘
`,
		wantColors: `
3333333333
3333333333
3333333333
3333333333
3333333333
`,
		wantDecorations: `
2222222222
2000000002
2000000002
2000000002
2222222222
`,
	},
}

func TestBox_Style(t *testing.T) {
	for _, tt := range styleBoxTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			surface := NewTestSurface(10, 5)
			painter := NewPainter(surface, tt.theme())
			painter.Repaint(tt.setup())

			if tt.wantContents != "" && surface.String() != tt.wantContents {
				t.Errorf("wrong contents: got = \n%s\n\nwant = \n%s", surface.String(), tt.wantContents)
			}
			if tt.wantColors != "" && surface.FgColors() != tt.wantColors {
				t.Errorf("wrong colors: got = \n%s\n\nwant = \n%s", surface.FgColors(), tt.wantColors)
			}
			if tt.wantDecorations != "" && surface.Decorations() != tt.wantDecorations {
				t.Errorf("wrong decorations: got = \n%s\n\nwant = \n%s", surface.Decorations(), tt.wantDecorations)
			}
		})
	}
}

func TestBox_IsFocused(t *testing.T) {
	btn := NewButton("Test box focus")
	box := NewVBox(btn)
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
	setup func() *Box
	index int
	want  string
}{
	{
		test:  "Insert at beginning of box",
		index: 0,
		want: `
┌──────────────────┐
│Insertion         │
│                  │
│Test 0            │
│                  │
│Test 1            │
│                  │
│Test 2            │
│                  │
└──────────────────┘
`,
	},
	{
		test:  "Insert in the middle",
		index: 1,
		want: `
┌──────────────────┐
│Test 0            │
│                  │
│Insertion         │
│                  │
│Test 1            │
│                  │
│Test 2            │
│                  │
└──────────────────┘
`,
	},
	{
		test:  "Slice index out of range",
		index: 5,
		want: `
┌──────────────────┐
│Test 0            │
│                  │
│                  │
│Test 1            │
│                  │
│                  │
│Test 2            │
│                  │
└──────────────────┘
`,
	},
	{
		test:  "Append widget",
		index: 3,
		want: `
┌──────────────────┐
│Test 0            │
│                  │
│Test 1            │
│                  │
│Test 2            │
│                  │
│Insertion         │
│                  │
└──────────────────┘
`,
	},
}

func TestBox_Insert(t *testing.T) {
	for _, tt := range insertWidgetTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			surface := NewTestSurface(20, 10)
			painter := NewPainter(surface, NewTheme())

			label0 := NewLabel("Test 0")
			label1 := NewLabel("Test 1")
			label2 := NewLabel("Test 2")

			b := NewVBox(label0, label1, label2)

			insertLabel := NewLabel("Insertion")
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
│Prepend           │
│                  │
│Test 0            │
│                  │
│Test 1            │
│                  │
│Test 2            │
│                  │
└──────────────────┘
`
	surface := NewTestSurface(20, 10)
	painter := NewPainter(surface, NewTheme())

	label0 := NewLabel("Test 0")
	label1 := NewLabel("Test 1")
	label2 := NewLabel("Test 2")

	b := NewVBox(label0, label1, label2)

	label := NewLabel("Prepend")
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
│Test 0            │
│                  │
│Test 2            │
│                  │
└──────────────────┘
`
	surface := NewTestSurface(20, 6)
	painter := NewPainter(surface, NewTheme())

	label0 := NewLabel("Test 0")
	label1 := NewLabel("Test 1")
	label2 := NewLabel("Test 2")

	b := NewVBox(label0, label1, label2)

	b.Remove(1)
	b.Remove(10)

	b.SetBorder(true)

	painter.Repaint(b)

	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}
