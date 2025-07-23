package iteration

import "strings"

const repeatCount = 5

func Repeat(character string, repeats int) string {
	var repeated strings.Builder
	// Use a loop to append the character 'repeatCount' times
	for i := 0; i < repeats; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
