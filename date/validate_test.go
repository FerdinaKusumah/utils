package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDateIsValid(t *testing.T) {
	t.Log("test is valid datetime")
	{
		res := IsDateIsValid("May 8, 2009 5:57:51 PM")
		assert.Equal(t, res, true, "value is not equals")
	}
}
