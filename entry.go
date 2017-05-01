package tui

import (
	"image"
	"strings"

	wordwrap "github.com/mitchellh/go-wordwrap"
)

var _ Widget = &Entry{}

// Entry is a one-line text editor. It lets the user supply your application
// with text, e.g. to input user and password information.
type Entry struct {
	WidgetBase

	text string

	onTextChange func(*Entry)
	onSubmit     func(*Entry)
}

// NewEntry returns a new Entry.
func NewEntry() *Entry {
	return &Entry{}
}

// Draw draws the entry.
func (e *Entry) Draw(p *Painter) {
	style := "entry"
	if e.IsFocused() {
		style += ".focused"
	}
	p.WithStyle(style, func(p *Painter) {
		s := e.Size()

		tw := stringWidth(e.text)
		offx := tw - s.X

		// Make room for cursor.
		if e.IsFocused() {
			offx++
		}

		text := e.text
		if tw >= s.X {
			text = text[offx:]
		}

		p.FillRect(0, 0, s.X, 1)
		p.DrawText(0, 0, text)

		if e.IsFocused() {
			p.DrawCursor(stringWidth(text), 0)
		}
	})
}

// SizeHint returns the recommended size for the entry.
func (e *Entry) SizeHint() image.Point {
	return image.Point{10, 1}
}

// OnEvent handles terminal events.
func (e *Entry) OnEvent(ev Event) {
	if !e.IsFocused() || ev.Type != EventKey {
		return
	}

	if ev.Key != 0 {
		switch ev.Key {
		case KeyEnter:
			if e.onSubmit != nil {
				e.onSubmit(e)
			}
		case KeySpace:
			e.text = e.text + string(' ')
			if e.onTextChange != nil {
				e.onTextChange(e)
			}
		case KeyBackspace2:
			if len(e.text) > 0 {
				e.text = trimRightLen(e.text, 1)
				if e.onTextChange != nil {
					e.onTextChange(e)
				}
			}
		}
		return
	}

	e.text = e.text + string(ev.Ch)
	if e.onTextChange != nil {
		e.onTextChange(e)
	}
}

// OnChanged sets a function to be run whenever the content of the entry has
// been changed.
func (e *Entry) OnChanged(fn func(entry *Entry)) {
	e.onTextChange = fn
}

// OnSubmit sets a function to be run whenever the user submits the entry (by
// pressing KeyEnter).
func (e *Entry) OnSubmit(fn func(entry *Entry)) {
	e.onSubmit = fn
}

// SetText sets the text content of the entry.
func (e *Entry) SetText(text string) {
	e.text = text
}

// Text returns the text content of the entry.
func (e *Entry) Text() string {
	return e.text
}

func (e *Entry) heightForWidth(w int) int {
	return len(strings.Split(wordwrap.WrapString(e.text, uint(w)), "\n"))
}
