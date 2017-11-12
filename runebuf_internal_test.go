package tui

import (
	"image"
	"reflect"
	"testing"
)

func TestRuneBuffer_MoveToLineStart(t *testing.T) {
	for _, tt := range []struct {
		curr RuneBuffer
		want RuneBuffer
	}{
		{RuneBuffer{idx: 3, buf: []rune("foo")}, RuneBuffer{idx: 0, buf: []rune("foo")}},
		{RuneBuffer{idx: 0, buf: []rune("foo")}, RuneBuffer{idx: 0, buf: []rune("foo")}},
	} {
		t.Run("", func(t *testing.T) {
			tt.curr.MoveToLineStart()

			if tt.want.idx != tt.curr.idx {
				t.Fatalf("want = %v; got = %v", tt.want.idx, tt.curr.idx)
			}
			if !reflect.DeepEqual(tt.want.buf, tt.curr.buf) {
				t.Fatalf("want = %v; got = %v", tt.want.buf, tt.curr.buf)
			}
		})
	}
}

func TestRuneBuffer_MoveToLineEnd(t *testing.T) {
	for _, tt := range []struct {
		curr RuneBuffer
		want RuneBuffer
	}{
		{RuneBuffer{idx: 3, buf: []rune("foo")}, RuneBuffer{idx: 3, buf: []rune("foo")}},
		{RuneBuffer{idx: 0, buf: []rune("foo")}, RuneBuffer{idx: 3, buf: []rune("foo")}},
	} {
		t.Run("", func(t *testing.T) {
			tt.curr.MoveToLineEnd()

			if tt.want.idx != tt.curr.idx {
				t.Fatalf("want = %v; got = %v", tt.want.idx, tt.curr.idx)
			}
			if !reflect.DeepEqual(tt.want.buf, tt.curr.buf) {
				t.Fatalf("want = %v; got = %v", tt.want.buf, tt.curr.buf)
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

func TestRuneBuffer_CursorPos(t *testing.T) {
	for _, tt := range []struct {
		text        string
		screenWidth int
		idx         int
		out         image.Point
	}{
		{"Lorem ipsum dolor sit amet.", 12, 27, image.Pt(5, 2)},
		{"Lorem ipsum dolor sit amet.", 16, 27, image.Pt(15, 1)},
		{"Lorem ipsum dolor sit amet.", 27, 20, image.Pt(20, 0)},
	} {
		t.Run("", func(t *testing.T) {
			var r RuneBuffer
			r.wordwrap = true
			r.SetWithIdx(tt.idx, []rune(tt.text))

			if got := r.CursorPos(tt.screenWidth); tt.out != got {
				t.Fatalf("want = %s; got = %s", tt.out, got)
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
		{"Lorem ipsum dolor sit amet.", 12, true, []string{"Lorem ipsum", "dolor sit", "amet."}},
		{"Lorem ipsum dolor sit amet.", 27, true, []string{"Lorem ipsum dolor sit amet."}},
		{"Lorem ipsum dolor sit amet.", 12, false, []string{"Lorem ipsum dolor sit amet."}},
	} {
		got := getSplitByLine([]rune(tt.text), tt.width, tt.wrap)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("want = %#v; got = %#v", tt.want, got)
		}
	}
}

func TestRuneBuffer_CursorPosWithWordWrap(t *testing.T) {
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
		{"Lorem ipsum dolor sit amet.", 12, 12, true, image.Pt(0, 1)},
		{"Lorem ipsum dolor sit amet.", 12, 13, true, image.Pt(1, 1)},

		// Lorem ipsum dolor
		// sit amet.
		{"Lorem ipsum dolor sit amet.", 19, 17, true, image.Pt(17, 0)},
		{"Lorem ipsum dolor sit amet.", 19, 18, true, image.Pt(0, 1)},
		{"Lorem ipsum dolor sit amet.", 19, 19, true, image.Pt(1, 1)},
		{"Lorem ipsum dolor sit amet.", 19, 20, true, image.Pt(2, 1)},
		{"Lorem ipsum dolor sit amet.", 19, 21, true, image.Pt(3, 1)},
	} {
		t.Skip("Skip until a more sophisticated word wrap has been implemented.")
		t.Run("", func(t *testing.T) {
			var r RuneBuffer
			r.wordwrap = tt.wrap
			r.SetWithIdx(tt.idx, []rune(tt.text))

			if got := r.CursorPos(tt.screenWidth); tt.want != got {
				t.Fatalf("want = %s; got = %s", tt.want, got)
			}
		})
	}
}
