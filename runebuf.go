package tui

import (
	"image"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	wordwrap "github.com/mitchellh/go-wordwrap"
)

// RuneBuffer provides readline functionality for text widgets.
type RuneBuffer struct {
	buf []rune
	idx int

	wordwrap bool
}

// Width returns the width of the rune buffer, taking into account for CJK.
func (r *RuneBuffer) Width() int {
	return runewidth.StringWidth(string(r.buf))
}

// Set the buffer and the index at the end of the buffer.
func (r *RuneBuffer) Set(buf []rune) {
	r.SetWithIdx(len(buf), buf)
}

// SetWithIdx set the the buffer with a given index.
func (r *RuneBuffer) SetWithIdx(idx int, buf []rune) {
	r.buf = buf
	r.idx = idx
}

// WriteRune appends a rune to the buffer.
func (r *RuneBuffer) WriteRune(s rune) {
	r.WriteRunes([]rune{s})
}

// WriteRunes appends runes to the buffer.
func (r *RuneBuffer) WriteRunes(s []rune) {
	tail := append(s, r.buf[r.idx:]...)
	r.buf = append(r.buf[:r.idx], tail...)
	r.idx += len(s)
}

// Pos returns the current index in the buffer.
func (r *RuneBuffer) Pos() int {
	return r.idx
}

// Len returns the number of runes in the buffer.
func (r *RuneBuffer) Len() int {
	return len(r.buf)
}

// SplitByLine returns the lines for a given width.
func (r *RuneBuffer) SplitByLine(width int) []string {
	var text string
	if r.wordwrap {
		text = wordwrap.WrapString(r.String(), uint(width))
	} else {
		text = r.String()
	}
	return strings.Split(text, "\n")
}

func getSplitByLine(rs []rune, width int, wrap bool) []string {
	var text string
	if wrap {
		text = wordwrap.WrapString(string(rs), uint(width))
	} else {
		text = string(rs)
	}
	return strings.Split(text, "\n")
}

// CursorPos returns the coordinate for the cursor for a given width.
func (r *RuneBuffer) CursorPos(width int) image.Point {
	if width == 0 {
		return image.ZP
	}

	sp := getSplitByLine(r.buf[:r.idx], width, r.wordwrap)

	return image.Pt(stringWidth(sp[len(sp)-1]), len(sp)-1)
}

func (r *RuneBuffer) String() string {
	return string(r.buf)
}

// MoveBackward moves the cursor back by one rune.
func (r *RuneBuffer) MoveBackward() {
	if r.idx == 0 {
		return
	}
	r.idx--
}

// MoveForward moves the cursor forward by one rune.
func (r *RuneBuffer) MoveForward() {
	if r.idx == len(r.buf) {
		return
	}
	r.idx++
}

// MoveToLineStart moves the cursor to the start of the current line.
func (r *RuneBuffer) MoveToLineStart() {
	for i := r.idx; i > 0; i-- {
		if r.buf[i-1] == '\n' {
			r.idx = i
			return
		}
	}
	r.idx = 0
}

// MoveToLineEnd moves the cursor to the end of the current line.
func (r *RuneBuffer) MoveToLineEnd() {
	for i := r.idx; i < len(r.buf)-1; i++ {
		if r.buf[i+1] == '\n' {
			r.idx = i
			return
		}
	}
	r.idx = len(r.buf)
}

// Backspace deletes the rune left of the cursor.
func (r *RuneBuffer) Backspace() {
	if r.idx == 0 {
		return
	}
	r.idx--
	r.buf = append(r.buf[:r.idx], r.buf[r.idx+1:]...)
}

// Delete deletes the rune at the current cursor position.
func (r *RuneBuffer) Delete() {
	if r.idx == len(r.buf) {
		return
	}
	r.buf = append(r.buf[:r.idx], r.buf[r.idx+1:]...)
}

// Kill deletes all runes from the cursor until the end of the line.
func (r *RuneBuffer) Kill() {
	newlineIdx := strings.IndexRune(string(r.buf[r.idx:]), '\n')
	if newlineIdx < 0 {
		r.buf = r.buf[:r.idx]
	} else {
		r.buf = append(r.buf[:r.idx], r.buf[r.idx+newlineIdx+1:]...)
	}
}

func (r *RuneBuffer) heightForWidth(w int) int {
	return len(r.SplitByLine(w))
}
