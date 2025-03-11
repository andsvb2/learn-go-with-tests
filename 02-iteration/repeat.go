package iteration

import "strings"

func Repeat(character string, repeatLimit int) string {
	var repeated strings.Builder
	for i := 0; i < repeatLimit; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
