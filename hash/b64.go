package hash

import (
	b64 "encoding/base64"
)

// Base64Encode ...
func Base64Encode(text string) string {
	return b64.StdEncoding.EncodeToString([]byte(text))
}

// Base64Decode ...
func Base64Decode(text string) string {
	dec, err := b64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}
	return string(dec)
}
