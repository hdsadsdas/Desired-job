package main

import (
	"fmt"
	"数据结构与算法/数组/练习"
)



func main() {

	list := 练习.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	for i:=0;i<10;i++ {

		list.Insert(1,"x5")
        fmt.Println(list)
	}

	fmt.Println(list.String())

	
}
