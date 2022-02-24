package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {

	str := "小薇xiaoweixiaoweixiaoweixiaoweixiaoweixiaoweixiaoweixiaoweixiaowei"

	hexString := BytesToHexString([]byte("zzzzzzzz"))
	fmt.Println(hexString)

	//转二进制
	bin := HexToBin(hexString)
	fmt.Println(bin)
	fmt.Println(len(bin))

	//转10进制
	Dec := HexToDec(hexString)
	fmt.Println(Dec)


	//返回的2进制字符串
	binary := StringToBinary(str)


	//分为几组
	GroupCount := len(binary)/64

	//用来存储分割后的数据
	Group := make(map[int]string)

	for i := 0 ; i < GroupCount ; i++ {

		Group[i+1] = binary[i*64:i*64+64]

	}
	//最后一组
    if len(bin) % 64 != 0{
		Group[GroupCount+1] = binary[GroupCount*64:]
		fmt.Println("分为",GroupCount+1,"组")
		fmt.Println("数据是",Group)
	}

}

//将字符串转为2进制
func StringToBinary(s string) (Bin string) {

	for _,c := range s {
		Bin = fmt.Sprintf("%s%b",Bin,c)
	}

	return

}

//将字符串转为16进制字符串
func BytesToHexString(arr []byte) string{

	return hex.EncodeToString(arr)

}

//将十六进制转为二进制
func HexToBin(str string) string {
	base, _ := strconv.ParseInt(str, 16, 64)
	return strconv.FormatInt(base, 2)
}

//将hex转为10进制
func HexToDec(str string) int64  {

	parseInt, _ := strconv.ParseInt(str, 16, 64)
    return parseInt

}
