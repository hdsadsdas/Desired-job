package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {
	//密钥
	key := []byte("00000000")

	//要加密的数据
	arr := "哈哈哈"
	fmt.Println("_____________DES 加密解密字节数组")
	fmt.Println("加密前：",arr)

	//得到加密后的数据
	resultArr, _ := DesEncrypt([]byte(arr), key)
	fmt.Printf("加密后数据:%x\n",resultArr)

	//得到解密后的数据
	resultArr , _ = Desdecrypt(resultArr, key)
 	fmt.Println("解密后数据",string(resultArr))

	//DES加密
	cipherText ,_ := DesEncrypt([]byte(arr),key)
	fmt.Println("DES加密后",cipherText)

	//DES解密
	originalText, _ := Desdecrypt(cipherText, key)
	fmt.Println("DES解密后",string(originalText))

}



//DES 解密字节数组 返回字节数组
func Desdecrypt(cipherBytes , key []byte)([]byte,error){

    //生成数据块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//选择解密模式
	blockMode := cipher.NewCBCDecrypter(block, key)

    //创建一个用来储存解密后数据的 字节切片
	originalText := make([]byte,len(cipherBytes))

	blockMode.CryptBlocks(originalText,cipherBytes)

	originalText =  PKCS5UnPadding(originalText)

	return originalText,nil

}


//DES 加密字节数组，返回加密后的字节数据
func DesEncrypt(originalBytes , key []byte)([]byte , error)  {

	block, err := des.NewCipher(key) //生成数据块
	if err != nil {
		return nil, err
	}

	//得到补码后的数据
	originalBytes  = PKCS5Padding(originalBytes,block.BlockSize())

	//选择加密模式 （CBC加密模式）
	blockMode := cipher.NewCBCEncrypter(block, key)

	//生成存储加密后数据的 字节切片
	cipherArr := make([]byte,len(originalBytes))
	//进行加密
	blockMode.CryptBlocks(cipherArr,originalBytes)
	return cipherArr,nil

}

//尾部填充  补码
func PKCS5Padding(ciphertext []byte,blockSize int) []byte{

	//得到最后差的位数
	padding := blockSize - len(ciphertext)%blockSize

	padtext := bytes.Repeat([]byte{byte(padding)},padding)

	return append(ciphertext,padtext...)

}

//尾部去码 去码
func PKCS5UnPadding(origData []byte)[]byte {

	length := len(origData)

	unpadding := int(origData[length-1])

	return origData[:length-unpadding]

}
