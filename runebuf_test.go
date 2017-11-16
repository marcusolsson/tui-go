package tui

import (
	"image"
	"reflect"
	"testing"
)

func TestRuneBuffer_MoveForward(t *testing.T) {
	for _, tt := range []struct {
		text    string
		in, out int
	}{
		{"foo", 0, 1},
		{"foo", 2, 3},
		{"foo", 3, 3},
		{"Lorem ipsum dolor \nsit amet.", 17, 18},
	} {
		t.Run("", func(t *testing.T) {
			var buf RuneBuffer
			buf.SetWithIdx(tt.in, []rune(tt.text))

			buf.MoveForward()

			if tt.out != buf.idx {
				t.Fatalf("want = %v; got = %v", tt.out, buf.idx)
			}
		})
	}
}

func TestRuneBuffer_MoveBackward(t *testing.T) {
	for _, tt := range []struct {
		text    string
		in, out int
	}{
		{"foo", 0, 0},
		{"foo", 2, 1},
		{"foo", 3, 2},
		{"Lorem ipsum dolor \nsit amet.", 18, 17},
	} {
		t.Run("", func(t *testing.T) {
			var buf RuneBuffer
			buf.SetWithIdx(tt.in, []rune(tt.text))

			buf.MoveBackward()

			if tt.out != buf.idx {
				t.Fatalf("want = %v; got = %v", tt.out, buf.idx)
			}
		})
	}
}

func TestRuneBuffer_MoveToLineStart(t *testing.T) {
	for _, tt := range []struct {
		text    string
		in, out int
	}{
		{"foo", 3, 0},
		{"foo", 0, 0},
		{"Lorem ipsum dolor \nsit amet.", 21, 19},
		{"Lorem ipsum dolor \n\nsit amet.", 21, 20},
		{"Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, \n\nsit amet fringilla nisl pharetra quis.", 40, 33},
		{"Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, \n\nsit amet fringilla nisl pharetra quis.", 90, 79},
		{"Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, \n\nsit amet fringilla nisl pharetra quis.", 79, 79},
		// On a empty line.
		{"Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, \n\nsit amet fringilla nisl pharetra quis.", 78, 78},
		// On newline character.
		{"Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, \n\nsit amet fringilla nisl pharetra quis.", 77, 33},
	} {
		t.Run("", func(t *testing.T) {
			var buf RuneBuffer
			buf.SetWithIdx(tt.in, []rune(tt.text))

			buf.MoveToLineStart()

			if tt.out != buf.idx {
				t.Fatalf("want = %v; got = %v", tt.out, buf.idx)
			}
		})
	}
}

func TestRuneBuffer_MoveToLineEnd(t *testing.T) {
	for _, tt := range []struct {
		text    string
		in, out int
	}{
		{"foo", 0, 3},
		{"foo", 3, 3},
		{"Lorem ipsum dolor \nsit amet.", 0, 18},
		{"Lorem ipsum dolor \n\nsit amet.", 20, 29},
		{"Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, \n\nsit amet fringilla nisl pharetra quis.", 0, 31},
		{"Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, \n\nsit amet fringilla nisl pharetra quis.", 33, 77},
		{"Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, \n\nsit amet fringilla nisl pharetra quis.", 79, 117},
		{"Sed maximus tempor condimentum.\n\nNam et risus est. Cras ornare iaculis orci, \n\nsit amet fringilla nisl pharetra quis.", 77, 77},
	} {
		t.Run("", func(t *testing.T) {
			var buf RuneBuffer
			buf.SetWithIdx(tt.in, []rune(tt.text))

			buf.MoveToLineEnd()

			if tt.out != buf.idx {
				t.Fatalf("want = %v; got = %v", tt.out, buf.idx)
			}
		})
	}
}

func TestRuneBuffer_Backspace(t *testing.T) {
	for _, tt := range []struct {
		curr RuneBuffer
		want RuneBuffer
	}{
		{RuneBuffer{idx: 0, buf: []rune("foo bar")}, RuneBuffer{idx: 0, buf: []rune("foo bar")}},
		{RuneBuffer{idx: 1, buf: []rune("foo bar")}, RuneBuffer{idx: 0, buf: []rune("oo bar")}},
		{RuneBuffer{idx: 7, buf: []rune("foo bar")}, RuneBuffer{idx: 6, buf: []rune("foo ba")}},
		{RuneBuffer{idx: 4, buf: []rune("foo bar")}, RuneBuffer{idx: 3, buf: []rune("foobar")}},
	} {
		t.Run("", func(t *testing.T) {
			tt.curr.Backspace()

			if tt.want.idx != tt.curr.idx {
				t.Fatalf("want = %v; got = %v", tt.want.idx, tt.curr.idx)
			}
			if !reflect.DeepEqual(tt.want.buf, tt.curr.buf) {
				t.Fatalf("want = %q; got = %q", string(tt.want.buf), string(tt.curr.buf))
			}
		})
	}
}

