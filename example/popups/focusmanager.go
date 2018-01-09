package main

import (
	"github.com/marcusolsson/tui-go"
)

type FocusManager struct {
	chain tui.FocusChain
	current tui.Widget
}

func (fm *FocusManager) Set(c tui.FocusChain) {
	fm.chain = c
	if fm.current != nil {
		fm.current.SetFocused(false)
	}
	fm.current = fm.chain.FocusDefault()
	fm.current.SetFocused(true)
}

func (fm *FocusManager) next() {
		if fm.chain == nil {
			return
		}
		fm.current.SetFocused(false)
		fm.current = fm.chain.FocusNext(fm.current)
		fm.current.SetFocused(true)
}

func (fm *FocusManager) prev() {
		if fm.chain == nil {
			return
		}
		fm.current.SetFocused(false)
		fm.current = fm.chain.FocusPrev(fm.current)
		fm.current.SetFocused(true)
}

func (fm *FocusManager) Attach(ui tui.UI) {
	ui.SetKeybinding("Down", fm.next)
	ui.SetKeybinding("Right", fm.next)
	ui.SetKeybinding("Tab", fm.next)

	ui.SetKeybinding("Up", fm.prev)
	ui.SetKeybinding("Left", fm.prev)
}
