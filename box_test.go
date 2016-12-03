package tui

import (
	"image"
	"testing"
)

var verticalBoxSizeTests = []struct {
	test     string
	setup    func() *Box
	size     image.Point
	sizeHint image.Point
}{
	{
		test: "Stretch empty box",
		setup: func() *Box {
			b := NewVBox()
			b.SetBorder(true)
			return b
		},
		size:     image.Point{2, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "Stretch empty box",
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
	setup func() *Box
	want  string
}{
	{
		test: "Empty box",
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
	setup    func() *Box
	size     image.Point
	sizeHint image.Point
}{
	{
		test: "Stretch empty box",
		setup: func() *Box {
			b := NewHBox()
			b.SetBorder(true)
			return b
		},
		size:     image.Point{2, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "Stretch empty box",
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
	size  image.Point
	setup func() *Box
	want  string
}{
	{
		test: "Empty box",
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
		setup: func() *Box {
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
	{
		test: "Nested boxes expands equally",
		size: image.Point{32, 5},
		setup: func() *Box {
			nested := NewVBox(
				NewLabel("test"),
			)
			nested.SetBorder(true)
			nested.SetSizePolicy(Expanding, Minimum)

			b := NewHBox(
				NewSpacer(),
				nested,
				NewSpacer(),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)
			return b
		},
		want: `┌──────────────────────────────┐
│..........┌────────┐..........│
│..........│test....│..........│
│..........└────────┘..........│
└──────────────────────────────┘
`,
	},
	{
		test: "Nested boxes expands equally",
		size: image.Point{32, 5},
		setup: func() *Box {
			nested := NewVBox()
			nested.SetBorder(true)
			nested.SetSizePolicy(Expanding, Expanding)

			b := NewHBox(
				NewSpacer(),
				nested,
				NewSpacer(),
			)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)
			return b
		},
		want: `┌──────────────────────────────┐
│..........┌────────┐..........│
│..........│........│..........│
│..........└────────┘..........│
└──────────────────────────────┘
`,
	},
}

func TestHBox_Draw(t *testing.T) {
	for _, tt := range drawHBoxTests {
		var surface *testSurface
		if tt.size.X == 0 && tt.size.Y == 0 {
			surface = newTestSurface(10, 5)
		} else {
			surface = newTestSurface(tt.size.X, tt.size.Y)
		}
		painter := NewPainter(surface, NewPalette())

		b := tt.setup()

		b.Resize(surface.size)
		b.Draw(painter)

		if surface.String() != tt.want {
			t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
		}
	}
}
