package tui

const (
	runeMinus      = '-'
	runePlusSmall  = '+'
	runeVLineSmall = '|'

	runePlus     = '┼'
	runeHLine    = '─'
	runeVLine    = '│'
	runeTTee     = '┬'
	runeRTee     = '┤'
	runeLTee     = '├'
	runeBTee     = '┴'
	runeULCorner = '┌'
	runeURCorner = '┐'
	runeLLCorner = '└'
	runeLRCorner = '┘'
)

var runeFallbacks = map[rune]rune{}

func init() {
	defaultBorder()
}

// default
//	┌──────────────────┐
//	│┌────────┬───────┐│
//	││A       │apple  ││
//	│└────────┴───────┘│
//	└──────────────────┘
func defaultBorder() {
	runeFallbacks = map[rune]rune{
		runePlus:     runePlus,
		runeHLine:    runeHLine,
		runeVLine:    runeVLine,
		runeTTee:     runeTTee,
		runeRTee:     runeRTee,
		runeLTee:     runeLTee,
		runeBTee:     runeBTee,
		runeULCorner: runeULCorner,
		runeURCorner: runeURCorner,
		runeLLCorner: runeLLCorner,
		runeLRCorner: runeLRCorner,
	}
}

// small
//	+------------------+
//	|+--------+-------+|
//	||A       |apple  ||
//	|+--------+-------+|
//	+------------------+
func simpleBorder() {
	runeFallbacks = map[rune]rune{
		runeHLine: runeMinus,
		runeVLine: runeVLineSmall,

		runePlus:     runePlusSmall,
		runeLLCorner: runePlusSmall,
		runeLRCorner: runePlusSmall,
		runeTTee:     runePlusSmall,
		runeRTee:     runePlusSmall,
		runeLTee:     runePlusSmall,
		runeBTee:     runePlusSmall,
		runeULCorner: runePlusSmall,
		runeURCorner: runePlusSmall,
	}
}

// GetBorder get a simple or default border in a compatible way
func GetBorder(k rune) rune {
	return runeFallbacks[k]
}
