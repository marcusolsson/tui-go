package tui

import (
	"image"
	"testing"

	"github.com/kr/pretty"
)

var labelTests = []struct {
	test     string
	setup    func() *Label
	size     image.Point
	sizeHint image.Point
}{
	{
		test: "Empty",
		setup: func() *Label {
			return NewLabel("")
		},
		size:     image.Point{0, 1},
		sizeHint: image.Point{0, 1},
	},
	{
		test: "Single word",
		setup: func() *Label {
			return NewLabel("test")
		},
		size:     image.Point{4, 1},
		sizeHint: image.Point{4, 1},
	},
}

func TestLabel_Size(t *testing.T) {
	for _, tt := range labelTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			t.Parallel()

			l := tt.setup()
			l.Resize(image.Point{100, 00})

			if got := l.Size(); got != tt.size {
				t.Errorf("l.Size() = %s; want = %s", got, tt.size)
			}
			if got := l.SizeHint(); got != tt.sizeHint {
				t.Errorf("l.SizeHint() = %s; want = %s", got, tt.sizeHint)
			}
		})
	}
}

func TestLabel_Draw(t *testing.T) {
	surface := newTestSurface(10, 5)
	painter := NewPainter(surface)

	label := NewLabel("test")
	label.Resize(surface.size)
	label.Draw(painter)

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
