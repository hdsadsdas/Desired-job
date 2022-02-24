package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	 BASE_URL = "http://apis.juhe.cn/simpleWeather/query"
	 KEY = "4f40b620803ec85ef0a196e7aeae4e51"
)

//go语言字段或者函数/方法等的命名首字母的大小写规范：
//通过大小写来规范代码的可见性
//首字母大写表示外部可以访问
//首字母小写表示外部不可以访问
//外部的内部的 用包来进行区分
type Weather struct {
	Reason     string `json:"reason"`
	Result     Result `json:"result"`
	Error_code int    `json:"error_code"`
}

type Result struct {
	City     string   `json:"city"`
	Realtime RealTime `json:"realtime"`
	Future   []Future `json:"future"`
}

type RealTime struct {
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
	Info        string `json:"info"`
	Wid         string `json:"wid"`
	Driect      string `json:"driect"`
	Power       string `json:"power"`
	Aqi         string `json:"aqi"`
}

type Future struct {
	Date        string `json:"date"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
	Wid         Wid    `json:"wid"`
	Direct      string `json:"direct"`
}

type Wid struct {
	Day   string `json:"day"`
	Night string `json:"night"`
}


func main() {

	//1 . 读取用户的键盘输入

	var city string
    fmt.Println("请输入城市")
	_, err := fmt.Scan(&city)

	if err != nil {

		panic(err)
		return

	}


	//2 . 将用户的输入内容 和 第三方api 进行拼接

	url := BASE_URL + "?city=" + city + "&key="+ KEY

	//发起网络请求
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("网络错误",resp.StatusCode)
	}

	//3 . 接受网络请求的响应结果

	data , err := ioutil.ReadAll(resp.Body)

	if err != nil{
		fmt.Println("读取数据错误",err.Error())
		return
	}

	fmt.Println("获取成功"+string(data))

	//将得到的结果反序列化

	var netReturn map[string]interface{}

	err = json.Unmarshal(data, &netReturn)
	if err != nil {
		fmt.Println("反序列化错误",err.Error())
		return
	}

	/* {
	    "reason": ,
	    "result": {},
	    "future": [],
	    "error_code": 0
	   }
	 */
	data2 := netReturn["result"]
		realtime := data2.(map[string]interface{})["realtime"]

		fmt.Printf("温度：%v\n湿度：%v\n天气：%v\n风向：%v\n风力：%v\n空气质量：%v",
			realtime.(map[string]interface{})["temperature"],
			realtime.(map[string]interface{})["humidity"],
			realtime.(map[string]interface{})["info"],
			realtime.(map[string]interface{})["direct"],
			realtime.(map[string]interface{})["power"],
			realtime.(map[string]interface{})["aqi"],
		)


	 future	:= realtime.(map[string]interface{})["future"]
     fmt.Printf("%T\n",future)

    var weater Weather

	err = json.Unmarshal(data, &weater)
	if err != nil {

		fmt.Println("反序列化失败",err)

	}

	for _,future := range weater.Result.Future{

		fmt.Println("日期:",future.Date,"温度",future.Temperature,"天气",future.Weather,"风向",future.Direct)

	}

}
