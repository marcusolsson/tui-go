package tui_test

import (
	"image"
	"testing"
	"github.com/marcusolsson/tui-go"
)

var drawEntryTests = []struct {
	test  string
	size  image.Point
	setup func() *tui.Entry
	want  string
}{
	{
		test: "Empty entry",
		size: image.Point{15, 2},
		setup: func() *tui.Entry {
			return tui.NewEntry()
		},
		want: `
               
...............
`,
	},
	{
		test: "Entry with text",
		size: image.Point{15, 2},
		setup: func() *tui.Entry {
			e := tui.NewEntry()
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
		setup: func() *tui.Entry {
			e := tui.NewEntry()
			e.SetText("Lorem ipsum dolor sit amet")
			e.Offset = 11
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
		setup: func() *tui.Entry {
			e := tui.NewEntry()
			e.SetText("Lorem ipsum dolor sit amet")
			e.SetFocused(true)
			e.Offset = 12
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

func TestEntry_OnChanged(t *testing.T) {
	e := tui.NewEntry()

	var invoked bool
	e.OnChanged(func(e *tui.Entry) {
		invoked = true
		if e.Text() != "t" {
			t.Errorf("e.Text() = %s; want = %s", e.Text(), "t")
		}
	})

	ev := tui.KeyEvent{
		Key:  tui.KeyRune,
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
	e := tui.NewEntry()

	var invoked bool
	e.OnSubmit(func(e *tui.Entry) {
		invoked = true
	})

	ev := tui.KeyEvent{
		Key: tui.KeyEnter,
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
	setup func() *tui.Box
	want  string
}{
	{
		test: "Preferred",
		setup: func() *tui.Box {
			e := tui.NewEntry()
			e.SetSizePolicy(tui.Preferred, tui.Preferred)

			b := tui.NewHBox(e)
			b.SetBorder(true)
			b.SetSizePolicy(tui.Expanding, tui.Expanding)

			return b
		},
		want: `
┌──────────────────┐
│                  │
│..................│
│..................│
└──────────────────┘
`,
	},
	{
		test: "Preferred/Preferred",
		setup: func() *tui.Box {
			e1 := tui.NewEntry()
			e1.SetSizePolicy(tui.Preferred, tui.Preferred)
			e1.SetText("0123456789foo")
			e1.Offset = 4

			e2 := tui.NewEntry()
			e2.SetSizePolicy(tui.Preferred, tui.Preferred)
			e2.SetText("0123456789bar")
			e2.Offset = 4

			b := tui.NewHBox(e1, e2)
			b.SetBorder(true)
			b.SetSizePolicy(tui.Expanding, tui.Expanding)

			return b
		},
		want: `
┌──────────────────┐
│456789foo456789bar│
│..................│
│..................│
└──────────────────┘
`,
	},
	{
		test: "Preferred/Minimum",
		setup: func() *tui.Box {
			e1 := tui.NewEntry()
			e1.SetSizePolicy(tui.Preferred, tui.Preferred)
			e1.SetText("0123456789foo")
			e1.Offset = 5

			e2 := tui.NewEntry()
			e2.SetSizePolicy(tui.Minimum, tui.Preferred)
			e2.SetText("0123456789bar")
			e2.Offset = 3

			b := tui.NewHBox(e1, e2)
			b.SetBorder(true)

			return b
		},
		want: `
┌──────────────────┐
│56789foo3456789bar│
│..................│
│..................│
└──────────────────┘
`,
	},
	{
		test: "tui.Minimum/tui.Preferred",
		setup: func() *tui.Box {
			e1 := tui.NewEntry()
			e1.SetSizePolicy(tui.Minimum, tui.Preferred)
			e1.SetText("0123456789foo")
			e1.Offset = 3

			e2 := tui.NewEntry()
			e2.SetSizePolicy(tui.Preferred, tui.Preferred)
			e2.SetText("0123456789bar")
			e2.Offset = 5

			b := tui.NewHBox(e1, e2)
			b.SetBorder(true)

			return b
		},
		want: `
┌──────────────────┐
│3456789foo56789bar│
│..................│
│..................│
└──────────────────┘
`,
	},
	{
		test: "Preferred/Expanding",
		setup: func() *tui.Box {
			e1 := tui.NewEntry()
			e1.SetSizePolicy(tui.Preferred, tui.Preferred)
			e1.SetText("foo")

			e2 := tui.NewEntry()
			e2.SetSizePolicy(tui.Expanding, tui.Preferred)
			e2.SetText("bar")

			b := tui.NewHBox(e1, e2)
			b.SetBorder(true)
			b.SetSizePolicy(tui.Expanding, tui.Expanding)

			return b
		},
		want: `
┌──────────────────┐
│foo       bar     │
│..................│
│..................│
└──────────────────┘
`,
	},
	{
		test: "tui.Expanding/tui.Preferred",
		setup: func() *tui.Box {
			e1 := tui.NewEntry()
			e1.SetText("foo")
			e1.SetSizePolicy(tui.Expanding, tui.Preferred)

			e2 := tui.NewEntry()
			e2.SetText("bar")
			e2.SetSizePolicy(tui.Preferred, tui.Preferred)

			b := tui.NewHBox(e1, e2)
			b.SetBorder(true)

			return b
		},
		want: `
┌──────────────────┐
│foo     bar       │
│..................│
│..................│
└──────────────────┘
`,
	},
}

func TestEntry_Layout(t *testing.T) {
	for _, tt := range layoutEntryTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			surface := newTestSurface(20, 5)
			painter := tui.NewPainter(surface, tui.NewTheme())
			painter.Repaint(tt.setup())

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}

func TestEntry_OnEvent(t *testing.T) {
	e := tui.NewEntry()
	e.SetText("Lorem ipsum")
	e.SetFocused(true)

	surface := newTestSurface(4, 1)
	painter := tui.NewPainter(surface, tui.NewTheme())
	painter.Repaint(e)

	want := `
Lore
`
	if e.Offset != 0 {
		t.Errorf("Offset = %d; want = %d", e.Offset, 0)
	}
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}

	e.OnKeyEvent(tui.KeyEvent{Key: tui.KeyRight})
	painter.Repaint(e)

	want = `
orem
`
	if e.Offset != 1 {
		t.Errorf("Offset = %d; want = %d", e.Offset, 1)
	}
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}

	e.OnKeyEvent(tui.KeyEvent{Key: tui.KeyLeft})
	e.OnKeyEvent(tui.KeyEvent{Key: tui.KeyLeft})
	painter.Repaint(e)

	want = `
Lore
`
	if e.Offset != 0 {
		t.Errorf("Offset = %d; want = %d", e.Offset, 0)
	}
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}

	repeatKeyEvent(e, tui.KeyEvent{Key: tui.KeyRight}, 20)
	painter.Repaint(e)

	want = `
sum 
`
	if e.Offset != 8 {
		t.Errorf("Offset = %d; want = %d", e.Offset, 8)
	}
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestEntry_MoveToStartAndEnd(t *testing.T) {
	t.Run("Given an entry with too long text", func(t *testing.T) {
		e := tui.NewEntry()
		e.SetText("Lorem ipsum")
		e.SetFocused(true)
		e.Offset = 6

		surface := newTestSurface(5, 1)
		painter := tui.NewPainter(surface, tui.NewTheme())

		t.Run("When cursor is moved to the start", func(t *testing.T) {
			repeatKeyEvent(e, tui.KeyEvent{Key: tui.KeyCtrlA}, 1)
			painter.Repaint(e)

			want := "\nLorem\n"

			if e.Offset != 0 {
				t.Errorf("Offset should be %d, but was %d", 0, e.Offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When cursor is moved to the end", func(t *testing.T) {
			repeatKeyEvent(e, tui.KeyEvent{Key: tui.KeyCtrlE}, 1)
			painter.Repaint(e)

			want := "\npsum \n"

			if e.Offset != 7 {
				t.Errorf("Offset should be %d, but was %d", 7, e.Offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
	})
}
func TestEntry_OnKeyBackspaceEvent(t *testing.T) {
	t.Run("Given an entry with too long text", func(t *testing.T) {
		e := tui.NewEntry()
		e.SetText("Lorem ipsum")
		e.SetFocused(true)
		e.Offset = 6

		surface := newTestSurface(5, 1)
		painter := tui.NewPainter(surface, tui.NewTheme())

		t.Run("When cursor is moved to the middle", func(t *testing.T) {
			repeatKeyEvent(e, tui.KeyEvent{Key: tui.KeyLeft}, 2)
			painter.Repaint(e)

			want := "\nm ips\n"

			if e.Offset != 4 {
				t.Errorf("Offset should be %d, but was %d", 4, e.Offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When character in the middle is deleted", func(t *testing.T) {
			repeatKeyEvent(e, tui.KeyEvent{Key: tui.KeyBackspace2}, 1)
			painter.Repaint(e)

			want := "\nm ipu\n"

			if e.Offset != 4 {
				t.Errorf("Offset should be %d, but was %d", 4, e.Offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When cursor is moved to the end", func(t *testing.T) {
			repeatKeyEvent(e, tui.KeyEvent{Key: tui.KeyRight}, 6)
			painter.Repaint(e)

			want := "\nipum \n"

			if e.Offset != 6 {
				t.Errorf("Offset should be %d, but was %d", 6, e.Offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When last character is deleted", func(t *testing.T) {
			repeatKeyEvent(e, tui.KeyEvent{Key: tui.KeyBackspace2}, 1)
			painter.Repaint(e)

			want := "\n ipu \n"

			if e.Offset != 5 {
				t.Errorf("Offset should be %d, but was %d", 5, e.Offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When all characters are deleted", func(t *testing.T) {
			repeatKeyEvent(e, tui.KeyEvent{Key: tui.KeyBackspace2}, 9)
			painter.Repaint(e)

			want := "\n     \n"

			if e.Offset != 0 {
				t.Errorf("Offset should be %d, but was %d", 0, e.Offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
		t.Run("When deleting an empty text", func(t *testing.T) {
			repeatKeyEvent(e, tui.KeyEvent{Key: tui.KeyBackspace2}, 1)
			painter.Repaint(e)

			want := "\n     \n"

			if e.Offset != 0 {
				t.Errorf("Offset should be %d, but was %d", 0, e.Offset)
			}
			if surface.String() != want {
				t.Errorf("surface should be:\n%s\nbut was:\n%s", want, surface.String())
			}
		})
	})
}

func repeatKeyEvent(e *tui.Entry, ev tui.KeyEvent, n int) {
	for i := 0; i < n; i++ {
		e.OnKeyEvent(ev)
	}
}


