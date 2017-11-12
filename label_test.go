package tui_test

import (
	"image"
	"testing"
	"github.com/marcusolsson/tui-go"
)

var labelTests = []struct {
	test     string
	setup    func() *tui.Label
	size     image.Point
	sizeHint image.Point
}{
	{
		test: "Empty",
		setup: func() *tui.Label {
			return tui.NewLabel("")
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{0, 1},
	},
	{
		test: "Single word",
		setup: func() *tui.Label {
			return tui.NewLabel("test")
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{4, 1},
	},
	{
		test: "Wide word",
		setup: func() *tui.Label {
			return tui.NewLabel("あäa")
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{4, 1},
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
	setup func() *tui.Label
	want  string
}{
	{
		test: "Simple label",
		setup: func() *tui.Label {
			return tui.NewLabel("test")
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
		setup: func() *tui.Label {
			l := tui.NewLabel("this will wrap")
			l.SetWordWrap(true)
			l.SetSizePolicy(tui.Expanding, tui.Expanding)
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
		surface := newTestSurface(10, 5)

		painter := tui.NewPainter(surface, tui.NewTheme())
		painter.Repaint(tt.setup())

		if surface.String() != tt.want {
			t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
		}
	}
}
