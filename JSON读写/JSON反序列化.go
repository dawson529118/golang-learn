package main

import (
	"encoding/json"
	"fmt"
)

func mainAllJosn() {
	// 将json转换为map
	jsonStr := `{"Name":"小明","Age":24,"Hobby":["篮球","网球"]}`
	jsonBytes := []byte(jsonStr)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		fmt.Println("json转map错误，err=", err)
		return
	}
	fmt.Println(jsonMap)

	// 将json转换为结构体
	type Person struct {
		Name  string
		Age   int
		Hobby []string
	}
	person := new(Person)
	err = json.Unmarshal(jsonBytes, &person)
	if err != nil {
		fmt.Println("json转结构体错误，err=", err)
		return
	}
	fmt.Println(*person)

	// json转map切片
	jsonStr = `[{"hobby":["抽中华","喝牛栏山"], "name":"王钢蛋"}]`
	dataSlice := make([]map[string]interface{}, 0)
	jsonBytes = []byte(jsonStr)
	err = json.Unmarshal(jsonBytes, &dataSlice)
	if err != nil {
		fmt.Println("json转map切片错误，err=", err)
		return
	}
	fmt.Println(dataSlice)

	// json转结构体切片
	jsonStr = `[{"hobby":["抽中华","喝牛栏山"], "name":"王钢蛋", "age":24}]`
	persons := make([]Person, 0)
	jsonBytes = []byte(jsonStr)
	err = json.Unmarshal(jsonBytes, &persons)
	if err != nil {
		fmt.Println("json转结构体切片错误，err=", err)
		return
	}
	fmt.Println(persons)
}
