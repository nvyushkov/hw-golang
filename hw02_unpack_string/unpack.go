package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder

	for i, r := range str {
		if i+1 >= len(str) {
			result.WriteRune(r)
			continue
		}
		next := rune(str[i+1])
		nextIsDigit := unicode.IsDigit(next)

		if unicode.IsDigit(r) {
			if nextIsDigit || i == 0 {
				return "", ErrInvalidString
			}
			continue
		}

		if !nextIsDigit {
			result.WriteRune(r)
			continue
		}

		count := int(next - '0')
		if next < 1 {
			continue
		}
		s := strings.Repeat(string(r), count)
		result.WriteString(s)
	}
	return result.String(), nil
}
