package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

// Sha1HashString ...
func Sha1HashString(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
