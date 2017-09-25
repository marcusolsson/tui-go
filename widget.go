package tui

import "image"

// SizePolicy determines the space occupied by a widget.
type SizePolicy int

const (
	// Preferred interprets the size hint as the preferred size.
	Preferred SizePolicy = iota
	// Minimum allows the widget to shrink down to the size hint.
	Minimum
	// Maximum allows the widget to grow up to the size hint.
	Maximum
	// Expanding makes the widget expand to the available space.
	Expanding
)

// Widget defines common operations on widgets.
type Widget interface {
	Draw(p *Painter)
	MinSizeHint() image.Point
	Size() image.Point
	SizeHint() image.Point
	SizePolicy() (SizePolicy, SizePolicy)
	Resize(size image.Point)
	OnKeyEvent(ev KeyEvent)
	SetFocused(bool)
	IsFocused() bool
}

// WidgetBase defines base attributes and operations for all widgets.
type WidgetBase struct {
	size image.Point

	sizePolicyX SizePolicy
	sizePolicyY SizePolicy

	focused bool
}

// Draw is an empty operation to fulfill the Widget interface.
func (w *WidgetBase) Draw(p *Painter) {
}

// SetFocused focuses the widget.
func (w *WidgetBase) SetFocused(f bool) {
	w.focused = f
}

// IsFocused returns whether the widget is focused.
func (w *WidgetBase) IsFocused() bool {
	return w.focused
}

// MinSizeHint returns the size below which the widget cannot shrink.
func (w *WidgetBase) MinSizeHint() image.Point {
	return image.Point{1, 1}
}

// Size returns the current size of the widget.
func (w *WidgetBase) Size() image.Point {
	return w.size
}

// SizeHint returns the size hint of the widget.
func (w *WidgetBase) SizeHint() image.Point {
	return image.ZP
}

// SetSizePolicy sets the size policy for horizontal and vertical directions.
func (w *WidgetBase) SetSizePolicy(h, v SizePolicy) {
	w.sizePolicyX = h
	w.sizePolicyY = v
}

// SizePolicy returns the current size policy.
func (w *WidgetBase) SizePolicy() (SizePolicy, SizePolicy) {
	return w.sizePolicyX, w.sizePolicyY
}

// Resize sets the size of the widget.
func (w *WidgetBase) Resize(size image.Point) {
	w.size = size
}

// OnKeyEvent is an empty operation to fulfill the Widget interface.
func (w *WidgetBase) OnKeyEvent(ev KeyEvent) {
}
