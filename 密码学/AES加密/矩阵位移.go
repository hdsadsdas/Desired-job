package main

import "fmt"

func main() {

	arrs := [4][4]string{

		[4]string{"a","b","c","d"},

		[4]string{"a","b","c","d"},

		[4]string{"a","b","c","d"},

		[4]string{"a","b","c","d"},
	}

	Md(&arrs)

	for i := 0 ; i < len(arrs) ; i++ {

		fmt.Println(arrs[i])

	}
}


func Md(arrs *[4][4]string){

	for i := 0 ; i < len(arrs) ; i ++{

		arr := arrs[i]

		for j := 0 ; j < len(arrs[i]) ; j++ {



			if j+i > len(arrs[i])-1 {

				outindex := j+i - len(arrs[i])
				arrs[i][outindex] = arr[j]

			}else {


				arrs[i][j+i] = arr[j]

			}



		}


	}

}
