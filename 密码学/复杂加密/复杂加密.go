package main

import (
	"fmt"
	"math/rand"
	"time"
)

func jia(miwen []int32)[]int32{

	rand.Seed(time.Now().UnixNano())

     var my []int32

	for k,v := range miwen {

		key := rand.Intn(11)
		my = append(my,int32(key))
		miwen[k] = v+int32(key)
        fmt.Println(v+int32(key))
	}

	return my

}

func jie(miwen []int32,my []int32){

	for k,v := range miwen {

		miwen[k] = v - my[k]

	}

}


func main() {

	str := "hello"
	miwen := []rune(str)



	my := jia(miwen)
	fmt.Println("密钥是",my)
	fmt.Println("加密",string(miwen))

	jie(miwen,my)

	fmt.Println("解密",string(miwen))


}