package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	URL = "http://v.juhe.cn/toutiao/index"
	KEY = "d8c6db2c5ac0000b78f0f2ca007bb888"
)

func main() {

	params := make(map[string][]string)
	params["key"] = []string{KEY}
	params["type"] = []string{"top"}

	resp, err := http.PostForm(URL,params)
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

		fmt.Println("读取数据失败",err.Error())
		return

	}

	var data map[string]interface{}

	err = json.Unmarshal(all, &data)
	if err != nil {

		fmt.Println("反序列化失败",err.Error())
		return

	}

	fmt.Println(data["reason"])

	result := data["result"]

	//fmt.Println(result.(map[string]interface{})["data"])

	datt := result.(map[string]interface{})["data"]

	fmt.Printf("%T\n",datt)

	fmt.Println(datt.([]interface{})[1])

	aaa := datt.([]interface{})[1]

	fmt.Printf("%T\n",aaa)
	fmt.Println(aaa.(map[string]interface{})["date"])

}
