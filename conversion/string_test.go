package conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToString(t *testing.T) {

	t.Log(`Convert 1 to string`)
	{
		num := 1
		res := IntToString(num)
		assert.Equal(t, res, "1", "value is not equals")
	}
}

func TestInt8ToString(t *testing.T) {

	t.Log(`Convert 1 to string`)
	{
		num := int8(1)
		res := Int8ToString(num)
		assert.Equal(t, res, "1", "value is not equals")
	}
}

func TestInt16ToString(t *testing.T) {

	t.Log(`Convert 1 to string`)
	{
		num := int16(1)
		res := Int16ToString(num)
		assert.Equal(t, res, "1", "value is not equals")
	}
}

func TestInt32ToString(t *testing.T) {

	t.Log(`Convert 1 to string`)
	{
		num := int32(1)
		res := Int32ToString(num)
		assert.Equal(t, res, "1", "value is not equals")
	}
}

func TestInt64ToString(t *testing.T) {

	t.Log(`Convert 1 to string`)
	{
		num := int64(1)
		res := Int64ToString(num)
		assert.Equal(t, res, "1", "value is not equals")
	}
}

func TestFloat32ToString(t *testing.T) {

	t.Log(`Convert 1.0 to string`)
	{
		num := float32(1.0)
		res := Float32ToString(num)
		assert.Equal(t, res, "1.000000", "value is not equals")
	}
}

func TestFloat64ToString(t *testing.T) {

	t.Log(`Convert 1.0 to string`)
	{
		num := 1.0
		res := Float64ToString(num)
		assert.Equal(t, res, "1.000000", "value is not equals")
	}
}

func TestBoolToString(t *testing.T) {

	t.Log(`Convert true to string`)
	{
		num := true
		res := BoolToString(num)
		assert.Equal(t, res, "true", "value is not equals")
	}

	t.Log(`Convert false to string`)
	{
		num := false
		res := BoolToString(num)
		assert.Equal(t, res, "false", "value is not equals")
	}
}

func TestEscapeString(t *testing.T) {

	t.Log(`Convert escape string "Fran & Freddie's Diner" <tasty@example.com>"`)
	{
		text := `"Fran & Freddie's Diner" <tasty@example.com>"`
		res := EscapeString(text)
		assert.Equal(t, res, "&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;&#34;", "value is not equals")
	}

	t.Log(`Convert escape string "your_name@gmail.com && my_name@gmail.com"`)
	{
		text := "your_name@gmail.com && my_name@gmail.com"
		res := EscapeString(text)
		assert.Equal(t, res, "your_name@gmail.com &amp;&amp; my_name@gmail.com", "value is not equals")
	}
}

func TestUnescapeString(t *testing.T) {

	t.Log(`Convert escape string "Fran & Freddie's Diner" <tasty@example.com>"`)
	{
		text := `&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;&#34;`
		res := UnescapeString(text)
		assert.Equal(t, res, `"Fran & Freddie's Diner" <tasty@example.com>"`, "value is not equals")
	}

	t.Log(`Convert escape string "your_name@gmail.com && my_name@gmail.com"`)
	{
		text := "your_name@gmail.com &amp;&amp; my_name@gmail.com"
		res := UnescapeString(text)
		assert.Equal(t, res, "your_name@gmail.com && my_name@gmail.com", "value is not equals")
	}
}
