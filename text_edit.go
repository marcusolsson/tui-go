package tui

import (
	"image"
	"strings"
)

var _ Widget = &TextEdit{}

// TextEdit is a multi-line text editor.
type TextEdit struct {
	WidgetBase

	text   RuneBuffer
	offset int

	onTextChange func(*TextEdit)
}

// NewTextEdit returns a new TextEdit.
func NewTextEdit() *TextEdit {
	return &TextEdit{}
}

// Draw draws the entry.
func (e *TextEdit) Draw(p *Painter) {
	style := "entry"
	if e.IsFocused() {
		style += ".focused"
	}
	p.WithStyle(style, func(p *Painter) {
		s := e.Size()
		e.text.SetMaxWidth(s.X)

		lines := e.text.SplitByLine()
		for i, line := range lines {
			p.FillRect(0, i, s.X, 1)
			p.DrawText(0, i, line)
		}
		if e.IsFocused() {
			pos := e.text.CursorPos()
			p.DrawCursor(pos.X, pos.Y)
		}
	})
}

// SizeHint returns the recommended size for the entry.
func (e *TextEdit) SizeHint() image.Point {
	var max int
	lines := strings.Split(e.text.String(), "\n")
	for _, line := range lines {
		if w := stringWidth(line); w > max {
			max = w
		}
	}
	return image.Point{max, e.text.heightForWidth(max)}
}

// OnKeyEvent handles key events.
func (e *TextEdit) OnKeyEvent(ev KeyEvent) {
	if !e.IsFocused() {
		return
	}

	screenWidth := e.Size().X
	e.text.SetMaxWidth(screenWidth)

	if ev.Key != KeyRune {
		switch ev.Key {
		case KeyEnter:
			e.text.WriteRune('\n')
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
			left := e.text.Width() - (screenWidth - 1)
			if left >= 0 {
				e.offset = left
			}
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

// OnTextChanged sets a function to be run whenever the text content of the
// widget has been changed.
func (e *TextEdit) OnTextChanged(fn func(entry *TextEdit)) {
	e.onTextChange = fn
}

// SetText sets the text content of the entry.
func (e *TextEdit) SetText(text string) {
	e.text.Set([]rune(text))
}

// Text returns the text content of the entry.
func (e *TextEdit) Text() string {
	return e.text.String()
}

// SetWordWrap sets whether the text should wrap or not.
func (e *TextEdit) SetWordWrap(enabled bool) {
	e.text.wordwrap = enabled
}

func (e *TextEdit) isTextRemaining() bool {
	return e.text.Width()-e.offset > e.Size().X
}
