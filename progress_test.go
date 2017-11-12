package tui_test

import (
	"testing"

	"github.com/kr/pretty"
	"github.com/marcusolsson/tui-go"
)

func TestProgress_Draw(t *testing.T) {
	p := tui.NewProgress(100)
	p.SetSizePolicy(tui.Expanding, tui.Minimum)
	p.SetCurrent(50)

	surface := newTestSurface(11, 2)
	painter := tui.NewPainter(surface, tui.NewTheme())
	painter.Repaint(p)

	want := `
[===>-----]
...........
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
	}
}
