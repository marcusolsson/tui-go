package tui

import (
	"image"
	"testing"
)

var entrySizeTests = []struct {
	setup    func() *Entry
	size     image.Point
	sizeHint image.Point
}{
	{
		setup: func() *Entry {
			return NewEntry()
		},
		size:     image.Point{10, 1},
		sizeHint: image.Point{10, 1},
	},
}

func TestEntry_Size(t *testing.T) {
	for _, tt := range entrySizeTests {
		e := tt.setup()
		e.Resize(image.Point{100, 00})

		if got := e.Size(); got != tt.size {
			t.Errorf("e.Size() = %s; want = %s", got, tt.size)
		}
		if got := e.SizeHint(); got != tt.sizeHint {
			t.Errorf("e.SizeHint() = %s; want = %s", got, tt.sizeHint)
		}
	}
}

var drawEntryTests = []struct {
	test  string
	size  image.Point
	setup func() *Entry
	want  string
}{
	{
		test: "Simple",
		size: image.Point{15, 5},
		setup: func() *Entry {
			return NewEntry()
		},
		want: `          .....
...............
...............
...............
...............
`,
	},
	{
		test: "Entry with text",
		size: image.Point{15, 5},
		setup: func() *Entry {
			e := NewEntry()
			e.SetText("test")
			return e
		},
		want: `test      .....
...............
...............
...............
...............
`,
	},
	{
		test: "Scrolling entry",
		size: image.Point{15, 5},
		setup: func() *Entry {
			e := NewEntry()
			e.SetText("Lorem ipsum dolor sit amet")
			return e
		},
		want: `r sit amet.....
...............
...............
...............
...............
`,
	},
	{
		test: "Scrolling entry when focused",
		size: image.Point{15, 5},
		setup: func() *Entry {
			e := NewEntry()
			e.SetText("Lorem ipsum dolor sit amet")
			e.SetFocused(true)
			return e
		},
		want: ` sit amet .....
...............
...............
...............
...............
`,
	},
}

func TestEntry_Draw(t *testing.T) {
	for _, tt := range drawEntryTests {
		var surface *testSurface
		if tt.size.X == 0 && tt.size.Y == 0 {
			surface = newTestSurface(10, 5)
		} else {
			surface = newTestSurface(tt.size.X, tt.size.Y)
		}
		painter := NewPainter(surface, NewPalette())

		b := tt.setup()

		b.Resize(surface.size)
		b.Draw(painter)

		if surface.String() != tt.want {
			t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
		}
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

	ev := Event{
		Type: EventKey,
		Ch:   't',
	}

	t.Run("When entry is not focused", func(t *testing.T) {
		e.OnEvent(ev)
		if invoked {
			t.Errorf("entry should not be submitted")
		}
	})

	invoked = false
	e.SetFocused(true)

	t.Run("When entry is focused", func(t *testing.T) {
		e.OnEvent(ev)
		if !invoked {
			t.Errorf("entry should be submitted")
		}
	})
}

func TestEntry_OnSubmit(t *testing.T) {
	e := NewEntry()

	var invoked bool
	e.OnSubmit(func(e *Entry) {
		invoked = true
	})

	ev := Event{
		Type: EventKey,
		Key:  KeyEnter,
	}

	t.Run("When entry is not focused", func(t *testing.T) {
		e.OnEvent(ev)
		if invoked {
			t.Errorf("entry should not be submitted")
		}
	})

	invoked = false
	e.SetFocused(true)

	t.Run("When entry is focused", func(t *testing.T) {
		e.OnEvent(ev)
		if !invoked {
			t.Errorf("button should be submitted")
		}
	})
}
