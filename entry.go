package tui

import (
	"image"
	"strings"
)

var _ Widget = &Entry{}

// EchoMode is used to determine the visibility of Entry text.
type EchoMode int

const (
	// EchoModeNormal displays the characters as they're being entered.
	EchoModeNormal EchoMode = iota

	// EchoModeNoEcho disables text display. This is useful for when the length
	// of the password should be kept secret.
	EchoModeNoEcho

	// EchoModePassword replaces all characters with asterisks.
	EchoModePassword
)

// Entry is a one-line text editor. It lets the user supply the application
// with text, e.g., to input user and password information.
type Entry struct {
	WidgetBase

	text RuneBuffer

	onTextChange func(*Entry)
	onSubmit     func(*Entry)

	echoMode EchoMode
	offset   int
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

		switch e.echoMode {
		case EchoModeNormal:
			p.DrawText(0, 0, text)
		case EchoModePassword:
			p.DrawText(0, 0, strings.Repeat("*", len(text)))
		}

		if e.IsFocused() {
			var off int
			if e.echoMode != EchoModeNoEcho {
				off = e.text.CursorPos().X - e.offset
			}
			p.DrawCursor(off, 0)
		}
	})
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
		case KeyBackspace:
			fallthrough
		case KeyBackspace2:
			e.text.Backspace()
			if e.offset > 0 && !e.isTextRemaining() {
				e.offset--
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
			if e.offset > 0 {
				e.offset--
			}
		case KeyRight, KeyCtrlF:
			e.text.MoveForward()

			isCursorTooFar := e.text.CursorPos().X >= screenWidth
			isTextLeft := (e.text.Width() - e.offset) > (screenWidth - 1)

			if isCursorTooFar && isTextLeft {
				e.offset++
			}
		case KeyHome, KeyCtrlA:
			e.text.MoveToLineStart()
			e.offset = 0
		case KeyEnd, KeyCtrlE:
			e.text.MoveToLineEnd()
			e.ensureCursorIsVisible()
		case KeyCtrlK:
			e.text.Kill()
		}
		return
	}

	e.text.WriteRune(ev.Rune)
	if e.text.CursorPos().X >= screenWidth {
		e.offset++
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

// SetEchoMode sets the echo mode of the entry.
func (e *Entry) SetEchoMode(m EchoMode) {
	e.echoMode = m
}

// SetText sets the text content of the entry.
func (e *Entry) SetText(text string) {
	e.text.Set([]rune(text))
	// TODO: Enable when RuneBuf supports cursor movement for CJK.
	// e.ensureCursorIsVisible()
	e.offset = 0
}

func (e *Entry) ensureCursorIsVisible() {
	left := e.text.Width() - (e.Size().X - 1)
	if left >= 0 {
		e.offset = left
	} else {
		e.offset = 0
	}
}

// Text returns the text content of the entry.
func (e *Entry) Text() string {
	return e.text.String()
}

func (e *Entry) visibleText() string {
	text := e.text
	if text.Len() == 0 {
		return ""
	}
	windowStart := e.offset
	windowEnd := e.Size().X + windowStart
	if windowEnd > text.Len() {
		windowEnd = text.Len()
	}
	return string(text.Runes()[windowStart:windowEnd])
}

func (e *Entry) isTextRemaining() bool {
	return e.text.Width()-e.offset > e.Size().X
}
