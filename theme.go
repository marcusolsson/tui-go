package tui

// Color represents a color.
type Color int

// Common colors.
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

// Style determines how a cell should be painted.
type Style struct {
	Fg      Color
	Bg      Color
	Reverse bool

	Bold      bool
	Underline bool
}

// Theme defines the styles for a set of identifiers.
type Theme struct {
	styles map[string]Style
}

// DefaultTheme is a theme with reasonable defaults.
var DefaultTheme = &Theme{
	styles: map[string]Style{
		"list.item.selected":  {Reverse: true},
		"table.cell.selected": {Reverse: true},
		"button.focused":      {Reverse: true},
		"box.focused":         {Reverse: true},
	},
}

// NewTheme return an empty theme.
func NewTheme() *Theme {
	return &Theme{
		styles: make(map[string]Style),
	}
}

// SetStyle sets a style for a given identifier.
func (p *Theme) SetStyle(n string, i Style) {
	p.styles[n] = i
}

// Style returns the style associated with an identifier.
func (p *Theme) Style(name string) Style {
	if c, ok := p.styles[name]; ok {
		return c
	}
	return Style{Fg: ColorDefault, Bg: ColorDefault}
}

// HasStyle returns whether an identifier is associated with an identifier.
func (p *Theme) HasStyle(name string) bool {
	_, ok := p.styles[name]
	return ok
}
