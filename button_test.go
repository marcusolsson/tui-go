package tui

import (
	"testing"
)

func TestButton_OnActivated(t *testing.T) {
	btn := NewButton("test")

	var invoked bool
	btn.OnActivated(func(b *Button) {
		invoked = true
	})

	ev := KeyEvent{
		Key: KeyEnter,
	}

	t.Run("When button is not focused", func(t *testing.T) {
		btn.OnKeyEvent(ev)
		if invoked {
			t.Errorf("button should not be activated")
		}
	})

	invoked = false
	btn.SetFocused(true)

	t.Run("When button is focused", func(t *testing.T) {
		btn.OnKeyEvent(ev)
		if !invoked {
			t.Errorf("button should be activated")
		}
	})
}

func TestButton_Draw(t *testing.T) {
	surface := newTestSurface(10, 5)
	painter := NewPainter(surface, NewTheme())

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
