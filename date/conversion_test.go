package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestStringToDate ...
func TestStringToDate(t *testing.T) {
	t.Log("string to date current datetime")
	{
		currentTime := time.Now()
		stringDate := currentTime.String()
		res := StringToDate(stringDate)
		assert.Equal(t, res.Format(time.RFC3339), currentTime.Format(time.RFC3339), "value is not equals")
	}
}

// TestStringToUTCDate ...
func TestStringToUTCDate(t *testing.T) {
	t.Log("string to date current datetime")
	{
		currentTime := time.Now().UTC()
		stringDate := currentTime.String()
		res := StringToUTCDate(stringDate)
		assert.Equal(t, res.Format(time.RFC3339), currentTime.Format(time.RFC3339), "value is not equals")
	}
}
