package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	res := strings.Builder{}
	r := rune(0)

	for _, c := range s {
		if unicode.IsDigit(c) {
			if r == 0 {
				return "", ErrInvalidString
			}
			ctr, _ := strconv.Atoi(string(c))
			res.WriteString(strings.Repeat(string(r), ctr))
			r = rune(0)
		} else {
			if r != 0 {
				res.WriteRune(r)
			}
			r = c
		}
	}
	if r != 0 {
		res.WriteRune(r)
	}
	return res.String(), nil
}
