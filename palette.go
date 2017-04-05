package tui

type Color int

const (
	ColorDefault Color = iota
	ColorBlack
	ColorWhite
	ColorBlue
	ColorRed
)

type PaletteItem struct {
	Fg Color
	Bg Color
}

type Palette struct {
	items map[string]PaletteItem
}

var DefaultPalette = &Palette{
	items: map[string]PaletteItem{
		"normal":              {ColorDefault, ColorDefault},
		"list.item.selected":  {ColorWhite, ColorBlue},
		"table.cell.selected": {ColorWhite, ColorBlue},
		"button.focused":      {ColorWhite, ColorBlue},
	},
}

func NewPalette() *Palette {
	return &Palette{
		items: make(map[string]PaletteItem),
	}
}

func (p *Palette) SetItem(n string, i PaletteItem) {
	p.items[n] = i
}

func (p *Palette) Item(name string) PaletteItem {
	if c, ok := p.items[name]; ok {
		return c
	}
	return PaletteItem{ColorDefault, ColorDefault}
}
