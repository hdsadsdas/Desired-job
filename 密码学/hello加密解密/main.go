package main

import (
	"fmt"
	"math/rand"
	"time"
)

func jia(str string,str2 *string)(int){

	for _,v := range str{
		
		*str2 += string(v+3)

	}

	fmt.Println("加密后",*str2)
   return 3
}

func jie(str2 string,my int)  {

	var str string

	for _,v := range str2{

		str += string(v-3)

	}

	fmt.Println("解密后",str)
	
}

func jia2(miwen []int32)int{

	rand.Seed(time.Now().UnixNano())
	my := rand.Intn(15)


	for k,v := range miwen {

			miwen[k] = v+int32(my)

	}

	return my

}

func jie2(miwen []int32,my2 int32){

	for k,v := range miwen {

		miwen[k] = v - my2

	}

}


func main() {

	str := "hello"
	//var str2 string
	miwen := []rune(str)


	//my := jia(str,&str2)
	//
	//jie(str2,my)

	 my2 := jia2(miwen)
    fmt.Println("密钥是",my2)
	fmt.Println("加密",string(miwen))

	jie2(miwen,int32(my2))

	 fmt.Println("解密",string(miwen))



}
