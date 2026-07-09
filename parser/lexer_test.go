package parser

import (
	"testing"
)

func TestParseSrcInvalidRune(t *testing.T) {
	// Regression test for issue #341. The scanner did not advance past
	// runes it could not tokenize, so the parser looped forever.
	tests := []string{
		"",
		"a = ",
		"",
		"@",
		"$",
	}
	for _, src := range tests {
		_, err := ParseSrc(src)
		if err == nil {
			t.Errorf("ParseSrc(%q) expected syntax error, got nil", src)
		}
	}
}