func TestRuneBuffer_Kill(t *testing.T) {
	for _, tt := range []struct {
		curr RuneBuffer
		want RuneBuffer
	}{
		{RuneBuffer{idx: 0, buf: []rune("foo bar")}, RuneBuffer{idx: 0, buf: []rune("")}},
		{RuneBuffer{idx: 1, buf: []rune("foo bar")}, RuneBuffer{idx: 1, buf: []rune("f")}},
		{RuneBuffer{idx: 7, buf: []rune("foo bar")}, RuneBuffer{idx: 7, buf: []rune("foo bar")}},
		{RuneBuffer{idx: 4, buf: []rune("foo bar")}, RuneBuffer{idx: 4, buf: []rune("foo ")}},
		{RuneBuffer{idx: 0, buf: []rune("foo \nbar")}, RuneBuffer{idx: 0, buf: []rune("bar")}},
		{RuneBuffer{idx: 5, buf: []rune("foo \nbar")}, RuneBuffer{idx: 5, buf: []rune("foo \n")}},
		{RuneBuffer{idx: 6, buf: []rune("foo \nbar")}, RuneBuffer{idx: 6, buf: []rune("foo \nb")}},
		{RuneBuffer{idx: 0, buf: []rune("\n")}, RuneBuffer{idx: 0, buf: []rune("")}},
	} {
		t.Run("", func(t *testing.T) {
			tt.curr.Kill()

			if tt.want.idx != tt.curr.idx {
				t.Fatalf("want = %v; got = %v", tt.want.idx, tt.curr.idx)
			}
			if !reflect.DeepEqual(tt.want.buf, tt.curr.buf) {
				t.Fatalf("want = %q; got = %q", string(tt.want.buf), string(tt.curr.buf))
			}
		})
	}
}

func TestRuneBuffer_SplitByLines(t *testing.T) {
	for _, tt := range []struct {
		text  string
		width int
		wrap  bool
		want  []string
	}{
		{"Lorem ipsum dolor sit amet.", 12, true, []string{"Lorem ipsum ", "dolor sit ", "amet."}},
		{"Lorem ipsum dolor sit amet.", 27, true, []string{"Lorem ipsum dolor sit amet."}},
		{"Lorem ipsum dolor sit amet.", 12, false, []string{"Lorem ipsum dolor sit amet."}},
	} {
		var buf RuneBuffer
		buf.Set([]rune(tt.text))
		buf.SetMaxWidth(tt.width)
		buf.wordwrap = tt.wrap

		got := buf.SplitByLine()
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want = %#v; got = %#v", tt.want, got)
		}
	}
}

func TestRuneBuffer_CursorPos(t *testing.T) {
	for _, tt := range []struct {
		text        string
		screenWidth int
		idx         int
		wrap        bool
		want        image.Point
	}{
		// Lorem ipsum
		// dolor sit amet.
		{"Lorem ipsum dolor sit amet.", 12, 11, true, image.Pt(11, 0)},
		{"Lorem ipsum dolor sit amet.", 12, 12, true, image.Pt(12, 0)},
		{"Lorem ipsum dolor sit amet.", 12, 13, true, image.Pt(0, 1)},

		// Lorem ipsum dolor
		// sit amet.
		{"Lorem ipsum dolor sit amet.", 19, 17, true, image.Pt(17, 0)},
		{"Lorem ipsum dolor sit amet.", 19, 18, true, image.Pt(18, 0)},
		{"Lorem ipsum dolor sit amet.", 19, 19, true, image.Pt(0, 1)},
		{"Lorem ipsum dolor sit amet.", 19, 20, true, image.Pt(1, 1)},
		{"Lorem ipsum dolor sit amet.", 19, 21, true, image.Pt(2, 1)},

		// aa bb
		//
		// cc dd
		{"aa bb\n\ncc dd", 10, 4, true, image.Pt(4, 0)},
		{"aa bb\n\ncc dd", 10, 5, true, image.Pt(5, 0)},
		{"aa bb\n\ncc dd", 10, 6, true, image.Pt(0, 1)},
		{"aa bb\n\ncc dd", 10, 7, true, image.Pt(0, 2)},
	} {
		t.Run("", func(t *testing.T) {
			var r RuneBuffer
			r.wordwrap = tt.wrap
			r.SetWithIdx(tt.idx, []rune(tt.text))
			r.SetMaxWidth(tt.screenWidth)

			if got := r.CursorPos(); tt.want != got {
				t.Fatalf("want = %s; got = %s", tt.want, got)
			}
		})
	}
}
