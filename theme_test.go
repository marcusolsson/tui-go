package tui

import (
	"testing"
)

var base = Style{Fg: ColorWhite, Bg: ColorBlack, Bold: DecorationOff, Underline: DecorationOff}

var mergeTests = []struct {
	test  string
	chain []Style
	want  Style
}{
	{
		test: "zero inherits",
		chain: []Style{
			base,
			Style{},
		},
		want: base,
	},
	{
		test: "orthogonal inherits",
		chain: []Style{
			base,
			Style{Reverse: DecorationOn},
			Style{Bold: DecorationOn},
		},
		want: Style{Fg: ColorWhite, Bg: ColorBlack, Reverse: DecorationOn, Bold: DecorationOn, Underline: DecorationOff},
	},
	{
		test: "serial inherits",
		chain: []Style{
			base,
			Style{Reverse: DecorationOn, Bold: DecorationOn},
			Style{Bold: DecorationOff},
		},
		want: Style{Fg: ColorWhite, Bg: ColorBlack, Reverse: DecorationOn, Bold: DecorationOff, Underline: DecorationOff},
	},
}

func TestStyle_Merge(t *testing.T) {
	for _, tt := range mergeTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			var got Style
			for _, s := range tt.chain {
				got = got.mergeIn(s)
			}
			if got != tt.want {
				t.Errorf("got = \n%v\nwant = \n%v", got, tt.want)
			}
		})
	}
}
