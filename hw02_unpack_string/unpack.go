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
	if !unicode.IsLetter(rune(s[0])) {
		return "", ErrInvalidString
	}

	var result string
	var prev rune

	for i, cur := range s {
		if i < 1 {
			prev = cur
			result += string(cur)
			continue
		}

		if !unicode.IsLetter(prev) && !unicode.IsLetter(cur) {
			return "", ErrInvalidString
		}

		if !unicode.IsLetter(cur) {
			cur, _ := strconv.ParseInt(string(cur), 10, 32)
			if cur == 0 {
				result = result[:i-1]
				continue
			}
			result += strings.Repeat(string(prev), int(cur) - 1)
		} else {
			result += string(cur)
		}

		prev = cur
	}

	return result, nil
}
