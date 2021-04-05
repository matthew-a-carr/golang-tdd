package iteration

import "strings"

const emptyString = ""

func Repeat(character string, times int) string {
	if times < 0 {
		return emptyString
	}
	return strings.Repeat(character, times)
}
