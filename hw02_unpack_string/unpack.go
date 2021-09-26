package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	firstRune, _ := utf8.DecodeRuneInString(s)
	if unicode.IsNumber(firstRune) {
		return "", ErrInvalidString
	}
	runeArr := []rune(s)
	var result strings.Builder
	var prev rune

	for i, cur := range runeArr {
		if i < 1 {
			prev = cur
			result.WriteRune(cur)
			continue
		}
		if unicode.IsNumber(prev) && unicode.IsNumber(cur) {
			return "", ErrInvalidString
		}

		if unicode.IsNumber(cur) {
			if cur == 48 {
				s := []rune(result.String())
				s = s[:i-1]
				result.Reset()
				result.WriteString(string(s))
				continue
			}

			cur, _ := strconv.ParseInt(string(cur), 10, 32)
			result.WriteString(strings.Repeat(string(runeArr[i-1]), int(cur)-1))
		} else {
			result.WriteRune(cur)
		}
		prev = cur
	}
	return result.String(), nil
}
