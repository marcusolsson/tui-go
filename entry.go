package tui

import (
	"image"
)

var _ Widget = &Entry{}

// Entry is a one-line text editor. It lets the user supply the application
// with text, e.g., to input user and password information.
type Entry struct {
	WidgetBase

	text RuneBuffer

	onTextChange func(*Entry)
	onSubmit     func(*Entry)

	Offset int
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
		e.text.SetMaxWidth(s.X)

		text := e.visibleText()

		p.FillRect(0, 0, s.X, 1)
		p.DrawText(0, 0, text)

		if e.IsFocused() {
			pos := e.text.CursorPos()
			p.DrawCursor(pos.X-e.Offset, 0)
		}
	})
}

// CursorPos returns the coordinate for the cursor for a given width.
func (e *Entry) CursorPos() image.Point {
	return e.text.CursorPos()
}

// SizeHint returns the recommended size hint for the entry.
func (e *Entry) SizeHint() image.Point {
	return image.Point{10, 1}
}

// OnKeyEvent handles key events.
func (e *Entry) OnKeyEvent(ev KeyEvent) {
	if !e.IsFocused() {
		return
	}

	screenWidth := e.Size().X
	e.text.SetMaxWidth(screenWidth)

	if ev.Key != KeyRune {
		switch ev.Key {
		case KeyEnter:
			if e.onSubmit != nil {
				e.onSubmit(e)
			}
		case KeyBackspace2:
			e.text.Backspace()
			if e.Offset > 0 && !e.isTextRemaining() {
				e.Offset--
			}
			if e.onTextChange != nil {
				e.onTextChange(e)
			}
		case KeyDelete, KeyCtrlD:
			e.text.Delete()
			if e.onTextChange != nil {
				e.onTextChange(e)
			}
		case KeyLeft, KeyCtrlB:
			e.text.MoveBackward()
			if e.Offset > 0 {
				e.Offset--
			}
		case KeyRight, KeyCtrlF:
			e.text.MoveForward()

			isCursorTooFar := e.text.CursorPos().X >= screenWidth
			isTextLeft := (e.text.Width() - e.Offset) > (screenWidth - 1)

			if isCursorTooFar && isTextLeft {
				e.Offset++
			}
		case KeyHome, KeyCtrlA:
			e.text.MoveToLineStart()
			e.Offset = 0
		case KeyEnd, KeyCtrlE:
			e.text.MoveToLineEnd()
			left := e.text.Width() - (screenWidth - 1)
			if left >= 0 {
				e.Offset = left
			}
		case KeyCtrlK:
			e.text.Kill()
		}
		return
	}

	e.text.WriteRune(ev.Rune)
	if e.text.CursorPos().X >= screenWidth {
		e.Offset++
	}
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
	e.text.Set([]rune(text))
}

// Text returns the text content of the entry.
func (e *Entry) Text() string {
	return e.text.String()
}

func (e *Entry) visibleText() string {
	text := e.text.String()
	if text == "" {
		return ""
	}
	windowStart := e.Offset
	windowEnd := e.Size().X + windowStart
	if windowEnd > len(text) {
		windowEnd = len(text)
	}
	return text[windowStart:windowEnd]
}

func (e *Entry) isTextRemaining() bool {
	return e.text.Width()-e.Offset > e.Size().X
}
