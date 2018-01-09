package main

import (
	"github.com/marcusolsson/tui-go"
)

type EditorLauncher interface {
	Edit(i *ListItem)
}
type WidgetSetter interface {
	SetWidget(tui.Widget)
}

type ListItem struct {
	*tui.Box

	focus bool

	name  *tui.Label
	value *tui.Label

	editor EditorLauncher
}

func NewItem(e EditorLauncher, name, value string) *ListItem {
	n := tui.NewLabel(name)
	v := tui.NewLabel(value)

	return &ListItem{
		name:   n,
		value:  v,
		Box:    tui.NewHBox(n, v),
		editor: e,
	}
}

func (i *ListItem) SetFocused(v bool) { i.focus = v }
func (i *ListItem) IsFocused() bool    { return i.focus }
func (i *ListItem) Draw(p *tui.Painter) {
	if i.IsFocused() {
		p.WithStyle("reversed", i.Box.Draw)
	} else {
		i.Box.Draw(p)
	}
}

func (i *ListItem) OnKeyEvent(ev tui.KeyEvent) {
	if i.IsFocused() && ev.Key == tui.KeyEnter && i.editor != nil {
		i.editor.Edit(i)
	}
}

type NewWidget struct {
	*tui.Label
	parent *List
}

func (n *NewWidget) OnKeyEvent(ev tui.KeyEvent) {
	if n.IsFocused() && ev.Key == tui.KeyEnter {
		n.parent.Edit(NewItem(n.parent, "", ""))
	}
}
func (n *NewWidget) Draw(p *tui.Painter) {
	if n.IsFocused() {
		p.WithStyle("reversed", n.Label.Draw)
	} else {
		n.Label.Draw(p)
	}
}

type List struct {
	*tui.Box
	tui.SimpleFocusChain
	fm *FocusManager

	contents []tui.Widget
	newWidget *NewWidget

	UI WidgetSetter
}

func NewList(fm *FocusManager) *List {
	l := &List{
		fm: fm,

	}
	l.newWidget = &NewWidget{
		Label:  tui.NewLabel("New item"),
		parent: l,
	}

	l.Box = tui.NewVBox(
		tui.NewSpacer(),
		l.newWidget,
	)
	l.Set(l.newWidget)
	l.fm.Set(l)
	return l
}

func (l *List) Prepend(i *ListItem) {
	l.Box.Prepend(i)
	l.contents = append([]tui.Widget{i}, l.contents...)
	l.Set(append([]tui.Widget{l.newWidget}, l.contents...)...)
}

func (l *List) Edit(i *ListItem) {
	r := NewItemEditor(l.fm, i)

	// No provisions for popups; set & reset the view.
	r.Done = func() {
		l.UI.SetWidget(l)
		l.fm.Set(l)
	}
	l.UI.SetWidget(r)
}

func (l *List) OnKeyEvent(ev tui.KeyEvent) {
	if !l.IsFocused() {
		return
	}
	switch ev.Key {
	default:
		l.Box.OnKeyEvent(ev)
	}
}


