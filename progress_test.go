package tui

import (
	"testing"

	"github.com/kr/pretty"
)

func TestProgress_Draw(t *testing.T) {
	surface := newTestSurface(11, 2)
	painter := NewPainter(surface, NewTheme())

	p := NewProgress(100)
	p.SetSizePolicy(Expanding, Minimum)
	p.SetCurrent(50)

	p.Resize(surface.size)
	p.Draw(painter)

	want := `
[===>-----]
...........
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
	}
}
