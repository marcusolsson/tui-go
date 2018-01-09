package tui

import (
	"image"
	"testing"
)

var zboxDrawTests = []struct {
	test  string
	size  image.Point
	setup func() Widget
	want  string
}{
	{
		test: "RudeOverwrite",
		setup: func() Widget {
			return NewZBox(
				NewLabel("long word!"),
				NewLabel("o'erwrite"),
			)
		},
		want: `
o'erwrite!
..........
..........
..........
..........
`,
	},
}

func TestZBox_Draw(t *testing.T) {
	for _, tt := range zboxDrawTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			var surface *TestSurface
			if tt.size.X == 0 && tt.size.Y == 0 {
				surface = NewTestSurface(10, 5)
			} else {
				surface = NewTestSurface(tt.size.X, tt.size.Y)
			}

			painter := NewPainter(surface, NewTheme())
			painter.Repaint(tt.setup())

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}

}
