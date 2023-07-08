package main

import (
	"fmt"
	"strconv"

	"github.com/NotKatsu/Go-Authentication/base64"
	"github.com/NotKatsu/Go-Authentication/random"
	"github.com/NotKatsu/Go-Authentication/time"
)

func get_authentication_token(unique_id uint64) string {
	return base64.Encode(strconv.Itoa(int(unique_id))) + "." + base64.Encode(strconv.Itoa(int(time.Since_Epoch()))) + "." + base64.Encode(random.String(25))
}

func main() {
	result := get_authentication_token(1110636929688539327)
	fmt.Println(result)
}
