package date

import "github.com/araddon/dateparse"

// IsDateIsValid ...
func IsDateIsValid(text string) bool {
	_, err := dateparse.ParseAny(text)
	return err == nil
}
