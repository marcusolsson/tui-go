package tui

import (
	"image"
	"testing"
)

var verticalBoxSizeTests = []struct {
	test     string
	setup    func() *VerticalBox
	size     image.Point
	sizeHint image.Point
}{
	{
		test: "Stretch empty box",
		setup: func() *VerticalBox {
			b := NewVerticalBox()
			b.SetBorder(true)
			return b
		},
		size:     image.Point{2, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "Stretch empty box",
		setup: func() *VerticalBox {
			b := NewVerticalBox()
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Minimum)
			return b
		},
		size:     image.Point{100, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "Stretch empty box",
		setup: func() *VerticalBox {
			b := NewVerticalBox()
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)
			return b
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "No stretch",
		setup: func() *VerticalBox {
			b := NewVerticalBox(
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
		setup: func() *VerticalBox {
			b := NewVerticalBox(
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
		setup: func() *VerticalBox {
			b := NewVerticalBox(
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
	{
		test: "With padding",
		setup: func() *VerticalBox {
			b := NewVerticalBox(
				NewLabel("test"),
				NewLabel("another test"),
			)
			b.SetBorder(true)
			b.SetPadding(2)

			return b
		},
		size:     image.Point{14, 6},
		sizeHint: image.Point{14, 6},
	},
	{
		test: "With margin",
		setup: func() *VerticalBox {
			b := NewVerticalBox(
				NewLabel("test"),
				NewLabel("another test"),
			)
			b.SetBorder(true)
			b.SetMargin(5)

			return b
		},
		size:     image.Point{24, 14},
		sizeHint: image.Point{24, 14},
	},
}

func TestVerticalBoxSize(t *testing.T) {
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

var horizontalBoxSizeTests = []struct {
	test     string
	setup    func() *HorizontalBox
	size     image.Point
	sizeHint image.Point
}{
	{
		test: "Stretch empty box",
		setup: func() *HorizontalBox {
			b := NewHorizontalBox()
			b.SetBorder(true)
			return b
		},
		size:     image.Point{2, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "Stretch empty box",
		setup: func() *HorizontalBox {
			b := NewHorizontalBox()
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Minimum)
			return b
		},
		size:     image.Point{100, 2},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "Stretch empty box",
		setup: func() *HorizontalBox {
			b := NewHorizontalBox()
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)
			return b
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{2, 2},
	},
	{
		test: "No stretch",
		setup: func() *HorizontalBox {
			b := NewHorizontalBox(
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
		setup: func() *HorizontalBox {
			b := NewHorizontalBox(
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
		setup: func() *HorizontalBox {
			b := NewHorizontalBox(
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
		setup: func() *HorizontalBox {
			nested := NewHorizontalBox(NewLabel("test"))
			nested.SetBorder(true)
			nested.SetSizePolicy(Expanding, Minimum)

			b := NewHorizontalBox(nested)
			b.SetBorder(true)
			b.SetSizePolicy(Expanding, Expanding)

			return b
		},
		size:     image.Point{100, 100},
		sizeHint: image.Point{8, 5},
	},
}

func TestHorizontalBoxSize(t *testing.T) {
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
