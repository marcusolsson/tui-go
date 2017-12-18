package tui

import (
	"image"
	"testing"
)

var drawTableTests = []struct {
	test  string
	size  image.Point
	setup func() *Box
	want  string
}{
	{
		test: "Long labels are masked (#31)",
		size: image.Point{32, 10},
		setup: func() *Box {
			first := NewTable(0, 0)
			first.AppendRow(NewLabel("ABC123"), NewLabel("test"))
			first.AppendRow(NewLabel("DEF456"), NewLabel("testing a longer text"))
			first.AppendRow(NewLabel("GHI789"), NewLabel("foo"))
			first.SetBorder(true)

			second := NewVBox(NewLabel("test"))
			second.SetBorder(true)

			third := NewHBox(first, second)
			third.SetBorder(true)

			return third
		},
		want: `
┌──────────────────────────────┐
│┌───────────┬──────────┐┌────┐│
││ABC123     │test      ││test││
││           │          ││    ││
│├───────────┼──────────┤│    ││
││DEF456     │testing a ││    ││
│├───────────┼──────────┤│    ││
││GHI789     │foo       ││    ││
│└───────────┴──────────┘└────┘│
└──────────────────────────────┘
`,
	},
	{
		test: "Remove a row from table",
		size: image.Point{20, 10},
		setup: func() *Box {
			table := NewTable(0, 0)
			table.AppendRow(NewLabel("A"), NewLabel("apple"))
			table.AppendRow(NewLabel("B"), NewLabel("box"))
			table.AppendRow(NewLabel("C"), NewLabel("cat"))
			table.AppendRow(NewLabel("D"), NewLabel("dog"))
			table.SetBorder(true)

			table.RemoveRow(1)

			box := NewHBox(table)
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
		setup: func() *Box {
			table := NewTable(0, 0)
			table.AppendRow(NewLabel("A"), NewLabel("apple"))
			table.AppendRow(NewLabel("B"), NewLabel("box"))
			table.AppendRow(NewLabel("C"), NewLabel("cat"))
			table.AppendRow(NewLabel("D"), NewLabel("dog"))
			table.SetBorder(true)

			table.RemoveRows()

			box := NewHBox(table)
			box.SetBorder(true)

			return box
		},
		want: `
┌──────────────────┐
│┌────────┬───────┐│
││        │       ││
││        │       ││
││        │       ││
││        │       ││
││        │       ││
││        │       ││
│└────────┴───────┘│
└──────────────────┘
`,
	},
}

func TestTable_Draw(t *testing.T) {
	for _, tt := range drawTableTests {
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
