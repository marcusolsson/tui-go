package tui

import (
	"image"
	"testing"
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

func TestLabelSize(t *testing.T) {
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
