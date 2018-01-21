package tui

import (
	"testing"
)

func TestList_Draw(t *testing.T) {
	surface := NewTestSurface(10, 5)
	painter := NewPainter(surface, NewTheme())

	l := NewList()
	l.AddItems("foo", "bar")
	painter.Repaint(l)

	want := `
foo       
bar       
..........
..........
..........
`

	if diff := surfaceEquals(surface, want); diff != "" {
		t.Error(diff)
	}
}

func TestList_RemoveItem(t *testing.T) {
	surface := NewTestSurface(5, 3)
	painter := NewPainter(surface, NewTheme())

	l := NewList()
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

	if diff := surfaceEquals(surface, want); diff != "" {
		t.Fatal(diff)
	}

	// Remove an item not visible.
	l.RemoveItem(3)

	painter.Repaint(l)

	if diff := surfaceEquals(surface, want); diff != "" {
		t.Error(diff)
	}

	// Selected item should not have changed.
	if l.Selected() != 1 {
		t.Errorf("got = %d; want = %d", l.Selected(), 1)
	}
}
