package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder
	prevDigital := false

	d, _ := strconv.Atoi(str)
	if d > 0 {
		return "", ErrInvalidString
	}

	for i, r := range str {
		if unicode.IsDigit(r) {
			if i == 0 || prevDigital {
				return "", ErrInvalidString
			}
			prevDigital = true
			continue
		}
		prevDigital = false

		if i+1 >= len(str) {
			result.WriteRune(r)
			break
		}
		next := rune(str[i+1])

		if !unicode.IsDigit(next) {
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
