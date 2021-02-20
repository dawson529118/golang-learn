package main

import (
	"encoding/json"
	"fmt"
)

type PersonInfo struct {
	Name  string
	Age   int
	Hobby []string
}

func mainSear() {

	// 使用结构体转json
	persion := PersonInfo{"小明", 24, []string{"篮球", "网球"}}
	data, err := json.Marshal(&persion)
	if err != nil {
		fmt.Println("序列化失败, err=", err)
		return
	}
	fmt.Println(string(data))

	// map转json
	dataMap := make(map[string]interface{})
	dataMap["name"] = "小明"
	dataMap["age"] = 24
	dataMap["hobby"] = []string{"篮球", "网球"}

	data, err = json.Marshal(dataMap)
	if err != nil {
		fmt.Println("map转换json错误，err=", err)
		return
	}
	fmt.Println(string(data))

	// 使用map切片转json
	dataMap1 := make(map[string]interface{})
	dataMap1["name"] = "小明"
	dataMap1["age"] = 24

	dataMap2 := make(map[string]interface{})
	dataMap2["name"] = "小蓝"
	dataMap2["age"] = 24

	splice := make([]map[string]interface{}, 0)
	dataSplice := append(splice, dataMap1, dataMap2)
	data, err = json.Marshal(dataSplice)
	if err != nil {
		fmt.Println("切片转换json错误，err=", err)
		return
	}
	fmt.Println(string(data))
}
