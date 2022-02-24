package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
)

func main() {
	
	str := "hello"
	fmt.Println("加密前",str)

	str2 := HASH(str,"sha256",false)
	fmt.Println("sha256加密",str2)

	str3 := HASH(str,"sha512",false)
	fmt.Println("sha512加密",str3)

	str4 := sha512Double(str,false)
	fmt.Println("sha512双重加密",str4)

	str5 := sha512DoubleString(str,false)
	fmt.Println("sha512双重加密 hex字符串",str5)
}

func HASH(text string,hashType string, isHex bool) string{
	
	var hashInstance hash.Hash

	switch hashType {

	case "md5" :
		hashInstance = md5.New()
	case "sha1" :
		hashInstance = sha1.New()
	case "sha256" :
		hashInstance = sha256.New()
	case "sha512" :
		hashInstance = sha512.New()
	}

	if isHex {
		arr,_ := hex.DecodeString(text)
		hashInstance.Write(arr)
	}else {
		hashInstance.Write([]byte(text))
	}

	cipherytes := hashInstance.Sum(nil)

	return fmt.Sprintf("%x",cipherytes)
	
}

func sha256Double(text string,isHex bool)[]byte{

	hashInstance := sha256.New()

	if isHex {
		arr,_ := hex.DecodeString(text)
		hashInstance.Write(arr)
	}else {
		hashInstance.Write([]byte(text))
	}
	cipherytes := hashInstance.Sum(nil)
	hashInstance.Reset()
	hashInstance.Write(cipherytes)
	cipherytes = hashInstance.Sum(nil)

	return cipherytes

}

func sha256DoubleString(text string,isHex bool)string{

	hashInstance := sha256.New()

	if isHex {
		arr,_ := hex.DecodeString(text)
		hashInstance.Write(arr)
	}else {
		hashInstance.Write([]byte(text))
	}
	cipherytes := hashInstance.Sum(nil)
	hashInstance.Reset()
	hashInstance.Write(cipherytes)
	cipherytes = hashInstance.Sum(nil)

	return fmt.Sprintf("%x" , cipherytes)

}


func sha512Double(text string,isHex bool)[]byte{

	hashInstance := sha512.New()

	if isHex {
		arr,_ := hex.DecodeString(text)
		hashInstance.Write(arr)
	}else {
		hashInstance.Write([]byte(text))
	}
	cipherytes := hashInstance.Sum(nil)
	hashInstance.Reset()
	hashInstance.Write(cipherytes)
	cipherytes = hashInstance.Sum(nil)

	return cipherytes

}

func sha512DoubleString(text string,isHex bool)string{

	hashInstance := sha512.New()

	if isHex {
		arr,_ := hex.DecodeString(text)
		hashInstance.Write(arr)
	}else {
		hashInstance.Write([]byte(text))
	}
	cipherytes := hashInstance.Sum(nil)
	hashInstance.Reset()
	hashInstance.Write(cipherytes)
	cipherytes = hashInstance.Sum(nil)

	return fmt.Sprintf("%x" , cipherytes)

}

