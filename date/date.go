package date

import (
	"time"
)

// GetCurrentDateString ...
func GetCurrentDateString() string {
	return time.Now().Format("2006-01-02")
}

// GetCurrentDatetime ...
func GetCurrentDatetime() time.Time {
	return time.Now().In(time.Local)
}

// GetCurrentUTCDatetime ...
func GetCurrentUTCDatetime() time.Time {
	return time.Now().In(time.UTC)
}

// GetCurrentUTCDateString ...
func GetCurrentUTCDateString() string {
	return time.Now().In(time.UTC).Format("2006-01-02")
}

// GetCurrentMonthNumber ...
func GetCurrentMonthNumber() int {
	return int(time.Now().Month())
}

// GetCurrentMonthString ...
func GetCurrentMonthString() string {
	return time.Now().Month().String()
}

// GetCurrentDay ...
func GetCurrentDay() int {
	return time.Now().Day()
}

// GetFirstDateOfMonth ...
func GetFirstDateOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// GetFirstDateCurrentMonth ...
func GetFirstDateCurrentMonth() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// GetLastDateOfMonth ...
func GetLastDateOfMonth(t time.Time) time.Time {
	return GetFirstDateOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

// GetLastDateCurrentMonth ...
func GetLastDateCurrentMonth() time.Time {
	t := time.Now()
	return GetFirstDateOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

// GetFirstDateOfYear ...
func GetFirstDateOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// GetFirstDateCurrentYear ...
func GetFirstDateCurrentYear() time.Time {
	t := time.Now()
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// GetLastDateOfYear ...
func GetLastDateOfYear(t time.Time) time.Time {
	return GetFirstDateOfYear(t).AddDate(1, 0, 0).Add(-time.Second)
}

// GetLastDateCurrentYear ...
func GetLastDateCurrentYear() time.Time {
	t := time.Now()
	return GetFirstDateOfYear(t).AddDate(1, 0, 0).Add(-time.Second)
}
