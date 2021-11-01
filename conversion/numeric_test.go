package conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToInt(t *testing.T) {

	t.Log("Convert 1 to int")
	{
		num := "1"
		res := StringToInt(num)
		assert.Equal(t, res, 1, "value is not equals")
	}

	t.Log("Convert 1.0 to int")
	{
		num := "1.0"
		res := StringToInt(num)
		assert.Equal(t, res, 0, "value is not equals")
	}

	t.Log("Convert 1,0 to int")
	{
		num := "1,0"
		res := StringToInt(num)
		assert.Equal(t, res, 0, "value is not equals")
	}

	t.Log("Convert 0.1 to int")
	{
		num := "0.1"
		res := StringToInt(num)
		assert.Equal(t, res, 0, "value is not equals")
	}

	t.Log("Convert a to int")
	{
		num := "a"
		res := StringToInt(num)
		assert.Equal(t, res, 0, "value is not equals")
	}
}

func TestStringToInt8(t *testing.T) {

	t.Log("Convert 1 to int8")
	{
		num := "1"
		res := StringToInt8(num)
		assert.Equal(t, res, int8(1), "value is not equals")
	}

	t.Log("Convert 1.0 to int8")
	{
		num := "1.0"
		res := StringToInt8(num)
		assert.Equal(t, res, int8(0), "value is not equals")
	}

	t.Log("Convert 1,0 to int8")
	{
		num := "1,0"
		res := StringToInt8(num)
		assert.Equal(t, res, int8(0), "value is not equals")
	}

	t.Log("Convert 0.1 to int8")
	{
		num := "0.1"
		res := StringToInt8(num)
		assert.Equal(t, res, int8(0), "value is not equals")
	}

	t.Log("Convert a to int8")
	{
		num := "a"
		res := StringToInt8(num)
		assert.Equal(t, res, int8(0), "value is not equals")
	}
}

func TestStringToInt16(t *testing.T) {

	t.Log("Convert 1 to int16")
	{
		num := "1"
		res := StringToInt16(num)
		assert.Equal(t, res, int16(1), "value is not equals")
	}

	t.Log("Convert 1.0 to int16")
	{
		num := "1.0"
		res := StringToInt16(num)
		assert.Equal(t, res, int16(0), "value is not equals")
	}

	t.Log("Convert 1.0 to int16")
	{
		num := "1.0"
		res := StringToInt16(num)
		assert.Equal(t, res, int16(0), "value is not equals")
	}

	t.Log("Convert 0.1 to int16")
	{
		num := "0.1"
		res := StringToInt16(num)
		assert.Equal(t, res, int16(0), "value is not equals")
	}

	t.Log("Convert a to int16")
	{
		num := "a"
		res := StringToInt16(num)
		assert.Equal(t, res, int16(0), "value is not equals")
	}
}

func TestStringToInt32(t *testing.T) {

	t.Log("Convert 1 to int32")
	{
		num := "1"
		res := StringToInt32(num)
		assert.Equal(t, res, int32(1), "value is not equals")
	}

	t.Log("Convert 1.0 to int32")
	{
		num := "1.0"
		res := StringToInt32(num)
		assert.Equal(t, res, int32(0), "value is not equals")
	}

	t.Log("Convert 1,0 to int32")
	{
		num := "1,0"
		res := StringToInt32(num)
		assert.Equal(t, res, int32(0), "value is not equals")
	}

	t.Log("Convert 0.1 to int32")
	{
		num := "0.1"
		res := StringToInt32(num)
		assert.Equal(t, res, int32(0), "value is not equals")
	}

	t.Log("Convert a to int32")
	{
		num := "a"
		res := StringToInt32(num)
		assert.Equal(t, res, int32(0), "value is not equals")
	}
}

