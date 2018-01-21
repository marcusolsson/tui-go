package tui

import (
	"image"
	"testing"
)

var drawTextEditTests = []struct {
	test  string
	size  image.Point
	setup func() *TextEdit
	want  string
}{
	{
		test: "Simple",
		size: image.Point{15, 5},
		setup: func() *TextEdit {
			e := NewTextEdit()
			e.SetText("Lorem ipsum dolor sit amet")
			e.SetWordWrap(true)
			return e
		},
		want: `
Lorem ipsum    
dolor sit amet 
...............
...............
...............
`,
	},
}

func TestTextEdit_Draw(t *testing.T) {
	for _, tt := range drawTextEditTests {
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

			if diff := surfaceEquals(surface, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
