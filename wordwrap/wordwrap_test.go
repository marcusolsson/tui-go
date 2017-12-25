package wordwrap

import (
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	for _, tt := range []struct {
		In    string
		Width int
		Out   string
	}{
		{"", 3, ""},
		{"a", 3, "a"},
		{"aa", 3, "aa"},
		{"aa bb", 3, "aa \nbb"},
		{"aa bb ddd", 7, "aa bb \nddd"},
		{"aaa bb cc ddddd", 7, "aaa bb \ncc \nddddd"},
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis egestas nisl urna, vel accumsan libero bibendum id. In consectetur facilisis iaculis. Vivamus cursus hendrerit neque, et bibendum leo accumsan ac.", 20, "Lorem ipsum dolor \nsit amet, \nconsectetur \nadipiscing elit. \nDuis egestas nisl \nurna, vel accumsan \nlibero bibendum id. \nIn consectetur \nfacilisis iaculis. \nVivamus cursus \nhendrerit neque, et \nbibendum leo \naccumsan ac."},
		{"aaa bb\n\ncc ddddd", 7, "aaa bb\n\ncc \nddddd"},
		{"Nulla lorem magna, efficitur interdum ante at, convallis sodales nulla. Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, sit amet fringilla nisl pharetra quis.", 30, "Nulla lorem magna, efficitur \ninterdum ante at, convallis \nsodales nulla. Sed maximus \ntempor condimentum.\n\nNam et risus est. Cras \nornare iaculis orci, sit amet \nfringilla nisl pharetra quis."},
		{"Nulla lorem magna, efficitur interdum ante at, convallis sodales nulla. Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, sit amet fringilla nisl pharetra quis.", 35, "Nulla lorem magna, efficitur \ninterdum ante at, convallis \nsodales nulla. Sed maximus tempor \ncondimentum.\n\nNam et risus est. Cras ornare \niaculis orci, sit amet fringilla \nnisl pharetra quis."},
		{"\n\nNam et risus est.", 30, "\n\nNam et risus est."},
		{"a\n\na\n\n", 6, "a\n\na\n\n"},
		{"null set ∅", 11, "null set ∅"},
	} {
		t.Run("", func(t *testing.T) {
			if got := WrapString(tt.In, tt.Width); got != tt.Out {
				padding := strings.Repeat(".", tt.Width)
				t.Fatalf("\n\ngot = \n\n%s\n%s\n%s\n\nwant = \n\n%s\n%s\n%s", padding, got, padding, padding, tt.Out, padding)
			}
		})
	}
}
