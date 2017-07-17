package tui

type Color int

const (
	ColorDefault Color = iota
	ColorBlack
	ColorWhite
	ColorRed
	ColorGreen
	ColorBlue
	ColorCyan
	ColorMagenta
	ColorYellow
)

type Style struct {
	Fg      Color
	Bg      Color
	Reverse bool
}

type Theme struct {
	styles map[string]Style
}

var DefaultTheme = &Theme{
	styles: map[string]Style{
		"list.item.selected":  {Reverse: true},
		"table.cell.selected": {Reverse: true},
		"button.focused":      {Reverse: true},
		"box.focused":         {Reverse: true},
	},
}

func NewTheme() *Theme {
	return &Theme{
		styles: make(map[string]Style),
	}
}

func (p *Theme) SetStyle(n string, i Style) {
	p.styles[n] = i
}

func (p *Theme) Style(name string) Style {
	if c, ok := p.styles[name]; ok {
		return c
	}
	return Style{Fg: ColorDefault, Bg: ColorDefault}
}

func (p *Theme) HasStyle(name string) bool {
	_, ok := p.styles[name]
	return ok
}
