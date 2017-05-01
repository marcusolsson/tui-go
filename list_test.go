package tui

import (
	"testing"

	"github.com/kr/pretty"
)

func TestList_Draw(t *testing.T) {
	surface := newTestSurface(10, 5)
	painter := NewPainter(surface, NewTheme())

	l := NewList()
	l.AddItems("foo", "bar")
	l.Resize(surface.size)
	l.Draw(painter)

	want := `
foo       
bar       
..........
..........
..........
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
	}
}