func TestStringToInt64(t *testing.T) {

	t.Log("Convert 1 to int64")
	{
		num := "1"
		res := StringToInt64(num)
		assert.Equal(t, res, int64(1), "value is not equals")
	}

	t.Log("Convert 1.0 to int64")
	{
		num := "1.0"
		res := StringToInt64(num)
		assert.Equal(t, res, int64(0), "value is not equals")
	}

	t.Log("Convert 1,0 to int64")
	{
		num := "1,0"
		res := StringToInt64(num)
		assert.Equal(t, res, int64(0), "value is not equals")
	}

	t.Log("Convert 0.1 to int64")
	{
		num := "0.1"
		res := StringToInt64(num)
		assert.Equal(t, res, int64(0), "value is not equals")
	}

	t.Log("Convert a to int64")
	{
		num := "a"
		res := StringToInt64(num)
		assert.Equal(t, res, int64(0), "value is not equals")
	}
}

func TestStringToFloat32(t *testing.T) {

	t.Log("Convert 1 to Float32")
	{
		num := "1"
		res := StringToFloat32(num)
		assert.Equal(t, res, float32(1), "value is not equals")
	}

	t.Log("Convert 1.0 to Float32")
	{
		num := "1.0"
		res := StringToFloat32(num)
		assert.Equal(t, res, float32(1.0), "value is not equals")
	}

	t.Log("Convert 1,0 to Float32")
	{
		num := "1,0"
		res := StringToFloat32(num)
		assert.Equal(t, res, float32(0), "value is not equals")
	}

	t.Log("Convert 0.1 to Float32")
	{
		num := "0.1"
		res := StringToFloat32(num)
		assert.Equal(t, res, float32(0.1), "value is not equals")
	}

	t.Log("Convert a to Float32")
	{
		num := "a"
		res := StringToFloat32(num)
		assert.Equal(t, res, float32(0), "value is not equals")
	}
}

func TestStringToFloat64(t *testing.T) {

	t.Log("Convert 1 to Float64")
	{
		num := "1"
		res := StringToFloat64(num)
		assert.Equal(t, res, float64(1), "value is not equals")
	}

	t.Log("Convert 1.0 to Float64")
	{
		num := "1.0"
		res := StringToFloat64(num)
		assert.Equal(t, res, float64(1.0), "value is not equals")
	}

	t.Log("Convert 1,0 to Float64")
	{
		num := "1,0"
		res := StringToFloat64(num)
		assert.Equal(t, res, float64(0), "value is not equals")
	}

	t.Log("Convert 0.1 to Float64")
	{
		num := "0.1"
		res := StringToFloat64(num)
		assert.Equal(t, res, 0.1, "value is not equals")
	}

	t.Log("Convert 1 to Float64")
	{
		num := "a"
		res := StringToFloat64(num)
		assert.Equal(t, res, float64(0), "value is not equals")
	}
}

func TestStringToBool(t *testing.T) {

	t.Log("Convert True to bool")
	{
		num := "True"
		res := StringToBool(num)
		assert.Equal(t, res, true, "value is not equals")
	}

	t.Log("Convert true to bool")
	{
		num := "true"
		res := StringToBool(num)
		assert.Equal(t, res, true, "value is not equals")
	}

	t.Log("Convert False to bool")
	{
		num := "False"
		res := StringToBool(num)
		assert.Equal(t, res, true, "value is not equals")
	}

	t.Log("Convert false to bool")
	{
		num := "false"
		res := StringToBool(num)
		assert.Equal(t, res, true, "value is not equals")
	}

	t.Log("Convert 1 to bool")
	{
		num := "1"
		res := StringToBool(num)
		assert.Equal(t, res, true, "value is not equals")
	}

	t.Log("Convert 1.0 to bool")
	{
		num := "1.0"
		res := StringToBool(num)
		assert.Equal(t, res, false, "value is not equals")
	}

	t.Log("Convert 1,0 to bool")
	{
		num := "1,0"
		res := StringToBool(num)
		assert.Equal(t, res, false, "value is not equals")
	}

	t.Log("Convert 0.1 to bool")
	{
		num := "0.1"
		res := StringToBool(num)
		assert.Equal(t, res, false, "value is not equals")
	}

	t.Log("Convert a to bool")
	{
		num := "a"
		res := StringToBool(num)
		assert.Equal(t, res, false, "value is not equals")
	}
}
