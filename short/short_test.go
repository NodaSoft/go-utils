package short

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
