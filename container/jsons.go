package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

type dataObj struct {
	Data []person `json:"data"`
	Msg  string   `json:"msg"`
}

func main() {
	data := `{
		"data":[{"name":"陈鑫", "sex":"男", "age":25}],
		"msg":"查询成功！"
	}`
	var newData dataObj
	if err := json.Unmarshal([]byte(data), &newData); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", newData)
	}
}
