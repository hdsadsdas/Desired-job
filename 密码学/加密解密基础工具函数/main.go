package main

import (
	"encoding/hex"
	"fmt"
)

func main() {

	arr := []byte{'h','e','l','l','o'}
	fmt.Println("加密前",string(arr))
	str := BytesToString(arr)
	fmt.Println("加密后",str)
	str = ReverseHexString(str)
	arr, _ = HexStringToBytes(str)
	//fmt.Printf("%x\n",arr)
	ReverseBytes(arr)
	fmt.Println(string(arr))
	
}

//将byte[]转成string

func BytesToString(arr []byte)string{

	return hex.EncodeToString(arr)

}

//将String转成byte[]

func HexStringToBytes(s string)(arr []byte,err error){

	arr,err = hex.DecodeString(s)
	return

}

//将字符串颠倒
//Reverse  相反
//Hex 十六进制

func ReverseHexString(hexStr string)string{

	arr, _ := hex.DecodeString(hexStr)
    ReverseBytes(arr)
	return hex.EncodeToString(arr)

}

//将byte[]颠倒

func ReverseBytes(data []byte){

	for i,j := 0 ,len(data)-1;i < j ; i,j = i + 1,j-1 {

		data[i],data[j] = data[j],data[i]

	}

}


