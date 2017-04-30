package tui

import (
	"image"
	"strings"

	wordwrap "github.com/mitchellh/go-wordwrap"
)

var _ Widget = &TextEdit{}

// TextEdit is a multi-line text editor.
type TextEdit struct {
	WidgetBase

	text string

	onTextChange func(*TextEdit)
}

// TextEdit returns a new TextEdit.
func NewTextEdit() *TextEdit {
	return &TextEdit{}
}

// Draw draws the entry.
func (e *TextEdit) Draw(p *Painter) {
	s := e.Size()

	style := "entry"
	if e.IsFocused() {
		style += ".focused"
	}

	p.WithStyle(style, func(p *Painter) {
		lines := strings.Split(wordwrap.WrapString(e.text, uint(s.X)), "\n")
		for i, line := range lines {
			p.FillRect(0, i, s.X, 1)
			p.DrawText(0, i, line)
		}

		if e.IsFocused() {
			p.DrawCursor(stringWidth(lines[len(lines)-1]), len(lines)-1)
		}

		return
	})
}

// SizeHint returns the recommended size for the entry.
func (e *TextEdit) SizeHint() image.Point {
	p := image.Point{10, 1}
	if e.Size().X > p.X {
		p.X = e.Size().X
	}

	return image.Point{p.X, e.heightForWidth(p.X)}
}

// OnEvent handles terminal events.
func (e *TextEdit) OnEvent(ev Event) {
	if !e.IsFocused() {
		return
	}

	if ev.Type != EventKey {
		return
	}

	if ev.Key != 0 {
		switch ev.Key {
		case KeyEnter:
			e.text = e.text + "\n"
			return
		case KeySpace:
			e.text = e.text + string(' ')
			if e.onTextChange != nil {
				e.onTextChange(e)
			}
			return
		case KeyBackspace2:
			if len(e.text) > 0 {
				e.text = trimRightLen(e.text, 1)
				if e.onTextChange != nil {
					e.onTextChange(e)
				}
			}
			return
		}
	} else {
		e.text = e.text + string(ev.Ch)
		if e.onTextChange != nil {
			e.onTextChange(e)
		}
	}
}

// OnTextChanged sets a function to be run whenever the text content of the
// widget has been changed.
func (e *TextEdit) OnTextChanged(fn func(entry *TextEdit)) {
	e.onTextChange = fn
}

// SetText sets the text content of the entry.
func (e *TextEdit) SetText(text string) {
	e.text = text
}

// Text returns the text content of the entry.
func (e *TextEdit) Text() string {
	return e.text
}

func (e *TextEdit) heightForWidth(w int) int {
	return len(strings.Split(wordwrap.WrapString(e.text, uint(w)), "\n"))
}
