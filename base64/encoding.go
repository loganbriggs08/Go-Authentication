package helpers

import (
	"encoding/base64"
)

func Encode(text_to_encode any) {

	base64.RawStdEncoding.EncodeToString([]byte(message_to_encode))
}
