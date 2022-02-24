package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	fmt.Println("请输入要加密的数据")
	var data string
	fmt.Scanln(&data)


	key := NewKey(len(data))
	encDate := encryption(key, []byte(data))

	Data := decrypt(key, encDate)
	fmt.Println("解密后",Data)

}

// 生成密钥
func NewKey(lens int)[]byte{

	enc := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	rand.Seed(time.Now().UnixNano())

	var keys []byte

	for i := 0 ; i < lens ; i++ {
		key := rand.Intn(len(enc))
		 r := enc[key]
		 keys = append(keys,r)
	}

	fmt.Println("生成的密钥是",string(keys))

	return keys

}

//加密
func encryption(key []byte,data []byte)int{

	keyHex := BytesToHexString([]byte(key))
	//fmt.Println("keyHex",keyHex)
	s1 := Hex2Dec(keyHex)

	dataHex := BytesToHexString([]byte(data))
	//fmt.Println("dataHex",dataHex)
	s2 := Hex2Dec(dataHex)

	encData := s1 ^ s2
	fmt.Println("加密后数据",encData)

	return encData
}

//解密

func decrypt(key []byte,encDate int)string{

	keyHex := BytesToHexString([]byte(key))
	//fmt.Println("keyHex",keyHex)
	s1 := Hex2Dec(keyHex)

	data := s1 ^ encDate
	dataHex := DecimalToHex(data)


	bytesData, err := HexStringToBytes(dataHex)
	if err != nil {
		fmt.Println("16进制转字符串失败",err.Error())
	}

	return string(bytesData)

}


//将字符串转为16进制字符串
func BytesToHexString(arr []byte) string{

	return hex.EncodeToString(arr)

}

//将16进制转为10进制

func Hex2Dec(x string) int {

	base, _ := strconv.ParseUint(x, 16, 64)
	return int(base)

}

//将10进制转为16进制
func DecimalToHex(Decimal int)(Hexstring string){

	Hexstring = fmt.Sprintf("%x",Decimal)
	return

}

//将十六进制转为字符串
func HexStringToBytes(s string)(arr []byte,err error){

	arr,err = hex.DecodeString(s)
	return

}