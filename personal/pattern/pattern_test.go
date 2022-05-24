package pattern

import (
	"testing"
)

type patternTest struct {
	in, out string
}

var patternTests = []patternTest{
	patternTest{"aaaaa", "a"},
	patternTest{"ababab", "ab"},
	patternTest{"abaaba", "aba"},
	patternTest{"c.c.c.", "c"},
	patternTest{"abcdefg", "abcdefg"},
}

func TestPatterns(t *testing.T) {
	for _, e := range patternTests {
		v := getPattern(e.in)
		if v != e.out {
			t.Errorf("getPattern(%s), but wnat %s", e.in, e.out)
		}
	}
}
