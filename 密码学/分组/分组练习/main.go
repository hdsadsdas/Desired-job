package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	str := "xiaoweixiaoweixiaoweixi小薇aoweixiaoweixiaoweixiaoweixiaoweixiaowei"


	hexString := BytesToHexString([]byte(str))
	fmt.Println(hexString)

	Group := make(map[int]string)
	GroupCount := len(hexString)/8

	for i := 0 ;i<GroupCount;i++ {

		Group[i+1] = hexString[i*8:i*8+8]

	}

	//最后一组
	Group[GroupCount+1] = hexString[GroupCount*8:]
	fmt.Println(Group)

	for i := 1 ; i<= len(Group) ; i++  {

		Group[i] = HexToBin(Group[i])

	}

	fmt.Println(Group)


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

