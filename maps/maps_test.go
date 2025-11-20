package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHas(t *testing.T) {
	// Check for an existing key
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	assert.True(t, Has(m, 1))
	assert.True(t, Has(m, 2))

	// Check for a non-existent key
	assert.False(t, Has(m, 4))

	// Check for an empty map
	emptyMap := make(map[int]string)
	assert.False(t, Has(emptyMap, 1))
}

func TestMerge(t *testing.T) {
	// Two non-overlapping maps
	a := map[int]string{1: "one", 2: "two"}
	b := map[int]string{3: "three", 4: "four"}
	expected := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
	result := Merge(a, b)
	assert.Equal(t, expected, result)

	// Overlapping keys, values from map 'a' should have priority
	a = map[int]string{1: "uno", 2: "dos"}
	b = map[int]string{1: "one", 3: "three"}
	expected = map[int]string{1: "uno", 2: "dos", 3: "three"}
	result = Merge(a, b)
	assert.Equal(t, expected, result)

	// Empty map 'b'
	a = map[int]string{1: "one", 2: "two"}
	b = map[int]string{}
	expected = map[int]string{1: "one", 2: "two"}
	result = Merge(a, b)
	assert.Equal(t, expected, result)

	// Empty map 'a'
	a = map[int]string{}
	b = map[int]string{3: "three", 4: "four"}
	expected = map[int]string{3: "three", 4: "four"}
	result = Merge(a, b)
	assert.Equal(t, expected, result)
}

func TestDiffKeys(t *testing.T) {
	// 'a' contains keys that are not in 'b'
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{2: "two", 4: "four"}
	expected := map[int]string{1: "one", 3: "three"}
	result := DiffKeys(a, b)
	assert.Equal(t, expected, result)

	// All keys from 'a' are in 'b'
	a = map[int]string{1: "one", 2: "two"}
	b = map[int]string{1: "one", 2: "two"}
	expected = map[int]string{}
	result = DiffKeys(a, b)
	assert.Equal(t, expected, result)

	// Empty map 'b'
	a = map[int]string{1: "one", 2: "two"}
	b = map[int]string{}
	expected = map[int]string{1: "one", 2: "two"}
	result = DiffKeys(a, b)
	assert.Equal(t, expected, result)

	// Empty map 'a'
	a = map[int]string{}
	b = map[int]string{1: "one", 2: "two"}
	expected = map[int]string{}
	result = DiffKeys(a, b)
	assert.Equal(t, expected, result)
}

func BenchmarkHas(b *testing.B) {
	m := make(map[int]string, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = "value"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Has(m, i%1000)
	}
}

func BenchmarkMerge(b *testing.B) {
	a := make(map[int]string, 1000)
	bm := make(map[int]string, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = "a_value"
		bm[i] = "b_value"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Merge(a, bm)
	}
}

func BenchmarkDiffKeys(b *testing.B) {
	a := make(map[int]string, 1000)
	bm := make(map[int]string, 500)
	for i := 0; i < 1000; i++ {
		a[i] = "value"
		if i < 500 {
			bm[i] = "value"
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffKeys(a, bm)
	}
}

func TestFilterByKeys(t *testing.T) {
	m := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}

	keys := []int{1, 3, 5}
	expected := map[int]string{1: "one", 3: "three", 5: "five"}
	result := FilterByKeys(m, keys)
	assert.Equal(t, expected, result)

	keys = []int{2, 4}
	expected = map[int]string{2: "two", 4: "four"}
	result = FilterByKeys(m, keys)
	assert.Equal(t, expected, result)

	keys = []int{10, 20}
	expected = map[int]string{}
	result = FilterByKeys(m, keys)
	assert.Equal(t, expected, result)

	keys = []int{}
	expected = map[int]string{}
	result = FilterByKeys(m, keys)
	assert.Equal(t, expected, result)
}

func BenchmarkFilterByKeys(b *testing.B) {
	m := make(map[int]string, 1000)
	keys := make([]int, 100)
	for i := 0; i < 1000; i++ {
		m[i] = "value"
		if i < 100 {
			keys[i] = i
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterByKeys(m, keys)
	}
}
