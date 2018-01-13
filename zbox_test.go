package tui

import (
	"fmt"
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
	{
		test: "AppendRemoveMiddle",
		setup: func() Widget {
			popupSpace := NewZBox(
				NewLabel("bottoms"),
			)
			midRemove := popupSpace.Append(NewLabel("middle"))
			_ = popupSpace.Append(NewLabel("top"))

			midRemove()
			return popupSpace
		},
		want: `
toptoms...
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

func ExampleZBox_popup() {
	s := NewTestSurface(15, 7)
	painter := NewPainter(s, NewTheme())

	contributors := NewVBox()
	for _, contributor := range []string{
		"marcusolsson",
		"cceckman",
		"raghuvanshy",
		"jsageryd",
		"fenimore",
	} {
		contributors.Append(NewLabel(contributor))
	}
	contributors.Append(NewSpacer())
	root := NewZBox(contributors)

	painter.Repaint(root)
	fmt.Print(s.String())
	center := func(w Widget) Widget {
		return NewVBox(NewSpacer(), NewHBox(NewSpacer(), w, NewSpacer()), NewSpacer())
	}
	thanks := NewLabel("Thanks for using tui-go!")
	thanks.SetWordWrap(true)
	thanksBox := NewVBox(thanks)
	thanksBox.SetBorder(true)
	thanksBox.SetFill(true)

	closePop := root.Append(center(thanksBox))

	painter.Repaint(root)
	// Repaint twice to get word-wrap behavior right: marcusolsson/tui-go#108
	painter.Repaint(root)
	fmt.Print(s.String())

	closePop()
	painter.Repaint(root)
	fmt.Printf(s.String())

	// Output:
	// marcusolsson...
	// cceckman.......
	// raghuvanshy....
	// jsageryd.......
	// fenimore.......
	// ...............
	// ...............
	//
	// marcusolsson...
	// cceckman.......
	// ┌─────────────┐
	// │Thanks for   │
	// │using tui-go!│
	// └─────────────┘
	// ...............
	//
	// marcusolsson...
	// cceckman.......
	// raghuvanshy....
	// jsageryd.......
	// fenimore.......
	// ...............
	// ...............
}

func ExampleZBox_popupOnPopup() {
	s := NewTestSurface(15, 7)
	painter := NewPainter(s, NewTheme())

	contributors := NewVBox()
	for _, contributor := range []string{
		"marcusolsson",
		"cceckman",
		"raghuvanshy",
		"jsageryd",
		"fenimore",
	} {
		contributors.Append(NewLabel(contributor))
	}
	contributors.Append(NewSpacer())
	root := NewZBox(contributors)

	painter.Repaint(root)
	fmt.Print(s.String())
	center := func(w Widget) Widget {
		return NewVBox(NewSpacer(), NewHBox(NewSpacer(), w, NewSpacer()), NewSpacer())
	}
	thanks := NewLabel("Thanks for using tui-go!")
	thanks.SetWordWrap(true)
	thanksBox := NewVBox(thanks)
	thanksBox.SetBorder(true)
	thanksBox.SetFill(true)

	closePop := root.Append(center(thanksBox))

	painter.Repaint(root)
	// Repaint twice to get word-wrap behavior right: marcusolsson/tui-go#108
	painter.Repaint(root)
	fmt.Print(s.String())

	smile := NewVBox(NewLabel("😀"))
	smile.SetBorder(true)
	// ┌──┐
	// │😀│
	// └──┘
	_ = root.Append(center(smile))
	painter.Repaint(root)
	fmt.Printf(s.String())


	closePop()
	painter.Repaint(root)
	fmt.Printf(s.String())

	// ┌──┐
	// │😀│
	// └──┘


	// Output:
	// marcusolsson...
	// cceckman.......
	// raghuvanshy....
	// jsageryd.......
	// fenimore.......
	// ...............
	// ...............
	//
	// marcusolsson...
	// cceckman.......
	// ┌─────┌──┐────┐
	// │Thank│😀│r   │
	// │using└──┘-go!│
	// └─────────────┘
	// ...............
	//

	// marcusolsson...
	// cceckman.......
	// raghuv┌──┐y...
	// jsager│😀│....
	// fenimo└──┘....
	// ...............
	// ...............
}
