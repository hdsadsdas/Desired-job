package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

const (
	URL = "http://web.juhe.cn/constellation/getAll"
	KEY = "7c8ea7514ddbb719778b58b2d43e25eb"
)

func main() {

	
	fmt.Println("请输入你要查询的星座")
	var consName string
	fmt.Scanln(&consName)
	consName = url2.QueryEscape(consName)
	
	fmt.Println("请输入查询的类型如today")
	var Type string
	fmt.Scanln(&Type)

	url := URL + "?consName=" + consName + "&type=" + Type + "&key=" + KEY

	resp, err := http.Get(url)

	if err != nil {

		fmt.Println("网络请求错误")
		return

	}

	if resp.StatusCode != 200 {

		fmt.Println("网络错误")
		return
	}

	all, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("读取数据错误",err.Error())
		return
	}


	var data map[string]interface{}

	err = json.Unmarshal(all,&data)

	if err != nil {

		fmt.Println("反序列化失败")
		return

	}

	fmt.Println(data["name"])
	fmt.Println(data["summary"])

}
