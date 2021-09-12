package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const SLASH rune = 92

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
	isEscaped := false
	slashCounter := 0

	for i, cur := range runeArr {
		if i < 1 {
			prev = cur
			result += string(cur)
			continue
		}
		if unicode.IsNumber(prev) && unicode.IsNumber(cur) && !isEscaped {
			return "", ErrInvalidString
		}

		if isSlash(prev) && unicode.IsNumber(cur) && slashCounter%2 == 1 {
			isEscaped = true
		}

		if unicode.IsNumber(cur) {
			if cur == 48 {
				result = result[:i-1]
				continue
			}

			if isSlash(prev) && isEscaped {
				result += string(cur)
			} else {
				cur, _ := strconv.ParseInt(string(cur), 10, 32)
				result += strings.Repeat(string(runeArr[i-1]), int(cur)-1)
			}
		} else if isSlash(cur) {
			slashCounter++
			if !isSlash(prev) {
				prev = cur
				continue
			}
			if isSlash(prev) && slashCounter%2 == 0 {
				result += string(cur)
				prev = cur
				continue
			}
		} else {
			result += string(cur)
			slashCounter = 0
		}
		prev = cur
	}
	return result, nil
}

func isSlash(r rune) bool {
	return r == SLASH
}
