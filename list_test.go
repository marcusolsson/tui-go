package tui_test

import (
	"testing"

	"github.com/kr/pretty"
	"github.com/marcusolsson/tui-go"
)

func TestList_Draw(t *testing.T) {
	surface := newTestSurface(10, 5)
	painter := tui.NewPainter(surface, tui.NewTheme())

	l := tui.NewList()
	l.AddItems("foo", "bar")
	painter.Repaint(l)

	want := `
foo       
bar       
..........
..........
..........
`

	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestList_RemoveItem(t *testing.T) {
	surface := newTestSurface(5, 3)
	painter := tui.NewPainter(surface, tui.NewTheme())

	l := tui.NewList()
	l.AddItems("one", "two", "three", "four", "five")
	l.SetSelected(1)

	painter.Repaint(l)

	want := `
one  
two  
three
`

	// Make sure okay before removing any items.
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}

	// Remove a visible item.
	l.RemoveItem(2)

	painter.Repaint(l)

	want = `
one  
two  
four 
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
		return
	}

	// Remove an item not visible.
	l.RemoveItem(3)

	painter.Repaint(l)

	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}

	// Selected item should not have changed.
	if l.Selected() != 1 {
		t.Error(pretty.Diff(l.Selected, 1))
	}
}
