package tui

type FocusChain interface {
	FocusNext(w Widget) Widget
	FocusPrev(w Widget) Widget
	FocusDefault() Widget
}

type KbFocusController struct {
	focusedWidget Widget

	chain FocusChain
}

func (c *KbFocusController) OnEvent(e Event) {
	if e.Type != EventKey || c.chain == nil {
		return
	}
	switch e.Key {
	case KeyTab:
		if c.focusedWidget != nil {
			c.focusedWidget.SetFocused(false)
			c.focusedWidget = c.chain.FocusNext(c.focusedWidget)
			c.focusedWidget.SetFocused(true)
		}
	case KeyBacktab:
		if c.focusedWidget != nil {
			c.focusedWidget.SetFocused(false)
			c.focusedWidget = c.chain.FocusPrev(c.focusedWidget)
			c.focusedWidget.SetFocused(true)
		}
	}
}

var DefaultFocusChain = &SimpleFocusChain{
	widgets: make([]Widget, 0),
}

type SimpleFocusChain struct {
	widgets []Widget
}

func (c *SimpleFocusChain) Set(ws ...Widget) {
	c.widgets = ws
}

func (c *SimpleFocusChain) FocusNext(current Widget) Widget {
	for i, w := range c.widgets {
		if w != current {
			continue
		}
		if i < len(c.widgets)-1 {
			return c.widgets[i+1]
		}
		return c.widgets[0]
	}
	return nil
}

func (c *SimpleFocusChain) FocusPrev(current Widget) Widget {
	for i, w := range c.widgets {
		if w != current {
			continue
		}
		if i <= 0 {
			return c.widgets[len(c.widgets)-1]
		}
		return c.widgets[i-1]
	}
	return nil
}

func (c *SimpleFocusChain) FocusDefault() Widget {
	if len(c.widgets) == 0 {
		return nil
	}
	return c.widgets[0]
}
