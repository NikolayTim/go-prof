package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var previousRune rune
	result := strings.Builder{}

	if str == result.String() {
		return "", nil
	}

	for i, currentRune := range str {
		if i == 0 {
			if unicode.IsDigit(currentRune) {
				return "", ErrInvalidString
			}

			previousRune = currentRune
			continue
		}

		if counts, err := strconv.Atoi(string(currentRune)); err == nil {
			if unicode.IsDigit(previousRune) {
				return "", ErrInvalidString
			}

			repeatedString := strings.Repeat(string(previousRune), counts)
			if _, err := result.WriteString(repeatedString); err != nil {
				return "", ErrInvalidString
			}
		} else if !unicode.IsDigit(previousRune) {
			if _, err := result.WriteRune(previousRune); err != nil {
				return "", ErrInvalidString
			}
		}

		previousRune = currentRune
	}

	if !unicode.IsDigit(previousRune) {
		if _, err := result.WriteRune(previousRune); err != nil {
			return "", ErrInvalidString
		}
	}

	return result.String(), nil
}
