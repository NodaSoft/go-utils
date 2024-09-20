package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 28, Max(-2, 3, 15, 28, 4))
	assert.Equal(t, 28.1, Max(-2.2, 3.2, 15, 28.1, 4.4))

	uintMax := Max([]uint{2, 3, 15, 28, 4}...)
	var uintExp uint = 28
	assert.Equal(t, uintExp, uintMax)

	int8Val := Max([]int8{-2, 3, 15, 28, 4}...)
	var expInt8 int8 = 28
	assert.Equal(t, expInt8, int8Val)
}

func TestMin(t *testing.T) {
	assert.Equal(t, -2, Min(1, -2, 3, 15, 28, 4))
	assert.Equal(t, -2.2, Min(1.1, -2.2, 3.2, 15, 28.1, 4.4))

	minUint := Min([]uint{4, 2, 3, 15, 28, 4}...)
	var expUint uint = 2
	assert.Equal(t, expUint, minUint)

	minInt8 := Min([]int8{4, -2, 3, 15, 28, 4}...)
	var expInt8 int8 = -2
	assert.Equal(t, expInt8, minInt8)
}

func TestSum(t *testing.T) {
	assert.Equal(t, 49, Sum(1, -2, 3, 15, 28, 4))
	assert.Equal(t, 49.6, Sum(1.1, -2.2, 3.2, 15, 28.1, 4.4))

	sumUint := Sum([]uint{4, 2, 3, 15, 28, 4}...)
	var uintExp uint = 56
	assert.Equal(t, uintExp, sumUint)

	sumInt8 := Sum([]int8{4, -2, 3, 15, 28, 4}...)
	var expInt8 int8 = 52
	assert.Equal(t, expInt8, sumInt8)
}

func BenchmarkMax(b *testing.B) {
	ints := []int{-2, 3, 15, 28, 4, 100, 99, 42}
	floats := []float64{-2.2, 3.2, 15.5, 28.1, 4.4, 100.8, 99.9, 42.2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(ints...)
		Max(floats...)
	}
}

func BenchmarkMin(b *testing.B) {
	ints := []int{1, -2, 3, 15, 28, 4, -100, 42}
	floats := []float64{1.1, -2.2, 3.2, 15.8, 28.1, 4.4, -100.7, 42.3}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(ints...)
		Min(floats...)
	}
}

func BenchmarkSum(b *testing.B) {
	ints := []int{1, -2, 3, 15, 28, 4, 100, 42, 10, 99}
	floats := []float64{1.1, -2.2, 3.2, 15.8, 28.1, 4.4, 100.7, 42.3, 99.9}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(ints...)
		Sum(floats...)
	}
}
