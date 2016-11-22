package tui

import (
	"image"
	"testing"

	"github.com/kr/pretty"
	termbox "github.com/nsf/termbox-go"
)

var buttonSizeTests = []struct {
	setup    func() *Button
	size     image.Point
	sizeHint image.Point
}{
	{
		setup: func() *Button {
			return NewButton("")
		},
		size:     image.Point{0, 1},
		sizeHint: image.Point{0, 1},
	},
	{
		setup: func() *Button {
			return NewButton("test")
		},
		size:     image.Point{4, 1},
		sizeHint: image.Point{4, 1},
	},
}

func TestButton_Size(t *testing.T) {
	for _, tt := range buttonSizeTests {
		b := tt.setup()
		b.Resize(image.Point{100, 00})

		if got := b.Size(); got != tt.size {
			t.Errorf("b.Size() = %s; want = %s", got, tt.size)
		}
		if got := b.SizeHint(); got != tt.sizeHint {
			t.Errorf("b.SizeHint() = %s; want = %s", got, tt.sizeHint)
		}
	}
}

func TestButton_OnActivated(t *testing.T) {
	btn := NewButton("test")

	var invoked bool
	btn.OnActivated(func(b *Button) {
		invoked = true
	})

	ev := termbox.Event{
		Key: termbox.KeyEnter,
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
	painter := NewPainter(surface)

	btn := NewButton("test")
	btn.Resize(surface.size)
	btn.Draw(painter)

	want := `test......
..........
..........
..........
..........
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
	}
}
