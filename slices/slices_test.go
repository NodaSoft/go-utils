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

func TestUnion(t *testing.T) {
	r := Union([]uint{1, 2, 3, 4}, []uint{3, 4, 5})
	assert.ElementsMatch(t, []uint{1, 2, 3, 4, 5}, r)
}

func TestCross(t *testing.T) {
	r := Cross([]uint{1, 2, 3}, []uint{2, 3, 4})
	assert.Equal(t, []uint{2, 3}, r)
}

func TestIsEqual(t *testing.T) {
	assert.True(t, IsEqual([]uint{1, 2, 3}, []uint{3, 2, 1}))
	assert.False(t, IsEqual([]uint{1, 2}, []uint{1, 2, 3}))
}

func TestHas(t *testing.T) {
	assert.True(t, Has([]uint{1, 2, 3}, 2))
	assert.False(t, Has([]uint{1, 2, 3}, 4))
}

func TestTrimStrings(t *testing.T) {
	r := TrimStrings([]string{"  hello  ", " world  ", "  ! "})
	assert.Equal(t, []string{"hello", "world", "!"}, r)
}

func TestToKeyMap(t *testing.T) {
	r := ToKeyMap([]string{"a", "b", "a", "c"})
	expected := map[string]bool{"a": true, "b": true, "c": true}
	assert.Equal(t, expected, r)
}

func TestSliceDiff(t *testing.T) {
	r := SliceDiff([]uint{1, 2, 3, 4}, []uint{2, 3})
	assert.Equal(t, []uint{1, 4}, r)
}

func TestSliceIntersect(t *testing.T) {
	r := SliceIntersect([]uint{1, 2, 3}, []uint{2, 3}, []uint{3, 4})
	assert.Equal(t, []uint{3}, r)
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

func BenchmarkFilterNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FilterNil([]uint{0, 0, 1, 2, 0, 3, 4, 0, 5})
	}
}

func BenchmarkUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unique([]uint{0, 1, 2, 2, 3, 1, 4, 5, 0, 6})
	}
}

func BenchmarkUnion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Union([]uint{1, 2, 3}, []uint{3, 4, 5})
	}
}

func BenchmarkCross(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cross([]uint{1, 2, 3}, []uint{2, 3, 4})
	}
}

func BenchmarkIsEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEqual([]uint{1, 2, 3}, []uint{3, 2, 1})
	}
}

func BenchmarkHas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Has([]uint{1, 2, 3}, 2)
	}
}

func BenchmarkTrimStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimStrings([]string{"  hello  ", " world  ", "  ! "})
	}
}

func BenchmarkToKeyMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToKeyMap([]string{"a", "b", "a", "c"})
	}
}

func BenchmarkSliceDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceDiff([]uint{1, 2, 3, 4}, []uint{2, 3})
	}
}

func BenchmarkSliceIntersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceIntersect([]uint{1, 2, 3}, []uint{2, 3}, []uint{3, 4})
	}
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max([]int{-2, 3, 15, 28, 4})
	}
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min([]int{1, -2, 3, 15, 28, 4})
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, -2, 3, 15, 28, 4})
	}
}

func TestStringsToType(t *testing.T) {
	// uint
	stringsUint := []string{"0", "1", "3", "15", "2006"}
	expectedUints := []uint{0, 1, 3, 15, 2006}
	uints, err := StringsToUints[uint](stringsUint)
	assert.NoError(t, err)
	assert.Equal(t, expectedUints, uints)

	// uint16
	stringsUint16 := []string{"0", "1", "3", "15", "2006"}
	expectedUint16s := []uint16{0, 1, 3, 15, 2006}

	uint16s, err := StringsToUints[uint16](stringsUint16)
	assert.NoError(t, err)
	assert.Equal(t, expectedUint16s, uint16s)

	// error
	stringsWithNegative := []string{"0", "-1", "3", "-15", "2006"}
	_, err = StringsToUints[uint16](stringsWithNegative)
	assert.NotNil(t, err)
}
