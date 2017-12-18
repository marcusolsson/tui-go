package tui

import "testing"

var drawCJKTests = []struct {
	test  string
	setup func() Widget
	want  string
}{
	{
		test: "Label",
		setup: func() Widget {
			return NewLabel("テスト")
		},
		want: `
テスト....
..........
..........
..........
`,
	},
	{
		test: "Box",
		setup: func() Widget {
			b := NewVBox(
				NewLabel("テスト１"),
				NewLabel("テスト２"),
			)
			b.SetBorder(true)
			return b
		},
		want: `
┌────────┐
│テスト１│
│テスト２│
└────────┘
`,
	},
	{
		test: "Box with title",
		setup: func() Widget {
			b := NewVBox(
				NewLabel("测试"),
			)
			b.SetTitle("标题")
			b.SetBorder(true)
			return b
		},
		want: `
┌标题────┐
│测试    │
│        │
└────────┘
`,
	},
	{
		test: "Entry",
		setup: func() Widget {
			e := NewEntry()
			e.SetText("テスト")
			b := NewVBox(e)
			b.SetBorder(true)
			return b
		},
		want: `
┌────────┐
│テスト  │
│        │
└────────┘
`,
	},
	{
		test: "List",
		setup: func() Widget {
			l := NewList()
			l.AddItems("テスト１", "テスト２")
			b := NewVBox(l)
			b.SetBorder(true)
			return b
		},
		want: `
┌────────┐
│テスト１│
│テスト２│
└────────┘
`,
	},
}

func TestCJK_Label(t *testing.T) {
	for _, tt := range drawCJKTests {
		surface := NewTestSurface(10, 4)

		painter := NewPainter(surface, NewTheme())
		painter.Repaint(tt.setup())

		if surface.String() != tt.want {
			t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
		}
	}
}
