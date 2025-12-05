package other

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstNonEmpty(t *testing.T) {
	stringCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{"AllEmpty", []string{"", "", ""}, ""},
		{"FirstNonEmpty", []string{"", "value", "another"}, "value"},
		{"OnlyNonEmpty", []string{"value"}, "value"},
		{"AllNonEmpty", []string{"value1", "value2"}, "value1"},
		{"MixedZeroAndValue", []string{"0", "1", "0"}, "0"},
	}

	for _, tt := range stringCases {
		t.Run(tt.name, func(t *testing.T) {
			result := FirstNonEmpty(tt.input...)
			assert.Equal(t, tt.expected, result)
		})
	}

	intCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"AllEmpty", []int{0, 0, 0}, 0},
		{"FirstNonEmpty", []int{0, 1, 2}, 1},
		{"OnlyNonEmpty", []int{1}, 1},
		{"AllNonEmpty", []int{1, 2}, 1},
		{"MixedZeroAndValue", []int{0, 1, 0}, 1},
	}

	for _, tt := range intCases {
		t.Run(tt.name, func(t *testing.T) {
			result := FirstNonEmpty(tt.input...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkFirstNonEmpty(b *testing.B) {
	stringCases := []struct {
		name string
		args []string
	}{
		{"AllEmpty", []string{"", "", ""}},
		{"FirstNonEmpty", []string{"", "value", "another"}},
		{"OnlyNonEmpty", []string{"value"}},
		{"AllNonEmpty", []string{"value1", "value2"}},
		{"MixedZeroAndValue", []string{"0", "1", "0"}},
	}

	for _, tt := range stringCases {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FirstNonEmpty(tt.args...)
			}
		})
	}

	intCases := []struct {
		name string
		args []int
	}{
		{"AllEmpty", []int{0, 0, 0}},
		{"FirstNonEmpty", []int{0, 1, 2}},
		{"OnlyNonEmpty", []int{1}},
		{"AllNonEmpty", []int{1, 2}},
		{"MixedZeroAndValue", []int{0, 1, 0}},
	}

	for _, tt := range intCases {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FirstNonEmpty(tt.args...)
			}
		})
	}
}
