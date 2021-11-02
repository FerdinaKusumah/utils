package date

import (
	"github.com/araddon/dateparse"
	"time"
)

// StringToDate ...
func StringToDate(text string) time.Time {
	t, err := dateparse.ParseAny(text)
	if err != nil || t.IsZero() {
		return time.Time{}
	}
	return t
}

// StringToUTCDate ...
func StringToUTCDate(text string) time.Time {
	t := StringToDate(text)
	if t.IsZero() {
		return time.Time{}
	}
	return t.UTC()
}
