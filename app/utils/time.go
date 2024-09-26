package utils

import "time"

// Create a new time instance at the start of the year for the given date
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}
