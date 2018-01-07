package tui

type DialogResult int

const (
	DialogAccepted DialogResult = iota
	DialogRejected
)

type Dialog struct {
	onFinished func(DialogResult)

	Box
}

func NewDialog() *Dialog {
	d := &Dialog{}
	d.SetBorder(true)
	d.alignment = Vertical
	return d
}

func (d *Dialog) OnFinished(fn func(res DialogResult)) {
	d.onFinished = fn
}

func (d *Dialog) OnKeyEvent(ev KeyEvent) {
	switch ev.Key {
	case KeyEnter:
		d.onFinished(DialogAccepted)
	case KeyEsc:
		d.onFinished(DialogRejected)
	}
}
