package tui

import (
	"testing"
)

func TestProgress_Draw(t *testing.T) {
	p := NewProgress(100)
	p.SetSizePolicy(Expanding, Minimum)
	p.SetCurrent(50)

	surface := NewTestSurface(11, 2)
	painter := NewPainter(surface, NewTheme())
	painter.Repaint(p)

	want := `
[===>-----]
...........
`

	if diff := surfaceEquals(surface, want); diff != "" {
		t.Error(diff)
	}
}
