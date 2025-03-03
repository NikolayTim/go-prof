package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var previousRune rune
	result := strings.Builder{}

	if str == "" {
		return "", nil
	}

	for i, currentRune := range str {
		if i == 0 {
			if _, err := strconv.Atoi(string(currentRune)); err == nil {
				return "", ErrInvalidString
			}

			previousRune = currentRune
			continue
		}

		if counts, err := strconv.Atoi(string(currentRune)); err == nil {
			if _, err := strconv.Atoi(string(previousRune)); err == nil {
				return "", ErrInvalidString
			}

			repeatedString := strings.Repeat(string(previousRune), counts)
			if _, err := result.WriteString(repeatedString); err != nil {
				return "", ErrInvalidString
			}
		} else if _, err := strconv.Atoi(string(previousRune)); err != nil {
			if _, err := result.WriteRune(previousRune); err != nil {
				return "", ErrInvalidString
			}
		}

		previousRune = currentRune
	}

	if _, err := strconv.Atoi(string(previousRune)); err != nil {
		if _, err := result.WriteRune(previousRune); err != nil {
			return "", ErrInvalidString
		}
	}

	return result.String(), nil
}

func UnpackAsterisk(str string) (string, error) {
	var previousRune rune
	var writeLastRune bool
	isSlash := false

	result := strings.Builder{}

	if str == "" {
		return "", nil
	}

	for i, currentRune := range str {
		writeLastRune = true

		if i == 0 {
			if _, err := strconv.Atoi(string(currentRune)); err == nil {
				return "", ErrInvalidString
			}

			previousRune = currentRune
			continue
		}

		if string(currentRune) == "\\" && !isSlash {
			isSlash = true
			continue
		}

		if counts, err := strconv.Atoi(string(currentRune)); err == nil {
			if !isSlash {
				repeatedString := strings.Repeat(string(previousRune), counts)
				if _, err := result.WriteString(repeatedString); err != nil {
					return "", ErrInvalidString
				}
				writeLastRune = false
			} else {
				if _, err := result.WriteRune(previousRune); err != nil {
					return "", ErrInvalidString
				}
			}
		} else {
			if isSlash && string(currentRune) != "\\" {
				return "", ErrInvalidString
			}

			if _, err := result.WriteRune(previousRune); err != nil {
				return "", ErrInvalidString
			}
		}

		isSlash = false
		previousRune = currentRune
	}

	if writeLastRune {
		if _, err := result.WriteRune(previousRune); err != nil {
			return "", ErrInvalidString
		}
	}

	return result.String(), nil
}
