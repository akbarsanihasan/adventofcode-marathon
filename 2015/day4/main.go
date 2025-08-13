package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func CreateMd5Hash(text string) string {
	hasher := md5.New()
	_, err := io.WriteString(hasher, text)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func answer(length int) int {
	PUZZLE := []byte("ckczppom")
	input := 1
	whole_input := string(PUZZLE) + fmt.Sprint(input)
	hash := CreateMd5Hash(whole_input)

	var condition string
	for range length {
		condition = condition + "0"
	}

	for hash[0:length] != condition {
		input++
		whole_input = string(PUZZLE) + fmt.Sprint(input)
		hash = CreateMd5Hash(whole_input)
	}

	return input
}

func main() {
	fmt.Println(answer(5))
	fmt.Println(answer(6))
}
