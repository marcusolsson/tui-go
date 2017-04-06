package tui

import (
	"image"
	"testing"

	"github.com/kr/pretty"
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

func TestEntry_Draw(t *testing.T) {
	surface := newTestSurface(15, 5)
	painter := NewPainter(surface, NewPalette())

	e := NewEntry()
	e.Resize(surface.size)
	e.Draw(painter)

	want := `         ......
...............
...............
...............
...............
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
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
