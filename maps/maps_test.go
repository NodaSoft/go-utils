package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHas(t *testing.T) {
	// Проверка существующего ключа
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	assert.True(t, Has(m, 1))
	assert.True(t, Has(m, 2))

	// Проверка отсутствующего ключа
	assert.False(t, Has(m, 4))

	// Проверка пустой карты
	emptyMap := make(map[int]string)
	assert.False(t, Has(emptyMap, 1))
}

func TestMerge(t *testing.T) {
	// Две непересекающиеся карты
	a := map[int]string{1: "one", 2: "two"}
	b := map[int]string{3: "three", 4: "four"}
	expected := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
	result := Merge(a, b)
	assert.Equal(t, expected, result)

	// Пересекающиеся ключи, значения карты a должны быть приоритетными
	a = map[int]string{1: "uno", 2: "dos"}
	b = map[int]string{1: "one", 3: "three"}
	expected = map[int]string{1: "uno", 2: "dos", 3: "three"}
	result = Merge(a, b)
	assert.Equal(t, expected, result)

	// Пустая карта b
	a = map[int]string{1: "one", 2: "two"}
	b = map[int]string{}
	expected = map[int]string{1: "one", 2: "two"}
	result = Merge(a, b)
	assert.Equal(t, expected, result)

	// Пустая карта a
	a = map[int]string{}
	b = map[int]string{3: "three", 4: "four"}
	expected = map[int]string{3: "three", 4: "four"}
	result = Merge(a, b)
	assert.Equal(t, expected, result)
}

func TestDiffKeys(t *testing.T) {
	// a содержит ключи, которых нет в b
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{2: "two", 4: "four"}
	expected := map[int]string{1: "one", 3: "three"}
	result := DiffKeys(a, b)
	assert.Equal(t, expected, result)

	// Все ключи из a есть в b
	a = map[int]string{1: "one", 2: "two"}
	b = map[int]string{1: "one", 2: "two"}
	expected = map[int]string{}
	result = DiffKeys(a, b)
	assert.Equal(t, expected, result)

	// Пустая карта b
	a = map[int]string{1: "one", 2: "two"}
	b = map[int]string{}
	expected = map[int]string{1: "one", 2: "two"}
	result = DiffKeys(a, b)
	assert.Equal(t, expected, result)

	// Пустая карта a
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
