package tui

import (
	"image"
	"testing"
)

var buttonSizeTests = []struct {
	setup       func() *Button
	minSizeHint image.Point
	sizeHint    image.Point
	size        image.Point
}{
	{
		setup: func() *Button {
			return NewButton("")
		},
		minSizeHint: image.Point{1, 1},
		sizeHint:    image.Point{1, 1},
		size:        image.Point{100, 100},
	},
	{
		setup: func() *Button {
			return NewButton("test")
		},
		minSizeHint: image.Point{1, 1},
		sizeHint:    image.Point{4, 1},
		size:        image.Point{100, 100},
	},
	{
		setup: func() *Button {
			return NewButton("あäa")
		},
		minSizeHint: image.Point{1, 1},
		sizeHint:    image.Point{4, 1},
		size:        image.Point{100, 100},
	},
}

func TestButton_Size(t *testing.T) {
	for _, tt := range buttonSizeTests {
		t.Run("", func(t *testing.T) {
			b := tt.setup()
			b.Resize(image.Point{100, 100})

			if got := b.Size(); got != tt.size {
				t.Errorf("b.Size() = %s; want = %s", got, tt.size)
			}
			if got := b.SizeHint(); got != tt.sizeHint {
				t.Errorf("b.SizeHint() = %s; want = %s", got, tt.sizeHint)
			}
		})
	}
}

func TestButton_OnActivated(t *testing.T) {
	btn := NewButton("test")

	var invoked bool
	btn.OnActivated(func(b *Button) {
		invoked = true
	})

	ev := Event{
		Type: EventKey,
		Key:  KeyEnter,
	}

	t.Run("When button is not focused", func(t *testing.T) {
		btn.OnEvent(ev)
		if invoked {
			t.Errorf("button should not be activated")
		}
	})

	invoked = false
	btn.SetFocused(true)

	t.Run("When button is focused", func(t *testing.T) {
		btn.OnEvent(ev)
		if !invoked {
			t.Errorf("button should be activated")
		}
	})
}

func TestButton_Draw(t *testing.T) {
	surface := newTestSurface(10, 5)
	painter := NewPainter(surface, NewPalette())

	btn := NewButton("test")
	btn.Resize(surface.size)
	btn.Draw(painter)

	want := `
test      
..........
..........
..........
..........
`

	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}
