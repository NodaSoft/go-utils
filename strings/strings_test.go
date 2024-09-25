package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxRunes uint
		expected string
	}{
		{"EmptyString", "", 10, ""},
		{"ShortString", "abc", 5, "abc"},
		{"ExactLength", "abcdef", 6, "abcdef"},
		{"Truncated", "abcdef", 4, "abcd"},
		{"LongString", "Привет мир", 6, "Привет"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Truncate(tt.input, tt.maxRunes)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkTruncate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Truncate("This is a long string for benchmark testing", 10)
	}
}
