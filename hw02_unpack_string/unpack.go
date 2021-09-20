package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if unicode.IsNumber(rune(s[0])) {
		return "", ErrInvalidString
	}
	runeArr := s
	var result string
	var prev rune

	for i, cur := range runeArr {
		if i < 1 {
			prev = cur
			result += string(cur)
			continue
		}
		if unicode.IsNumber(prev) && unicode.IsNumber(cur) {
			return "", ErrInvalidString
		}

		if unicode.IsNumber(cur) {
			if cur == 48 {
				result = result[:i-1]
				continue
			}

			cur, _ := strconv.ParseInt(string(cur), 10, 32)
			result += strings.Repeat(string(runeArr[i-1]), int(cur)-1)
		} else {
			result += string(cur)
		}
		prev = cur
	}
	return result, nil
}
