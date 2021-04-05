package iteration

import "strings"

const emptyString = ""

// Repeat returns a new string repeated the given number of times.
// If times is negative, then an empty string is returned.
func Repeat(character string, times int) string {
	if times < 0 {
		return emptyString
	}
	return strings.Repeat(character, times)
}
