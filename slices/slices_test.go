package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterNil(t *testing.T) {
	r := FilterNil([]uint{0, 0, 1, 2, 0, 0, 0, 3, 4, 0, 5, 0, 6, 0, 7, 0, 0, 0, 8, 9, 0, 0, 0})

	assert.Equal(t, []uint{1, 2, 3, 4, 5, 6, 7, 8, 9}, r)
}

func TestUnique(t *testing.T) {
	r := Unique([]uint{0, 1, 2, 2, 3, 1, 2, 4, 5, 0, 6, 4, 4, 5, 7, 8, 0, 3, 4, 5, 6, 0, 5, 4, 0, 0, 4, 0, 3, 0, 4, 3, 0, 0, 0, 9, 9})

	assert.Equal(t, []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, r)
}

func TestMax(t *testing.T) {
	assert.Equal(t, 28, Max([]int{-2, 3, 15, 28, 4}))
	assert.Equal(t, 28.1, Max([]float64{-2.2, 3.2, 15, 28.1, 4.4}))

	uintMax := Max([]uint{2, 3, 15, 28, 4})
	var uintExp uint = 28
	assert.Equal(t, uintExp, uintMax)

	int8Val := Max([]int8{-2, 3, 15, 28, 4})
	var expInt8 int8 = 28
	assert.Equal(t, expInt8, int8Val)
}

func TestMin(t *testing.T) {
	assert.Equal(t, -2, Min([]int{1, -2, 3, 15, 28, 4}))
	assert.Equal(t, -2.2, Min([]float64{1.1, -2.2, 3.2, 15, 28.1, 4.4}))

	minUint := Min([]uint{4, 2, 3, 15, 28, 4})
	var expUint uint = 2
	assert.Equal(t, expUint, minUint)

	minInt8 := Min([]int8{4, -2, 3, 15, 28, 4})
	var expInt8 int8 = -2
	assert.Equal(t, expInt8, minInt8)
}

func TestSum(t *testing.T) {
	assert.Equal(t, 49, Sum([]int{1, -2, 3, 15, 28, 4}))
	assert.Equal(t, 49.6, Sum([]float64{1.1, -2.2, 3.2, 15, 28.1, 4.4}))

	sumUint := Sum([]uint{4, 2, 3, 15, 28, 4})
	var uintExp uint = 56
	assert.Equal(t, uintExp, sumUint)

	sumInt8 := Sum([]int8{4, -2, 3, 15, 28, 4})
	var expInt8 int8 = 52
	assert.Equal(t, expInt8, sumInt8)
}
