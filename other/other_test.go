package other

import (
	"errors"
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

func TestIf(t *testing.T) {
	cases := []struct {
		name      string
		condition bool
		then      string
		otherwise string
		expected  string
	}{
		{"TrueCondition", true, "then", "otherwise", "then"},
		{"FalseCondition", false, "then", "otherwise", "otherwise"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result := If(tt.condition, tt.then, tt.otherwise)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIfFunc(t *testing.T) {
	t.Run("TrueCondition", func(t *testing.T) {
		thenCalled := false
		otherwiseCalled := false

		result := IfFunc(true,
			func() string {
				thenCalled = true
				return "then"
			},
			func() string {
				otherwiseCalled = true
				return "otherwise"
			},
		)

		assert.Equal(t, "then", result)
		assert.True(t, thenCalled)
		assert.False(t, otherwiseCalled)
	})

	t.Run("FalseCondition", func(t *testing.T) {
		thenCalled := false
		otherwiseCalled := false

		result := IfFunc(false,
			func() string {
				thenCalled = true
				return "then"
			},
			func() string {
				otherwiseCalled = true
				return "otherwise"
			},
		)

		assert.Equal(t, "otherwise", result)
		assert.False(t, thenCalled)
		assert.True(t, otherwiseCalled)
	})
}

func TestIfFuncE(t *testing.T) {
	t.Run("TrueConditionSuccess", func(t *testing.T) {
		thenCalled := false
		otherwiseCalled := false

		result, err := IfFuncE(true,
			func() (string, error) {
				thenCalled = true
				return "then", nil
			},
			func() (string, error) {
				otherwiseCalled = true
				return "otherwise", nil
			},
		)

		assert.Equal(t, "then", result)
		assert.NoError(t, err)
		assert.True(t, thenCalled)
		assert.False(t, otherwiseCalled)
	})

	t.Run("FalseConditionSuccess", func(t *testing.T) {
		thenCalled := false
		otherwiseCalled := false

		result, err := IfFuncE(false,
			func() (string, error) {
				thenCalled = true
				return "then", nil
			},
			func() (string, error) {
				otherwiseCalled = true
				return "otherwise", nil
			},
		)

		assert.Equal(t, "otherwise", result)
		assert.NoError(t, err)
		assert.False(t, thenCalled)
		assert.True(t, otherwiseCalled)
	})

	t.Run("TrueConditionError", func(t *testing.T) {
		testErr := errors.New("test error")
		thenCalled := false
		otherwiseCalled := false

		result, err := IfFuncE(true,
			func() (string, error) {
				thenCalled = true
				return "", testErr
			},
			func() (string, error) {
				otherwiseCalled = true
				return "otherwise", nil
			},
		)

		assert.Equal(t, "", result)
		assert.Error(t, err)
		assert.Equal(t, testErr, err)
		assert.True(t, thenCalled)
		assert.False(t, otherwiseCalled)
	})

	t.Run("FalseConditionError", func(t *testing.T) {
		testErr := errors.New("test error")
		thenCalled := false
		otherwiseCalled := false

		result, err := IfFuncE(false,
			func() (string, error) {
				thenCalled = true
				return "then", nil
			},
			func() (string, error) {
				otherwiseCalled = true
				return "", testErr
			},
		)

		assert.Equal(t, "", result)
		assert.Error(t, err)
		assert.Equal(t, testErr, err)
		assert.False(t, thenCalled)
		assert.True(t, otherwiseCalled)
	})
}
