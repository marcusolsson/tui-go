package main

import (
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
	if s.IsFocused() && ev.Key == tui.KeyEnter && s.OnSubmit != nil{
		s.OnSubmit()
	}
}

type ItemEditor struct {
	*tui.Box
	tui.SimpleFocusChain

	nameEdit  *tui.Entry
	valueEdit *tui.Entry

	Done func()
}

func NewItemEditor(fm *FocusManager, i *ListItem) *ItemEditor {
	e := &ItemEditor{
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
		Label: tui.NewLabel("[OK]"),
		OnSubmit: func() {
			if e.Done != nil {
				e.Done()
			}
		},
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
	fm.Set(e)

	return e
}
