package iteration

import "strings"

func Repeat(character string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}

func Replace(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}
