package main

import (
	"fmt"
	"log"

	"github.com/marcusolsson/tui-go"
)

type uiTab struct {
	label *tui.Label
	view  tui.Widget
}

func newTabWidget(views ...*uiTab) *tabWidget {
	topbar := tui.NewHBox(tui.NewLabel("> "))
	topbar.SetSizePolicy(tui.Minimum, tui.Maximum)
	view := &tabWidget{views: views}

	for i := 0; i < len(views); i++ {
		topbar.Append(views[i].label)
	}

	topbar.Append(tui.NewSpacer())
	view.style()

	vbox := tui.NewVBox(topbar, views[0].view)
	vbox.SetSizePolicy(tui.Maximum, tui.Preferred)
	view.Box = vbox
	return view
}

type tabWidget struct {
	*tui.Box

	views  []*uiTab
	active int
}

func (t *tabWidget) OnKeyEvent(ev tui.KeyEvent) {
	switch ev.Key {
	case tui.KeyTab, tui.KeyRight:
		t.Next()
	case tui.KeyBacktab, tui.KeyLeft:
		t.Previous()
	}

	t.Box.OnKeyEvent(ev)
}

func (t *tabWidget) setView(view tui.Widget) {
	t.Box.Remove(1)
	t.Box.Append(view)
}

func (t *tabWidget) style() {
	for i := 0; i < len(t.views); i++ {
		if i == t.active {
			t.views[i].label.SetStyleName("tab-selected")
			continue
		}
		t.views[i].label.SetStyleName("tab")
	}
}

func (t *tabWidget) Next() {
	t.active = clamp(t.active+1, 0, len(t.views)-1)
	t.style()
	t.setView(t.views[t.active].view)
}

func (t *tabWidget) Previous() {
	t.active = clamp(t.active-1, 0, len(t.views)-1)
	t.style()
	t.setView(t.views[t.active].view)
}

func clamp(n, min, max int) int {
	if n < min {
		return max
	}
	if n > max {
		return min
	}
	return n
}

// styledWidget allows painting of a custom style (primarily background color)
// of a given widget that doesn't natively support styling.
type styledWidget struct {
	name string
	*tui.Box
}

func (w *styledWidget) Draw(p *tui.Painter) {
	p.WithStyle(w.name, w.Box.Draw)
}

// keybind is a keybind with a given name, used to render a keybind-bar.
type keybind struct {
	key  string
	name string
}

// newKeybar returns a widget which renders a bar with a list of keybinds.
func newKeybar(keys ...keybind) *styledWidget {
	out := make([]tui.Widget, 0, len(keys)+1)
	for _, s := range keys {
		start := tui.NewLabel("[")
		key := tui.NewLabel(s.key)
		key.SetStyleName("keybar-key")
		end := tui.NewLabel("] " + s.name + "  ")
		out = append(out, start, key, end)
	}

	out = append(out, tui.NewSpacer())
	return &styledWidget{"keybar", tui.NewHBox(out...)}
}

func main() {
	extendedView := tui.NewList()
	for i := 0; i < 20; i++ {
		extendedView.AddItems(fmt.Sprintf("content here x%d", i))
	}

	tabLayout := newTabWidget(
		&uiTab{label: tui.NewLabel(" tab 1 "), view: extendedView},
		&uiTab{label: tui.NewLabel(" tab 2 "), view: tui.NewLabel("some other view here x2")},
		&uiTab{label: tui.NewLabel(" tab 3 "), view: tui.NewLabel("some other view here x3")},
		&uiTab{label: tui.NewLabel(" tab 4 "), view: tui.NewLabel("some other view here x4")},
	)
	tabLayout.SetSizePolicy(tui.Maximum, tui.Expanding)

	keybar := newKeybar(keybind{"left", "prev view"}, keybind{"right", "next view"}, keybind{"esc", "quit"})

	ui, err := tui.New(tui.NewVBox(tabLayout, keybar))
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	theme := tui.NewTheme()
	theme.SetStyle("keybar", tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack})
	theme.SetStyle("label.keybar-key", tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlue, Bold: tui.DecorationOn})
	theme.SetStyle("label.tab", tui.Style{Reverse: tui.DecorationOff})
	theme.SetStyle("label.tab-selected", tui.Style{Reverse: tui.DecorationOn, Fg: tui.ColorMagenta, Bg: tui.ColorWhite})
	ui.SetTheme(theme)

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
