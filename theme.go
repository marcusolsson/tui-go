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
	Fg Color
	Bg Color
}

type Theme struct {
	styles map[string]Style
}

var DefaultTheme = &Theme{
	styles: map[string]Style{
		"normal":              {ColorDefault, ColorDefault},
		"list.item.selected":  {ColorWhite, ColorBlue},
		"table.cell.selected": {ColorWhite, ColorBlue},
		"button.focused":      {ColorWhite, ColorBlue},
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
	return Style{ColorDefault, ColorDefault}
}
