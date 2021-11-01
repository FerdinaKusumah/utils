package conversion

import "strconv"

// StringToInt ...
func StringToInt(num string) int {
	r, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	return r
}

// StringToInt8 ...
func StringToInt8(num string) int8 {
	return int8(StringToInt(num))
}

// StringToInt16 ...
func StringToInt16(num string) int16 {
	return int16(StringToInt(num))
}

// StringToInt32 ...
func StringToInt32(num string) int32 {
	return int32(StringToInt(num))
}

// StringToInt64 ...
func StringToInt64(num string) int64 {
	return int64(StringToInt(num))
}

// StringToFloat32  ...
func StringToFloat32(num string) float32 {
	return float32(StringToFloat64(num))
}

// StringToFloat64  ...
func StringToFloat64(num string) float64 {
	r, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0.0
	}
	return r
}

// StringToBool ...
func StringToBool(num string) bool {
	_, err := strconv.ParseBool(num)
	return err == nil
}
