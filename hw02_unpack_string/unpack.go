package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func stringRepeat(num rune, r rune) (string, error) {
	count := int(num - '0')
	if num < 1 {
		return "", errors.New("count is 0")
	}
	return strings.Repeat(string(r), count), nil
} // повторяем строку если можно

func nextChar(i int, str string) (rune, error) {
	if i+1 >= len(str) {
		return 0, ErrInvalidString
	}
	result := rune(str[i+1])
	return result, nil
} // берем следующий символ

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

		// получаем следующий от текущего элемент
		next, err := nextChar(i, str)
		if err != nil {
			result.WriteRune(r)
			break
		}

		if !unicode.IsDigit(next) {
			result.WriteRune(r)
			continue
		}

		s, err := stringRepeat(next, r)
		if err != nil {
			continue
		}
		result.WriteString(s)
	}

	return result.String(), nil
}
