package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	text := "Zewqeqweqvs"

	fmt.Println("原文",text)

	cipherText := base64EncodeString(text)

	fmt.Println("base64编码后",cipherText)

	fmt.Println("解码后",base64DecodeString(cipherText))
	
}

func base64EncodeString(text string)string{

	return base64.StdEncoding.EncodeToString([]byte(text))

}

func base64DecodeString(text string)string{

	decodeString, _:= base64.StdEncoding.DecodeString(text)

	return string(decodeString)
}

