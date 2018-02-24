package tui

// FocusChain enables custom focus traversal when Tab or Backtab is pressed.
type FocusChain interface {
	FocusNext(w Widget) Widget
	FocusPrev(w Widget) Widget
	FocusDefault() Widget
}

type kbFocusController struct {
	focusedWidget Widget

	chain FocusChain
}

func (c *kbFocusController) setFocusChain(fc FocusChain) {
	c.chain = fc

	if w := c.chain.FocusDefault(); w != nil {
		w.SetFocused(true)
		c.focusedWidget = w
	}
}

func (c *kbFocusController) OnKeyEvent(e KeyEvent) {
	if c.chain == nil {
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

// DefaultFocusChain is the default focus chain.
var DefaultFocusChain = &SimpleFocusChain{
	widgets: make([]Widget, 0),
}

// SimpleFocusChain represents a ring of widgets where focus is loops to the
// first widget when it reaches the end.
type SimpleFocusChain struct {
	widgets []Widget
}

// Set sets the widgets in the focus chain. Widgets will received focus in the
// order widgets were passed.
func (c *SimpleFocusChain) Set(ws ...Widget) {
	c.widgets = ws
}

// FocusNext returns the widget in the ring that is after the given widget.
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

// FocusPrev returns the widget in the ring that is before the given widget.
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

// FocusDefault returns the default widget for when there is no widget
// currently focused.
func (c *SimpleFocusChain) FocusDefault() Widget {
	if len(c.widgets) == 0 {
		return nil
	}
	return c.widgets[0]
}
