package main

import "fmt"

func main() {

	//先创建一个二维数组，来模拟迷宫

	//规则
	//如果元素的值为一就表示墙
	//如果元素的值为0 表示没有走过的点
	//如果元素的值为2，是一个通路
	//如果元素值为3，表示为走过的点但走不通


	var  mymap [8][7]int

	for i := 0;i < 7;i++ {

		mymap[0][i] = 1

		mymap[7][i] = 1

	}

	for i := 0;i < 8;i++ {

		mymap[i][0] = 1

		mymap[i][6] = 1

	}

	mymap[3][1] = 1
	mymap[3][2] = 1

	for i := 0 ; i < 8 ;i++  {

		for j := 0 ; j < 7 ; j++ {

			fmt.Print(mymap[i][j]," ")

		}

       fmt.Println()

	}

	//测试

	SetWay(&mymap,1,1)


	fmt.Println("================探测完毕================")
	//探测完毕的地图
	for i := 0 ; i < 8 ;i++  {

		for j := 0 ; j < 7 ; j++ {

			fmt.Print(mymap[i][j]," ")

		}

		fmt.Println()

	}

}


//编写一个函数 完成 找路
//i , j 表示对地图的哪个点进行测试
func SetWay(mymap *[8][7]int,i int,j int) bool{

	//当mymap[6][5] = 2 表示找到出口

	if mymap[6][5] == 2 {

		return true

	}else {

		//继续

		if mymap[i][j] == 0 {//如果这个点没有探测的

			//假设这个点是可以通的,但是需要探测 下右上左
			mymap[i][j] = 2
			if SetWay(mymap,i+1,j) {//向下走

				return true

			}else if SetWay(mymap,i,j+1){//向右走

				return true


			}else if SetWay(mymap,i-1,j) {  //向上走

				return true

			}else if SetWay(mymap,i,j-1){//向左走

				return true

			}else {  //说明是死路

				mymap[i][j] = 3
				return false

			}

		}else { //说明这个点不能探测

			return false

		}


	}

}