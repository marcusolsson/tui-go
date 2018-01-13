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
	{
		test: "CoopOvewrite",
		setup: func() Widget {
			bgFill := NewVBox(
				NewLabel("background"),
				NewSpacer(),
			)
			popup := NewVBox(NewLabel("popup"))
			popup.SetBorder(true)

			fg := NewVBox(
				NewSpacer(),
				NewHBox(
					NewSpacer(),
					popup,
					NewSpacer(),
				),
				NewSpacer(),
			)

			return NewZBox(bgFill, fg)
		},
		want: `
background
..┌─────┐.
..│popup│.
..└─────┘.
..........
`,
	},
	{
		test: "PartOfScreen",
		setup: func() Widget {
			return NewVBox(
				NewLabel("tops"),
				NewSpacer(),
				NewZBox(
					NewLabel("bottoms"),
				),
			)
		},
		want: `
tops......
..........
..........
bottoms...
..........
`,
	},
	{
		test: "Empty",
		setup: func() Widget {
			return NewVBox(
				NewLabel("tops"),
				NewSpacer(),
				NewZBox(),
				NewLabel("bottoms"),
			)
		},
		want: `
tops......
..........
..........
bottoms...
..........
`,
	},

	{
		test: "PartialPopover",
		setup: func() Widget {
			pop := NewVBox(NewLabel("popup"))
			pop.SetFill(true)
			pop.SetBorder(true)
			base := NewVBox(
				NewLabel("tops"),
				NewSpacer(),
				NewZBox(
					NewLabel("bottoms"),
					pop,
				),
			)
			return base
		},
		want: `
tops......
..........
┌────────┐
│popup   │
└────────┘
`,
	},
	{
		test: "Append",
		setup: func() Widget {
			popupSpace := NewZBox(
				NewLabel("bottoms"),
			)

			base := NewVBox(
				NewLabel("tops"),
				NewSpacer(),
				popupSpace,
			)

			pop := NewVBox(NewLabel("popup"))
			pop.SetFill(true)
			pop.SetBorder(true)

			_ = popupSpace.Append(pop)

			return base
		},
		want: `
tops......
..........
┌────────┐
│popup   │
└────────┘
`,
	},
	{
		test: "AppendAndRemove",
		setup: func() Widget {
			popupSpace := NewZBox(
				NewLabel("bottoms"),
			)

			base := NewVBox(
				NewLabel("tops"),
				NewSpacer(),
				popupSpace,
			)

			pop := NewVBox(NewLabel("popup"))
			pop.SetFill(true)
			pop.SetBorder(true)

			remove := popupSpace.Append(pop)
			remove()

			return base
		},
		want: `
tops......
..........
..........
bottoms...
..........
`,
	},
	{
		test: "AppendAndRemove",
		setup: func() Widget {
			popupSpace := NewZBox(
				NewLabel("bottoms"),
			)

			base := NewVBox(
				NewLabel("tops"),
				NewSpacer(),
				popupSpace,
			)

			pop := NewVBox(NewLabel("popup"))
			pop.SetFill(true)
			pop.SetBorder(true)

			remove := popupSpace.Append(pop)
			remove()

			return base
		},
		want: `
tops......
..........
..........
bottoms...
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
