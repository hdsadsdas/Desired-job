package 练习

import (
	"errors"
	"fmt"
)

//接口
type List interface {

	Size() int//数组的大小

	Get(index int)(interface{},error)//抓取第几个元素

    Set(index int, newval interface{})error //设置修改数据

    Insert(index int,newval interface{})error//插入数据

    Append(newval interface{}) //追加

    Clear()//清空

    Delete(index int)error //删除

    String()string //返回字符串

}

//数据结构
type ArrayList struct {

    dataStore []interface{} //数组的存储
    theSize int //数组的大小

}


func NewArrayList() *ArrayList {

	list := new(ArrayList) // 初始化结构体

	list.dataStore = make([]interface{},0,10)//开辟空间

	list.theSize = 0

	return list

}

//数组的大小
func (list *ArrayList)Size() int  {

	return list.Size() //返回数组的大小

}

//追加
func (list *ArrayList)Append(newval interface{}) {

	list.dataStore = append(list.dataStore,newval)//追加数据

	list.theSize++
	
}


//抓取数据
func (list *ArrayList)Get(index int)(interface{},error){

	if index < 0 || index >= list.theSize {

		return nil,errors.New("索引越界")

	}

     return list.dataStore[index] , nil

}


//显示数组
func (list *ArrayList)String()string{

return fmt.Sprint(list.dataStore)

}

//设置修改数据
func (list *ArrayList)Set(index int, newval interface{})error{

	if index < 0 || index >= list.theSize {

		return errors.New("索引越界")

	}

	list.dataStore[index] = newval


	return nil

}

//判断数组是否满
func(list *ArrayList) checkisFull()  {

	if list.theSize == cap(list.dataStore){ //判断内存的使用

		newdataStore := make([]interface{},2*list.theSize,2*list.theSize) //开辟双倍内存
        copy(newdataStore,list.dataStore) // 将数据拷贝到新的空间里
        list.dataStore = newdataStore

	}

}

//插入数据
func (list *ArrayList) Insert(index int,newval interface{})error{

	list.checkisFull()//检查内存 如果满了自动追加
	list.dataStore = list.dataStore[:list.theSize+1]//插入数据，内存移动一位

	for i := list.theSize;i > index ; i-- { //从后往前移动

      list.dataStore[i] = list.dataStore[i-1]

	}

	list.dataStore[index] = newval
    list.theSize++

	return nil

}

//删除
func (list *ArrayList)Delete(index int)error{

	if index < 0 || index >= list.theSize {

		return errors.New("索引越界")

	}

	list.dataStore = append(list.dataStore[:index],list.dataStore[index+1:]...)
	list.theSize--

	return nil

}

//清空
func (list *ArrayList) Clear(){

	list.dataStore = make([]interface{},0,10)//开辟空间

	list.theSize = 0

}


















