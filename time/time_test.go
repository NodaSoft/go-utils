package time

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMidnight(t *testing.T) {
	midnight, err := Midnight()
	assert.NoError(t, err)

	assert.Equal(t, 0, midnight.Hour())
	assert.Equal(t, 0, midnight.Minute())
	assert.Equal(t, 0, midnight.Second())
	assert.Equal(t, 0, midnight.Nanosecond())
}

func TestMidnightByLocation(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	midnight, err := MidnightByLocation(loc)
	assert.NoError(t, err)

	assert.Equal(t, 0, midnight.Hour())
	assert.Equal(t, 0, midnight.Minute())
	assert.Equal(t, 0, midnight.Second())
	assert.Equal(t, 0, midnight.Nanosecond())
	assert.Equal(t, loc, midnight.Location())
}

func TestMidnightByTimeZone(t *testing.T) {
	midnight, err := MidnightByTimeZone("Europe/Moscow")
	assert.NoError(t, err)

	assert.Equal(t, 0, midnight.Hour())
	assert.Equal(t, 0, midnight.Minute())
	assert.Equal(t, 0, midnight.Second())
	assert.Equal(t, 0, midnight.Nanosecond())
	assert.Equal(t, "Europe/Moscow", midnight.Location().String())

	_, err = MidnightByTimeZone("Invalid/Timezone")
	assert.Error(t, err)
}

func BenchmarkMidnight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Midnight()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkMidnightByLocation(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	for i := 0; i < b.N; i++ {
		_, err := MidnightByLocation(loc)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkMidnightByTimeZone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MidnightByTimeZone("Asia/Tokyo")
		if err != nil {
			log.Fatal(err)
		}
	}
}
