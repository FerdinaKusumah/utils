package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestGetCurrentDateString ...
func TestGetCurrentDateString(t *testing.T) {
	t.Log("get current date as string")
	{
		currentTime := time.Now().Format("2006-01-02")
		res := GetCurrentDateString()
		assert.Equal(t, currentTime, res, "value is not equals")
	}
}

// TestGetCurrentDatetime ...
func TestGetCurrentDatetime(t *testing.T) {
	t.Log("get current datetime")
	{
		currentTime := time.Now().Format(time.RFC3339)
		res := GetCurrentDatetime()
		assert.Equal(t, currentTime, res.Format(time.RFC3339), "value is not equals")
	}
}

// TestGetCurrentUTCDatetime ...
func TestGetCurrentUTCDatetime(t *testing.T) {
	t.Log("get current utc datetime")
	{
		currentTime := time.Now().UTC().Format(time.RFC3339)
		res := GetCurrentUTCDatetime()
		assert.Equal(t, currentTime, res.Format(time.RFC3339), "value is not equals")
	}
}

// TestGetCurrentUTCDateString ...
func TestGetCurrentUTCDateString(t *testing.T) {
	t.Log("get current utc datetime string")
	{
		currentTime := time.Now().UTC().Format("2006-01-02")
		res := GetCurrentUTCDateString()
		assert.Equal(t, currentTime, res, "value is not equals")
	}
}

// TestGetCurrentMonthNumber ...
func TestGetCurrentMonthNumber(t *testing.T) {
	t.Log("get current month number")
	{
		currentTime := int(time.Now().Month())
		res := GetCurrentMonthNumber()
		assert.Equal(t, currentTime, res, "value is not equals")
	}
}

// TestGetCurrentMonthString ...
func TestGetCurrentMonthString(t *testing.T) {
	t.Log("get current month number as string")
	{
		currentTime := time.Now().Month().String()
		res := GetCurrentMonthString()
		assert.Equal(t, currentTime, res, "value is not equals")
	}
}

// TestGetCurrentDay ...
func TestGetCurrentDay(t *testing.T) {
	t.Log("get current day")
	{
		currentTime := time.Now().Day()
		res := GetCurrentDay()
		assert.Equal(t, currentTime, res, "value is not equals")
	}
}

// TestGetFirstDateOfMonth ...
func TestGetFirstDateOfMonth(t *testing.T) {
	t.Log("get first date of month")
	{
		currentTime := time.Now()
		expected := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
		res := GetFirstDateOfMonth(currentTime)
		assert.Equal(t, expected, res, "value is not equals")
	}
}

// TestGetFirstDateCurrentMonth ...
func TestGetFirstDateCurrentMonth(t *testing.T) {
	t.Log("get first date in current month")
	{
		currentTime := time.Now()
		expected := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
		res := GetFirstDateCurrentMonth()
		assert.Equal(t, expected, res, "value is not equals")
	}
}

// TestGetLastDateOfMonth ...
func TestGetLastDateOfMonth(t *testing.T) {
	t.Log("get last date of month")
	{
		currentTime := time.Now()
		expected := GetFirstDateOfMonth(currentTime).AddDate(0, 1, 0).Add(-time.Second)
		res := GetLastDateOfMonth(currentTime)
		assert.Equal(t, expected, res, "value is not equals")
	}
}

// TestGetLastDateCurrentMonth ...
func TestGetLastDateCurrentMonth(t *testing.T) {
	t.Log("get last date current month")
	{
		currentTime := time.Now()
		expected := GetFirstDateOfMonth(currentTime).AddDate(0, 1, 0).Add(-time.Second)
		res := GetLastDateCurrentMonth()
		assert.Equal(t, expected, res, "value is not equals")
	}
}

// TestGetFirstDateOfYear ...
func TestGetFirstDateOfYear(t *testing.T) {
	t.Log("get first date if year")
	{
		currentTime := time.Now()
		expected := GetFirstDateOfMonth(currentTime).AddDate(0, 1, 0).Add(-time.Second)
		res := GetFirstDateOfYear(currentTime)
		assert.Equal(t, expected, res, "value is not equals")
	}
}

// TestGetFirstDateCurrentYear ...
func TestGetFirstDateCurrentYear(t *testing.T) {
	t.Log("get first date current year")
	{
		currentTime := time.Now()
		expected := GetFirstDateOfMonth(currentTime).AddDate(0, 1, 0).Add(-time.Second)
		res := GetFirstDateCurrentYear()
		assert.Equal(t, expected, res, "value is not equals")
	}
}

// TestGetLastDateOfYear ...
func TestGetLastDateOfYear(t *testing.T) {
	t.Log("get first date current year")
	{
		currentTime := time.Now()
		expected := GetFirstDateOfYear(currentTime).AddDate(1, 0, 0).Add(-time.Second)
		res := GetLastDateOfYear(currentTime)
		assert.Equal(t, expected, res, "value is not equals")
	}
}

// TestGetLastDateCurrentYear ...
func TestGetLastDateCurrentYear(t *testing.T) {
	t.Log("get first date current year")
	{
		currentTime := time.Now()
		expected := GetFirstDateOfYear(currentTime).AddDate(1, 0, 0).Add(-time.Second)
		res := GetLastDateCurrentYear()
		assert.Equal(t, expected, res, "value is not equals")
	}
}
