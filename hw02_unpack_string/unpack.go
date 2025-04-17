package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(originalString string) (string, error) {
	var lastSymbol rune
	var unpackedString strings.Builder
	isEscape := false
	reps := false

	for i, s := range originalString {
		if i == 0 && unicode.IsDigit(s) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(s) && reps {
			return "", ErrInvalidString
		}

		if i == len(originalString)-1 && !isEscape && s == '\\' {
			return "", ErrInvalidString
		}

		if s == '\\' && !isEscape {
			isEscape = true
			unpackedString.WriteString(string(lastSymbol))
			lastSymbol = 0
			continue
		}

		if isEscape {
			if s != '\\' && !unicode.IsDigit(s) {
				return "", ErrInvalidString
			}
			isEscape = false
		}

		if unicode.IsDigit(s) && lastSymbol != 0 {
			n, _ := strconv.Atoi(string(s))
			unpackedString.WriteString(strings.Repeat(string(lastSymbol), n))
			lastSymbol = 0
			reps = true
			continue
		}

		if lastSymbol != 0 {
			unpackedString.WriteString(string(lastSymbol))
		}

		lastSymbol = s
		reps = false

		if i == len(originalString)-1 {
			unpackedString.WriteString(string(lastSymbol))
		}
	}

	return unpackedString.String(), nil
}
