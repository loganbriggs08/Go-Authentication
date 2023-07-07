package main

import (
	"fmt"

	"github.com/NotKatsu/Go-Authentication/base64"
	"github.com/NotKatsu/Go-Authentication/random"
)

func get_authentication_token(unique_id uint64) {
	fmt.Println(random.String(25))
}

func main() {
	encode_result := base64.Encode("this is my text")
	decode_result := base64.Decode(encode_result)

	fmt.Println(encode_result, decode_result)

	get_authentication_token(12345678232323232323)
}
