package maps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrderedMap_Len(t *testing.T) {
	// Map is nil
	var m *OrderedMap[int, int]
	var mFast *OrderedFastDeleteMap[int, int]
	assert.Equal(t, 0, m.Len())
	assert.Panics(t, func() { mFast.Len() })

	// Map is initialized without constructor
	m = &OrderedMap[int, int]{}
	mFast = &OrderedFastDeleteMap[int, int]{}
	assert.Equal(t, 0, m.Len())
	assert.Equal(t, 0, mFast.Len())

	// Empty map
	m = NewOrderedMap[int, int](3)
	mFast = NewOrderedFastDeleteMap[int, int](3)
	assert.Equal(t, 0, m.Len())
	assert.Equal(t, 0, mFast.Len())

	// After insert new value
	m.Set(100, 200)
	mFast.Set(100, 200)
	assert.Equal(t, 1, m.Len())
	assert.Equal(t, 1, mFast.Len())

	// After delete value
	m.Delete(100)
	mFast.Delete(100)
	assert.Equal(t, 0, m.Len())
	assert.Equal(t, 0, mFast.Len())
}

func TestOrderedMap_Has(t *testing.T) {
	// Map is nil
	var m *OrderedMap[int, string]
	var mFast *OrderedFastDeleteMap[int, string]
	assert.False(t, m.Has(1))
	assert.Panics(t, func() { mFast.Has(1) })

	// Map is initialized without constructor
	m = &OrderedMap[int, string]{}
	mFast = &OrderedFastDeleteMap[int, string]{}
	assert.False(t, m.Has(1))
	assert.False(t, mFast.Has(1))

	// Empty map
	m = NewOrderedMap[int, string](3)
	mFast = NewOrderedFastDeleteMap[int, string](3)
	assert.False(t, m.Has(1))
	assert.False(t, mFast.Has(1))

	m.Set(3, "3")
	m.Set(-300, "some string")
	mFast.Set(3, "3")
	mFast.Set(-300, "some string")

	// Check for non-existing key
	assert.False(t, m.Has(4))
	assert.False(t, mFast.Has(4))

	// Check for an existing key
	assert.True(t, m.Has(-300))
	assert.True(t, mFast.Has(-300))
}

func TestOrderedMap_GetValue(t *testing.T) {
	// Map is nil
	var m *OrderedMap[int, string]
	var mFast *OrderedFastDeleteMap[int, string]
	assert.Equal(t, "", m.GetValue(1))
	assert.Panics(t, func() { mFast.GetValue(1) })

	// Map is initialized without constructor
	m = &OrderedMap[int, string]{}
	mFast = &OrderedFastDeleteMap[int, string]{}
	assert.Equal(t, "", m.GetValue(1))
	assert.Equal(t, "", mFast.GetValue(1))

	// Get value from empty map
	m = NewOrderedMap[int, string](3)
	mFast = NewOrderedFastDeleteMap[int, string](3)
	assert.Equal(t, "", m.GetValue(1))
	assert.Equal(t, "", mFast.GetValue(1))

	m.Set(3, "three")
	m.Set(-300, "some string")
	mFast.Set(3, "3")
	mFast.Set(-300, "some string")

	// Get value by non-existent key
	assert.Equal(t, "", m.GetValue(1))
	assert.Equal(t, "", mFast.GetValue(1))

	// Get value by existent key
	assert.Equal(t, "some string", m.GetValue(-300))
	assert.Equal(t, "some string", mFast.GetValue(-300))
}

func TestOrderedMap_GetAndCheck(t *testing.T) {
	// Map is nil
	var m *OrderedMap[int, string]
	var mFast *OrderedFastDeleteMap[int, string]
	v, ok := m.GetAndCheck(1)
	assert.Equal(t, "", v)
	assert.False(t, ok)
	assert.Panics(t, func() { mFast.GetAndCheck(1) })

	// Map is initialized without constructor
	m = &OrderedMap[int, string]{}
	mFast = &OrderedFastDeleteMap[int, string]{}
	v, ok = m.GetAndCheck(1)
	assert.Equal(t, "", v)
	assert.False(t, ok)
	v, ok = mFast.GetAndCheck(1)
	assert.Equal(t, "", v)
	assert.False(t, ok)

	// Get from empty map
	m = NewOrderedMap[int, string](3)
	mFast = NewOrderedFastDeleteMap[int, string](3)
	v, ok = m.GetAndCheck(1)
	assert.Equal(t, "", v)
	assert.False(t, ok)
	v, ok = mFast.GetAndCheck(1)
	assert.Equal(t, "", v)
	assert.False(t, ok)

	// Get by non-existent key
	m.Set(3, "three")
	m.Set(-300, "some string")
	mFast.Set(3, "three")
	mFast.Set(-300, "some string")
	v, ok = m.GetAndCheck(1)
	assert.Equal(t, "", v)
	assert.False(t, ok)
	v, ok = mFast.GetAndCheck(1)
	assert.Equal(t, "", v)
	assert.False(t, ok)

	// Get by existent key
	v, ok = m.GetAndCheck(3)
	assert.Equal(t, "three", v)
	assert.True(t, ok)
	v, ok = mFast.GetAndCheck(3)
	assert.Equal(t, "three", v)
	assert.True(t, ok)
}

