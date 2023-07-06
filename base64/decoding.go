package base64

import (
	"encoding/base64"
)

func Decode(text_to_decode string) string {
	base64_decode_result, err := base64.RawStdEncoding.DecodeString(text_to_decode)

	if err == nil {
		return string(base64_decode_result)
	} else {
		return ""
	}
}
