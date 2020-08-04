package pkg

import "strings"

// Upper cases and adds exclamation mark
func MakeSomeNoise(sentence string) string {
	return strings.ToUpper(sentence) + "!"
}
