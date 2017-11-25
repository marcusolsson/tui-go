package tui_test

import (
	"image"
	"testing"
	"github.com/marcusolsson/tui-go"
)

var drawTableTests = []struct {
	test  string
	size  image.Point
	setup func() *tui.Box
	want  string
}{
	{
		test: "Long labels are masked (#31)",
		size: image.Point{32, 10},
		setup: func() *tui.Box {
			first := tui.NewTable(0, 0)
			first.AppendRow(tui.NewLabel("ABC123"), tui.NewLabel("test"))
			first.AppendRow(tui.NewLabel("DEF456"), tui.NewLabel("testing a longer text"))
			first.AppendRow(tui.NewLabel("GHI789"), tui.NewLabel("foo"))
			first.SetBorder(true)

			second := tui.NewVBox(tui.NewLabel("test"))
			second.SetBorder(true)

			third := tui.NewHBox(first, second)
			third.SetBorder(true)

			return third
		},
		want: `
┌──────────────────────────────┐
│┌───────────┬──────────┐┌────┐│
││ABC123     │test      ││test││
││           │          ││....││
│├───────────┼──────────┤│....││
││DEF456     │testing a ││....││
│├───────────┼──────────┤│....││
││GHI789     │foo       ││....││
│└───────────┴──────────┘└────┘│
└──────────────────────────────┘
`,
	},
	{
		test: "Remove a row from table",
		size: image.Point{20, 10},
		setup: func() *tui.Box {
			table := tui.NewTable(0, 0)
			table.AppendRow(tui.NewLabel("A"), tui.NewLabel("apple"))
			table.AppendRow(tui.NewLabel("B"), tui.NewLabel("box"))
			table.AppendRow(tui.NewLabel("C"), tui.NewLabel("cat"))
			table.AppendRow(tui.NewLabel("D"), tui.NewLabel("dog"))
			table.SetBorder(true)

			table.RemoveRow(1)

			box := tui.NewHBox(table)
			box.SetBorder(true)

			return box
		},
		want: `
┌──────────────────┐
│┌────────┬───────┐│
││A       │apple  ││
││        │       ││
│├────────┼───────┤│
││C       │cat    ││
│├────────┼───────┤│
││D       │dog    ││
│└────────┴───────┘│
└──────────────────┘
`,
	},
	{
		test: "Remove all rows from table",
		size: image.Point{20, 10},
		setup: func() *tui.Box {
			table := tui.NewTable(0, 0)
			table.AppendRow(tui.NewLabel("A"), tui.NewLabel("apple"))
			table.AppendRow(tui.NewLabel("B"), tui.NewLabel("box"))
			table.AppendRow(tui.NewLabel("C"), tui.NewLabel("cat"))
			table.AppendRow(tui.NewLabel("D"), tui.NewLabel("dog"))
			table.SetBorder(true)

			table.RemoveRows()

			box := tui.NewHBox(table)
			box.SetBorder(true)

			return box
		},
		want: `
┌──────────────────┐
│┌────────┬───────┐│
││........│.......││
││........│.......││
││........│.......││
││........│.......││
││........│.......││
││........│.......││
│└────────┴───────┘│
└──────────────────┘
`,
	},
}

func TestTable_Draw(t *testing.T) {
	for _, tt := range drawTableTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			var surface *testSurface
			if tt.size.X == 0 && tt.size.Y == 0 {
				surface = newTestSurface(10, 5)
			} else {
				surface = newTestSurface(tt.size.X, tt.size.Y)
			}

			painter := tui.NewPainter(surface, tui.NewTheme())
			painter.Repaint(tt.setup())

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}
