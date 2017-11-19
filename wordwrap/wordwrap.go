package wordwrap

import (
	"bytes"
	"unicode"
)

// WrapString wraps the input string by inserting newline characters. It does
// not remove whitespace, but preserves the original text.
func WrapString(s string, width int) string {
	if len(s) <= 1 {
		return s
	}

	var buf bytes.Buffer
	var word bytes.Buffer

	word.WriteByte(s[0])

	spaceLeft := width
	for i := 1; i < len(s); i++ {
		curr := s[i]
		prev := s[i-1]

		if curr == '\n' {
			if word.Len() > spaceLeft {
				spaceLeft = width
				buf.WriteRune('\n')
			} else {
				spaceLeft = width
			}
			// fmt.Printf("33: writing %q with %d spaces remaining out of %d\n", word.String(), spaceLeft, width)
			word.WriteTo(&buf)
		} else if unicode.IsSpace(rune(prev)) && !unicode.IsSpace(rune(curr)) {
			if word.Len() > spaceLeft {
				spaceLeft = width - word.Len()
				buf.WriteRune('\n')
			} else {
				spaceLeft -= word.Len()
			}
			// fmt.Printf("42: writing %q with %d spaces remaining out of %d\n", word.String(), spaceLeft, width)
			word.WriteTo(&buf)
		}
		word.WriteByte(curr)
	}

	if word.Len() > spaceLeft {
		buf.WriteRune('\n')
	}
	// fmt.Printf("51: writing %q with %d spaces remaining out of %d\n", word.String(), spaceLeft, width)
	word.WriteTo(&buf)

	return buf.String()
}
