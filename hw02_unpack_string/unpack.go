package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const Shield = '\\'

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder

	skip := true
	shielding := false
	current := '1'
	prev := '1'

	for _, current = range s {
		println(string(current))
		if prev == Shield && !shielding {
			if !unicode.IsDigit(current) && current != Shield {
				println("============" + result.String() + "===========")
				return result.String(), ErrInvalidString
			}
			shielding = true
			prev = current
			continue
		} else {
			if unicode.IsDigit(current) && unicode.IsDigit(prev) && !shielding {
				return result.String(), ErrInvalidString
			}
			shielding = false
		}

		if skip {
			skip = false
		} else {
			count, err := strconv.Atoi(string(current))
			if err == nil && !shielding {
				add := strings.Repeat(string(prev), count)
				result.WriteString(add)
				skip = true
			} else {
				result.WriteString(string(prev))
			}
		}
		prev = current
	}
	if !unicode.IsDigit(current) || shielding {
		result.WriteString(string(current))
	}
	if current == Shield && !shielding {
		return result.String(), ErrInvalidString
	}

	return result.String(), nil
}
