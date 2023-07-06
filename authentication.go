package main

import (
	"fmt"

	"github.com/NotKatsu/Go-Authentication/base64"
)

func main() {
	encode_result := base64.Encode("this is my text")
	decode_result := base64.Decode(encode_result)

	fmt.Println(encode_result, decode_result)
}
