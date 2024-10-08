package time

import (
	"time"
)

// Midnight returns the time at midnight (00:00:00) for the current day.
func Midnight() (time.Time, error) {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return midnight, nil
}

// MidnightByLocation returns the time at midnight in the given location.
func MidnightByLocation(loc *time.Location) (time.Time, error) {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	return midnight, nil
}

// MidnightByTimeZone returns the time at midnight in the specified time zone.
func MidnightByTimeZone(timeZone string) (time.Time, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return time.Time{}, err
	}

	return MidnightByLocation(loc)
}
