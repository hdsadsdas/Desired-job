package main

import (
	"bytes"
	"fmt"
)

func main() {

	str := "小薇小薇a"

	strBytes := []byte(str)
	fmt.Println("补码前数据",strBytes)
	paddingBtes := Group(strBytes)

	fmt.Println("补码后数据",paddingBtes)

}

//将进行补码
func Group(strBytes []byte) []byte  {

	if len(strBytes) % 8 == 0 {

		return strBytes

	}

	padding := PKCS5Padding(strBytes, 8)

	return padding

}

//补码
func PKCS5Padding(listGroup []byte,blockSize int) []byte {

	padding := blockSize-len(listGroup)%blockSize

	padtext := bytes.Repeat([]byte{byte(padding)},padding)

	return append(listGroup,padtext...)

}