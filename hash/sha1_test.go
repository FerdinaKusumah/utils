package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha1HashString(t *testing.T) {
	t.Log(`Has string "abcd"`)
	{
		text := "abcd"
		res := Sha1HashString(text)
		assert.Equal(t, res, "81fe8bfe87576c3ecb22426f8e57847382917acf", "value is not equals")
	}
}
