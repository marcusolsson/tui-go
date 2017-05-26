package tui

import "image"

type SizePolicy int

const (
	Preferred SizePolicy = iota
	Minimum
	Maximum
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
	OnEvent(ev Event)
	SetFocused(bool)
	IsFocused() bool
}

type WidgetBase struct {
	size image.Point

	sizePolicyX SizePolicy
	sizePolicyY SizePolicy

	focused bool
}

func (w *WidgetBase) Draw(p *Painter) {
}

func (w *WidgetBase) SetFocused(f bool) {
	w.focused = f
}

func (w *WidgetBase) IsFocused() bool {
	return w.focused
}

func (w *WidgetBase) MinSizeHint() image.Point {
	return image.Point{1, 1}
}

func (w *WidgetBase) Size() image.Point {
	return w.size
}

func (w *WidgetBase) SizeHint() image.Point {
	return image.ZP
}

func (w *WidgetBase) SetSizePolicy(h, v SizePolicy) {
	w.sizePolicyX = h
	w.sizePolicyY = v
}

func (w *WidgetBase) SizePolicy() (SizePolicy, SizePolicy) {
	return w.sizePolicyX, w.sizePolicyY
}

func (w *WidgetBase) Resize(size image.Point) {
	w.size = size
}

func (w *WidgetBase) OnEvent(ev Event) {
}
