package main

import (
	"strings"

	"github.com/marcusolsson/tui-go"
)

type SubmitLabel struct {
	*tui.Label

	OnSubmit func()
}

func (s *SubmitLabel) Draw(p *tui.Painter) {
	if s.IsFocused() {
		p.WithStyle("reversed", s.Label.Draw)
	} else {
		s.Label.Draw(p)
	}
}

func (s *SubmitLabel) OnKeyEvent(ev tui.KeyEvent) {
	if s.IsFocused() && ev.Key == tui.KeyEnter && s.OnSubmit != nil {
		s.OnSubmit()
	}
}

type ItemEditor struct {
	*tui.Box
	tui.SimpleFocusChain

	nameEdit  *tui.Entry
	valueEdit *tui.Entry
	parent    *List
	item      *ListItem

	Done func()
}

func NewItemEditor(parent *List, i *ListItem) *ItemEditor {
	e := &ItemEditor{
		parent:    parent,
		item:      i,
		nameEdit:  tui.NewEntry(),
		valueEdit: tui.NewEntry(),
	}
	e.nameEdit.SetText(i.name.Text())
	e.valueEdit.SetText(i.value.Text())

	cancel := &SubmitLabel{
		Label: tui.NewLabel("[Cancel]"),
		OnSubmit: func() {
			if e.Done != nil {
				e.Done()
			}
		},
	}
	done := &SubmitLabel{
		Label:    tui.NewLabel("[OK]"),
		OnSubmit: e.MaybeSubmit,
	}

	e.Box = tui.NewVBox(
		tui.NewHBox(
			tui.NewLabel("Username: "),
			e.nameEdit,
		),
		tui.NewHBox(
			tui.NewLabel("Real name: "),
			e.valueEdit,
		),
		tui.NewHBox(cancel, done),
		tui.NewSpacer(),
	)
	e.Box.SetBorder(true)

	// Set up focus chain
	e.Set(e.nameEdit, e.valueEdit, cancel, done)
	e.parent.FocusManager.Set(e)

	return e
}

func (e *ItemEditor) MaybeSubmit() {
	name := strings.TrimSpace(e.nameEdit.Text())
	value := strings.TrimSpace(e.valueEdit.Text())

	if name == "" {
		e.popErr("No username given!")
		return
	}
	if value == "" {
		e.popErr("No real name given!")
		return
	}

	e.item.name.SetText(name)
	e.item.value.SetText(value)

	// Looks OK; prepend & return.
	e.parent.Commit(e.item)
	e.Done()
}

func (e *ItemEditor) popErr(msg string) {
	done := func() {
		e.parent.FocusManager.Set(e)
		e.parent.UI.SetWidget(e)
	}

	submit := &SubmitLabel{
		Label:    tui.NewLabel("[OK]"),
		OnSubmit: done,
	}
	fc := &tui.SimpleFocusChain{}
	fc.Set([]tui.Widget{submit}...)
	e.parent.FocusManager.Set(fc)

	w := tui.NewVBox(tui.NewLabel(msg), submit)
	w.SetBorder(true)
	wrap := tui.NewVBox(
		tui.NewSpacer(),
		tui.NewHBox(
			tui.NewSpacer(),
			w,
			tui.NewSpacer(),
		),
		tui.NewSpacer(),
	)

	e.parent.UI.SetWidget(wrap)
}
