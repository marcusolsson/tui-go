package tui

import (
	"image"
	"testing"
)

var drawEntryTests = []struct {
	test  string
	size  image.Point
	setup func() *Entry
	want  string
}{
	{
		test: "Empty entry",
		size: image.Point{15, 2},
		setup: func() *Entry {
			return NewEntry()
		},
		want: `
               
...............
`,
	},
	{
		test: "Entry with text",
		size: image.Point{15, 2},
		setup: func() *Entry {
			e := NewEntry()
			e.SetText("test")
			return e
		},
		want: `
test           
...............
`,
	},
	{
		test: "Scrolling entry",
		size: image.Point{15, 2},
		setup: func() *Entry {
			e := NewEntry()
			e.SetText("Lorem ipsum dolor sit amet")
			e.offset = 11
			return e
		},
		want: `
 dolor sit amet
...............
`,
	},
	{
		test: "Scrolling entry when focused",
		size: image.Point{15, 2},
		setup: func() *Entry {
			e := NewEntry()
			e.SetText("Lorem ipsum dolor sit amet")
			e.SetFocused(true)
			e.offset = 12
			return e
		},
		want: `
dolor sit amet 
...............
`,
	},
}

func TestEntry_Draw(t *testing.T) {
	for _, tt := range drawEntryTests {
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

func TestEntry_OnChanged(t *testing.T) {
	e := NewEntry()

	var invoked bool
	e.OnChanged(func(e *Entry) {
		invoked = true
		if e.Text() != "t" {
			t.Errorf("e.Text() = %s; want = %s", e.Text(), "t")
		}
	})

	ev := KeyEvent{
		Key:  KeyRune,
		Rune: 't',
	}

	t.Run("When entry is not focused", func(t *testing.T) {
		e.OnKeyEvent(ev)
		if invoked {
			t.Errorf("entry should not receive key events")
		}
	})

	invoked = false
	e.SetFocused(true)

	t.Run("When entry is focused", func(t *testing.T) {
		e.OnKeyEvent(ev)
		if !invoked {
			t.Errorf("entry should receive key events")
		}
	})
}

func TestEntry_OnSubmit(t *testing.T) {
	e := NewEntry()

	var invoked bool
	e.OnSubmit(func(e *Entry) {
		invoked = true
	})

	ev := KeyEvent{
		Key: KeyEnter,
	}

	t.Run("When entry is not focused", func(t *testing.T) {
		e.OnKeyEvent(ev)
		if invoked {
			t.Errorf("entry should not be submitted")
		}
	})

	invoked = false
	e.SetFocused(true)

	t.Run("When entry is focused", func(t *testing.T) {
		e.OnKeyEvent(ev)
		if !invoked {
			t.Errorf("button should be submitted")
		}
	})
}

var layoutEntryTests = []struct {
	test  string
	setup func() *Box
	want  string
}{
	{
		test: "Preferred",
		setup: func() *Box {
			e := NewEntry()
			e.SetSizePolicy(Preferred, Preferred)

			b := NewHBox(e)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)

			return b
		},
		want: `
┌──────────────────┐
│                  │
│                  │
│                  │
└──────────────────┘
`,
	},
	{
		test: "Preferred/Preferred",
		setup: func() *Box {
			e1 := NewEntry()
			e1.SetSizePolicy(Preferred, Preferred)
			e1.SetText("0123456789foo")
			e1.offset = 4

			e2 := NewEntry()
			e2.SetSizePolicy(Preferred, Preferred)
			e2.SetText("0123456789bar")
			e2.offset = 4

			b := NewHBox(e1, e2)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)

			return b
		},
		want: `
┌──────────────────┐
│456789foo456789bar│
│                  │
│                  │
└──────────────────┘
`,
	},
	{
		test: "Preferred/Minimum",
		setup: func() *Box {
			e1 := NewEntry()
			e1.SetSizePolicy(Preferred, Preferred)
			e1.SetText("0123456789foo")
			e1.offset = 5

			e2 := NewEntry()
			e2.SetSizePolicy(Minimum, Preferred)
			e2.SetText("0123456789bar")
			e2.offset = 3

			b := NewHBox(e1, e2)
			b.SetBorder(true)

			return b
		},
		want: `
┌──────────────────┐
│56789foo3456789bar│
│                  │
│                  │
└──────────────────┘
`,
	},
	{
		test: "Minimum/Preferred",
		setup: func() *Box {
			e1 := NewEntry()
			e1.SetSizePolicy(Minimum, Preferred)
			e1.SetText("0123456789foo")
			e1.offset = 3

			e2 := NewEntry()
			e2.SetSizePolicy(Preferred, Preferred)
			e2.SetText("0123456789bar")
			e2.offset = 5

			b := NewHBox(e1, e2)
			b.SetBorder(true)

			return b
		},
		want: `
┌──────────────────┐
│3456789foo56789bar│
│                  │
│                  │
└──────────────────┘
`,
	},
	{
		test: "Preferred/Expanding",
		setup: func() *Box {
			e1 := NewEntry()
			e1.SetSizePolicy(Preferred, Preferred)
			e1.SetText("foo")

			e2 := NewEntry()
			e2.SetSizePolicy(Expanding, Preferred)
			e2.SetText("bar")

			b := NewHBox(e1, e2)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)

			return b
		},
		want: `
┌──────────────────┐
│foo       bar     │
│                  │
│                  │
└──────────────────┘
`,
	},
	{
		test: "Expanding/Preferred",
		setup: func() *Box {
			e1 := NewEntry()
			e1.SetText("foo")
			e1.SetSizePolicy(Expanding, Preferred)

			e2 := NewEntry()
			e2.SetText("bar")
			e2.SetSizePolicy(Preferred, Preferred)

			b := NewHBox(e1, e2)
			b.SetBorder(true)

			return b
		},
		want: `
┌──────────────────┐
│foo     bar       │
│                  │
│                  │
└──────────────────┘
`,
	},
}

