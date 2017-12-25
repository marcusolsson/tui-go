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
		size:     image.Point{100, 100},
		sizeHint: image.Point{0, 1},
	},
	{
		test: "Single word",
		setup: func() *Label {
			return NewLabel("test")
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{4, 1},
	},
	{
		test: "Wide word",
		setup: func() *Label {
			return NewLabel("あäa")
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{4, 1},
	},
	{
		test: "Unicode only",
		setup: func() *Label {
			return NewLabel("∅")
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{1, 1},
	},
	{
		test: "Tall string",
		setup: func() *Label {
			return NewLabel("Null set: ∅")
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{11, 1},
	},
}

func TestLabel_Size(t *testing.T) {
	for _, tt := range labelTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			t.Parallel()

			l := tt.setup()
			l.Resize(image.Point{100, 100})

			if got := l.Size(); got != tt.size {
				t.Errorf("l.Size() = %s; want = %s", got, tt.size)
			}
			if got := l.SizeHint(); got != tt.sizeHint {
				t.Errorf("l.SizeHint() = %s; want = %s", got, tt.sizeHint)
			}
		})
	}
}

var drawLabelTests = []struct {
	test  string
	setup func() *Label
	want  string
}{
	{
		test: "Simple label",
		setup: func() *Label {
			return NewLabel("test")
		},
		want: `
test......
..........
..........
..........
..........
`,
	},
	{
		test: "Word wrap",
		setup: func() *Label {
			l := NewLabel("this will wrap")
			l.SetWordWrap(true)
			l.SetSizePolicy(Expanding, Expanding)
			return l
		},
		want: `
this will.
wrap......
..........
..........
..........
`,
	},
}

func TestLabel_Draw(t *testing.T) {
	for _, tt := range drawLabelTests {
		surface := NewTestSurface(10, 5)

		painter := NewPainter(surface, NewTheme())
		painter.Repaint(tt.setup())

		if surface.String() != tt.want {
			t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
		}
	}
}
