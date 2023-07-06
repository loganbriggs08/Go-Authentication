package base64

import (
	"encoding/base64"
)

func Encode(text_to_encode string) string {
	base_64_encode_result := base64.RawStdEncoding.EncodeToString([]byte(text_to_encode))

	return base_64_encode_result
}
