package conversion

import (
	"fmt"
	"html"
	"strconv"
)

// IntToString ...
func IntToString(num int) string {
	return strconv.Itoa(num)
}

// Int8ToString ...
func Int8ToString(num int8) string {
	return IntToString(int(num))
}

// Int16ToString ...
func Int16ToString(num int16) string {
	return IntToString(int(num))
}

// Int32ToString ...
func Int32ToString(num int32) string {
	return IntToString(int(num))
}

// Int64ToString ...
func Int64ToString(num int64) string {
	return IntToString(int(num))
}

// Float32ToString  ...
func Float32ToString(num float32) string {
	return fmt.Sprintf(`%f`, num)
}

// Float64ToString  ...
func Float64ToString(num float64) string {
	return fmt.Sprintf(`%f`, num)
}

// BoolToString ...
func BoolToString(num bool) string {
	return fmt.Sprintf(`%t`, num)
}

// EscapeString ...
func EscapeString(text string) string {
	return html.EscapeString(text)
}

// UnescapeString ...
func UnescapeString(text string) string {
	return html.UnescapeString(text)
}