func TestEntry_Layout(t *testing.T) {
	for _, tt := range layoutEntryTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			surface := NewTestSurface(20, 5)
			painter := NewPainter(surface, NewTheme())
			painter.Repaint(tt.setup())

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}

func TestEntry_OnEvent(t *testing.T) {
	e := NewEntry()
	e.SetText("Lorem ipsum")
	e.SetFocused(true)

	surface := NewTestSurface(4, 1)
	painter := NewPainter(surface, NewTheme())
	painter.Repaint(e)

	want := `
Lore
`
	if e.offset != 0 {
		t.Errorf("offset = %d; want = %d", e.offset, 0)
	}
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}

	e.OnKeyEvent(KeyEvent{Key: KeyRight})
	painter.Repaint(e)

	want = `
orem
`
	if e.offset != 1 {
		t.Errorf("offset = %d; want = %d", e.offset, 1)
	}
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}

	e.OnKeyEvent(KeyEvent{Key: KeyLeft})
	e.OnKeyEvent(KeyEvent{Key: KeyLeft})
	painter.Repaint(e)

	want = `
Lore
`
	if e.offset != 0 {
		t.Errorf("offset = %d; want = %d", e.offset, 0)
	}
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}

	repeatKeyEvent(e, KeyEvent{Key: KeyRight}, 20)
	painter.Repaint(e)

	want = `
sum 
`
	if e.offset != 8 {
		t.Errorf("offset = %d; want = %d", e.offset, 8)
	}
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestEntry_MoveToStartAndEnd(t *testing.T) {
	t.Run("Given an entry with too long text", func(t *testing.T) {
		e := NewEntry()
		e.SetText("Lorem ipsum")
		e.SetFocused(true)
		e.text.SetMaxWidth(5)
		e.offset = 6

		surface := NewTestSurface(5, 1)
		painter := NewPainter(surface, NewTheme())

		t.Run("When cursor is moved to the start", func(t *testing.T) {
			repeatKeyEvent(e, KeyEvent{Key: KeyCtrlA}, 1)
			painter.Repaint(e)

			want := "\nLorem\n"

			if got := e.text.CursorPos(); got.X != 0 {
				t.Errorf("cursor position should be %d, but was %d", 0, got.X)
			}
			if e.offset != 0 {
				t.Errorf("offset should be %d, but was %d", 0, e.offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When cursor is moved to the end", func(t *testing.T) {
			repeatKeyEvent(e, KeyEvent{Key: KeyCtrlE}, 1)
			painter.Repaint(e)

			want := "\npsum \n"

			if got := e.text.CursorPos(); got.X != 11 {
				t.Errorf("cursor position should be %d, but was %d", 11, got.X)
			}
			if e.offset != 7 {
				t.Errorf("offset should be %d, but was %d", 7, e.offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
	})
}
func TestEntry_OnKeyBackspaceEvent(t *testing.T) {
	t.Run("Given an entry with too long text", func(t *testing.T) {
		e := NewEntry()
		e.SetText("Lorem ipsum")
		e.SetFocused(true)
		e.text.SetMaxWidth(5)
		e.offset = 6

		surface := NewTestSurface(5, 1)
		painter := NewPainter(surface, NewTheme())

		t.Run("When cursor is moved to the middle", func(t *testing.T) {
			repeatKeyEvent(e, KeyEvent{Key: KeyLeft}, 2)
			painter.Repaint(e)

			want := "\nm ips\n"

			if got := e.text.CursorPos(); got.X != 9 {
				t.Errorf("cursor position should be %d, but was %d", 9, got.X)
			}
			if e.offset != 4 {
				t.Errorf("offset should be %d, but was %d", 4, e.offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When character in the middle is deleted", func(t *testing.T) {
			repeatKeyEvent(e, KeyEvent{Key: KeyBackspace2}, 1)
			painter.Repaint(e)

			want := "\nm ipu\n"

			if e.offset != 4 {
				t.Errorf("offset should be %d, but was %d", 4, e.offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When cursor is moved to the end", func(t *testing.T) {
			repeatKeyEvent(e, KeyEvent{Key: KeyRight}, 6)
			painter.Repaint(e)

			want := "\nipum \n"

			if e.offset != 6 {
				t.Errorf("offset should be %d, but was %d", 6, e.offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When last character is deleted", func(t *testing.T) {
			repeatKeyEvent(e, KeyEvent{Key: KeyBackspace2}, 1)
			painter.Repaint(e)

			want := "\n ipu \n"

			if e.offset != 5 {
				t.Errorf("offset should be %d, but was %d", 5, e.offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When all characters are deleted", func(t *testing.T) {
			repeatKeyEvent(e, KeyEvent{Key: KeyBackspace2}, 9)
			painter.Repaint(e)

			want := "\n     \n"

			if e.offset != 0 {
				t.Errorf("offset should be %d, but was %d", 0, e.offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When deleting an empty text", func(t *testing.T) {
			repeatKeyEvent(e, KeyEvent{Key: KeyBackspace2}, 1)
			painter.Repaint(e)

			want := "\n     \n"

			if e.offset != 0 {
				t.Errorf("offset should be %d, but was %d", 0, e.offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
	})
}

func TestIsTextRemaining(t *testing.T) {
	for _, tt := range []struct {
		text   string
		offset int
		width  int
		want   bool
	}{
		{"Lorem ipsum", 0, 11, false},
		{"Lorem ipsum", 1, 11, false},
		{"Lorem ipsum", 0, 10, true},
		{"Lorem ipsum", 5, 5, true},
	} {
		t.Run("", func(t *testing.T) {
			e := NewEntry()
			e.SetText(tt.text)
			e.SetFocused(true)
			e.Resize(image.Pt(tt.width, 1))

			e.offset = tt.offset

			if e.isTextRemaining() != tt.want {
				t.Fatalf("want = %v; got = %v", tt.want, e.isTextRemaining())
			}
		})
	}
}

func repeatKeyEvent(e *Entry, ev KeyEvent, n int) {
	for i := 0; i < n; i++ {
		e.OnKeyEvent(ev)
	}
}
