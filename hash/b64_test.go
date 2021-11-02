package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestBase64Encode ...
func TestBase64Encode(t *testing.T) {
	t.Log(`test encode string "abcd"`)
	{
		text := "abcd"
		res := Base64Encode(text)
		assert.Equal(t, res, "YWJjZA==", "value is not equals")
	}
}

// TestBase64Decode ...
func TestBase64Decode(t *testing.T) {
	t.Log(`test decode string "abcd"`)
	{
		text := "YWJjZA=="
		res := Base64Decode(text)
		assert.Equal(t, res, "abcd", "value is not equals")
	}
}
