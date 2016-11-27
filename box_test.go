package tui

import (
	"image"
	"testing"
)

var verticalBoxSizeTests = []struct {
	test     string
	setup    func() *VBox
	size     image.Point
	sizeHint image.Point
}{
	{
		test: "Stretch empty box",
		setup: func() *VBox {
			b := NewVBox()
			b.SetBorder(true)
			return b
		},
		size:     image.Point{2, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "Stretch empty box",
		setup: func() *VBox {
			b := NewVBox()
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Minimum)
			return b
		},
		size:     image.Point{100, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "Stretch empty box",
		setup: func() *VBox {
			b := NewVBox()
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)
			return b
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "No stretch",
		setup: func() *VBox {
			b := NewVBox(
				NewLabel("test"),
				NewLabel("another test"),
			)
			b.SetBorder(true)
			return b
		},
		size:     image.Point{14, 4},
		sizeHint: image.Point{14, 4},
	},
	{
		test: "Stretchy width",
		setup: func() *VBox {
			b := NewVBox(
				NewLabel("test"),
				NewLabel("another test"),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Minimum)
			return b
		},
		size:     image.Point{100, 4},
		sizeHint: image.Point{14, 4},
	},
	{
		test: "Stretchy width and height",
		setup: func() *VBox {
			b := NewVBox(
				NewLabel("test"),
				NewLabel("another test"),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)
			return b
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{14, 4},
	},
}

func TestVBox_Size(t *testing.T) {
	for _, tt := range verticalBoxSizeTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			t.Parallel()

			b := tt.setup()
			b.Resize(image.Point{100, 100})

			if got := b.Size(); got != tt.size {
				t.Errorf("b.Size() = %s; want = %s", got, tt.size)
			}
			if got := b.SizeHint(); got != tt.sizeHint {
				t.Errorf("b.SizeHint() = %s; want = %s", got, tt.sizeHint)
			}
		})
	}
}

var drawVBoxTests = []struct {
	test  string
	setup func() *VBox
	want  string
}{
	{
		test: "Empty box",
		setup: func() *VBox {
			b := NewVBox()
			b.SetBorder(true)
			return b
		},
		want: `┌┐........
└┘........
..........
..........
..........
`,
	},
	{
		test: "Box containing one widget",
		setup: func() *VBox {
			b := NewVBox(
				NewLabel("test"),
			)
			b.SetBorder(true)
			return b
		},
		want: `┌────┐....
│test│....
└────┘....
..........
..........
`,
	},
	{
		test: "Box expands horizontally",
		setup: func() *VBox {
			b := NewVBox(
				NewLabel("test"),
				NewLabel("foo"),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Minimum)
			return b
		},
		want: `┌────────┐
│test....│
│foo.....│
└────────┘
..........
`,
	},
	{
		test: "Box expands vertically",
		setup: func() *VBox {
			b := NewVBox(
				NewLabel("test"),
				NewLabel("foo"),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Minimum, Expanding)
			return b
		},
		want: `┌────┐....
│test│....
│foo.│....
│....│....
└────┘....
`,
	},
	{
		test: "Box expands along both axes",
		setup: func() *VBox {
			b := NewVBox(
				NewLabel("test"),
				NewLabel("foo"),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)
			return b
		},
		want: `┌────────┐
│test....│
│foo.....│
│........│
└────────┘
`,
	},
}

func TestVBox_Draw(t *testing.T) {
	for _, tt := range drawVBoxTests {
		surface := newTestSurface(10, 5)
		painter := NewPainter(surface, NewPalette())

		b := tt.setup()

		b.Resize(surface.size)
		b.Draw(painter)

		if surface.String() != tt.want {
			t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
		}
	}
}

var horizontalBoxSizeTests = []struct {
	test     string
	setup    func() *HBox
	size     image.Point
	sizeHint image.Point
}{
	{
		test: "Stretch empty box",
		setup: func() *HBox {
			b := NewHBox()
			b.SetBorder(true)
			return b
		},
		size:     image.Point{2, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "Stretch empty box",
		setup: func() *HBox {
			b := NewHBox()
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Minimum)
			return b
		},
		size:     image.Point{100, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "Stretch empty box",
		setup: func() *HBox {
			b := NewHBox()
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)
			return b
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "No stretch",
		setup: func() *HBox {
			b := NewHBox(
				NewLabel("test"),
				NewLabel("another test"),
			)
			b.SetBorder(true)
			return b
		},
		size:     image.Point{18, 3},
		sizeHint: image.Point{18, 3},
	},
	{
		test: "Stretchy width",
		setup: func() *HBox {
			b := NewHBox(
				NewLabel("test"),
				NewLabel("another test"),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Minimum)
			return b
		},
		size:     image.Point{100, 3},
		sizeHint: image.Point{18, 3},
	},
	{
		test: "Stretchy width and height",
		setup: func() *HBox {
			b := NewHBox(
				NewLabel("test"),
				NewLabel("another test"),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)
			return b
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{18, 3},
	},
	{
		test: "Nested box",
		setup: func() *HBox {
			nested := NewHBox(NewLabel("test"))
			nested.SetBorder(true)
			nested.SetSizePolicy(Expanding, Minimum)

			b := NewHBox(nested)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)

			return b
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{8, 5},
	},
}

func TestHBox_Size(t *testing.T) {
	for _, tt := range horizontalBoxSizeTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			t.Parallel()

			b := tt.setup()
			b.Resize(image.Point{100, 100})

			if got := b.Size(); got != tt.size {
				t.Errorf("b.Size() = %s; want = %s", got, tt.size)
			}
			if got := b.SizeHint(); got != tt.sizeHint {
				t.Errorf("b.SizeHint() = %s; want = %s", got, tt.sizeHint)
			}
		})
	}
}

var drawHBoxTests = []struct {
	test  string
	setup func() *HBox
	want  string
}{
	{
		test: "Empty box",
		setup: func() *HBox {
			b := NewHBox()
			b.SetBorder(true)
			return b
		},
		want: `┌┐........
└┘........
..........
..........
..........
`,
	},
	{
		test: "Box containing one widget",
		setup: func() *HBox {
			b := NewHBox(
				NewLabel("test"),
			)
			b.SetBorder(true)
			return b
		},
		want: `┌────┐....
│test│....
└────┘....
..........
..........
`,
	},
	{
		test: "Box expands horizontally",
		setup: func() *HBox {
			b := NewHBox(
				NewLabel("test"),
				NewLabel("foo"),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Minimum)
			return b
		},
		want: `┌────────┐
│testfoo.│
└────────┘
..........
..........
`,
	},
	{
		test: "Box expands vertically",
		setup: func() *HBox {
			b := NewHBox(
				NewLabel("test"),
				NewLabel("foo"),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Minimum, Expanding)
			return b
		},
		want: `┌───────┐.
│testfoo│.
│.......│.
│.......│.
└───────┘.
`,
	},
	{
		test: "Box expands along both axes",
		setup: func() *HBox {
			b := NewHBox(
				NewLabel("test"),
				NewLabel("foo"),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)
			return b
		},
		want: `┌────────┐
│testfoo.│
│........│
│........│
└────────┘
`,
	},
}

func TestHBox_Draw(t *testing.T) {
	for _, tt := range drawHBoxTests {
		surface := newTestSurface(10, 5)
		painter := NewPainter(surface, NewPalette())

		b := tt.setup()

		b.Resize(surface.size)
		b.Draw(painter)

		if surface.String() != tt.want {
			t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
		}
	}
}
