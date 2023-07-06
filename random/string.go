package random

import (
	"math/rand"
	"time"
)

func String(length int) string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	end_string := make([]rune, length)

	for i := range end_string {
		end_string[i] = letters[rand.Intn(len(letters))]
	}

	return string(end_string)
}
