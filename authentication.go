package main

import (
	"fmt"

	"github.com/NotKatsu/Go-Authentication/base64"
)

func get_authentication_token(unique_id int) string {

}

func main() {
	encode_result := base64.Encode("this is my text")
	decode_result := base64.Decode(encode_result)

	fmt.Println(encode_result, decode_result)
}
