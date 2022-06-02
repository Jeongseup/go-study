package pattern

import "strings"

func getPattern(in string) string {
	strLen := len(in)
	for i := 0; i < (strLen / 2); i++ {
		pattern := in[0 : i+1]
		patternLen := len(pattern)
		result := strings.Repeat(pattern, strLen/patternLen)
		if result == in {
			return pattern
		}
	}
	return pattern
}