func TestOrderedMap_Set(t *testing.T) {
	// Panics if map is nil
	var m *OrderedMap[int, string]
	var mFast *OrderedFastDeleteMap[int, string]
	assert.Panics(t, func() { m.Set(1, "1") })
	assert.Panics(t, func() { mFast.Set(1, "1") })

	// Panics if map is not initialized by constructor
	m = &OrderedMap[int, string]{}
	mFast = &OrderedFastDeleteMap[int, string]{}
	assert.Panics(t, func() { m.Set(-300, "some string") })
	assert.Panics(t, func() { mFast.Set(-300, "some string") })

	// Pap is initialized by constructor
	m = NewOrderedMap[int, string](3)
	mFast = NewOrderedFastDeleteMap[int, string](3)
	m.Set(-300, "some string")
	mFast.Set(-300, "some string")
	assert.Equal(t, "some string", m.GetValue(-300))
	assert.Equal(t, "some string", mFast.GetValue(-300))
}

func TestOrderedMap_Delete(t *testing.T) {
	// Does not panic if map nil
	var m *OrderedMap[int, string]
	var mFast *OrderedFastDeleteMap[int, string]
	assert.NotPanics(t, func() { m.Delete(1) })
	assert.NotPanics(t, func() { mFast.Delete(1) })

	// Does not panic if map is not initialized by constructor
	m = &OrderedMap[int, string]{}
	mFast = &OrderedFastDeleteMap[int, string]{}
	assert.NotPanics(t, func() { m.Delete(1) })
	assert.NotPanics(t, func() { mFast.Delete(1) })

	m = NewOrderedMap[int, string](3)
	mFast = NewOrderedFastDeleteMap[int, string](3)
	m.Set(3, "three")
	m.Set(-300, "some string")
	length := m.Len()
	mFast.Set(3, "three")
	mFast.Set(-300, "some string")
	lengthFast := mFast.Len()

	// Delete value by non-existent key
	value := m.GetValue(8)
	m.Delete(8)
	assert.True(t, value == "" && m.GetValue(8) == "")
	assert.True(t, length == m.Len())

	valueFast := mFast.GetValue(8)
	mFast.Delete(8)
	assert.True(t, valueFast == "" && mFast.GetValue(8) == "")
	assert.True(t, lengthFast == mFast.Len())

	// Delete value by existent key
	value = m.GetValue(3)
	m.Delete(3)
	assert.True(t, value != "" && m.GetValue(3) == "")
	assert.True(t, length == m.Len()+1)

	valueFast = mFast.GetValue(3)
	mFast.Delete(3)
	assert.True(t, valueFast != "" && mFast.GetValue(3) == "")
	assert.True(t, lengthFast == mFast.Len()+1)
}

func TestOrderedMap_Iterate(t *testing.T) {
	// Map nil object
	var m *OrderedMap[int, int]
	var mFast *OrderedFastDeleteMap[int, int]
	for k, v := range m.Iterate() {
		require.Fail(t, fmt.Sprintf("iterate over nil object, key=%d, value='%d'", k, v))
	}
	assert.Panics(t, func() { mFast.Iterate() })

	// Map is not initialized by constructor
	m = &OrderedMap[int, int]{}
	mFast = &OrderedFastDeleteMap[int, int]{}
	for k, v := range mFast.Iterate() {
		require.Fail(t, fmt.Sprintf("iterate over nil map, key=%d, value='%d'", k, v))
	}

	// Initialize and check iteration result
	digits := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9}
	m = NewOrderedMap[int, int](len(digits))
	mFast = NewOrderedFastDeleteMap[int, int](len(digits))
	for k, v := range digits {
		m.Set(k, v)
		mFast.Set(k, v)
	}

	result := make([]int, 0, len(digits))
	for k, v := range m.Iterate() {
		if k == len(result) {
			result = append(result, v)
			continue
		}

		require.Equal(t, len(result), k)
	}
	assert.Equal(t, digits, result)

	resultFast := make([]int, 0, len(digits))
	for k, v := range mFast.Iterate() {
		if k == len(resultFast) {
			resultFast = append(resultFast, v)
			continue
		}

		require.Equal(t, len(resultFast), k)
	}
	assert.Equal(t, digits, resultFast)
}

func BenchmarkOrderedMap(b *testing.B) {
	size := 1000
	m := NewOrderedMap[int, int](size)

	for i := 0; i < b.N; i++ {
		for i := 0; i < size; i++ {
			m.Set(i, i)
		}

		for j := size - 1; j >= 0; j-- {
			m.Delete(j)
		}
	}
}

func BenchmarkOrderedFastDeleteMap(b *testing.B) {
	size := 1000
	m := NewOrderedFastDeleteMap[int, int](size)

	for i := 0; i < b.N; i++ {
		for i := 0; i < size; i++ {
			m.Set(i, i)
		}

		for j := size - 1; j >= 0; j-- {
			m.Delete(j)
		}
	}
}
