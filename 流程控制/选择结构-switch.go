package main

import (
	"fmt"
	"strconv"
)

func mainSwitch() {
	fmt.Println("请输入今天星期几：")
	var userInput string
	fmt.Scan(&userInput)

	var weekday int
	weekday, _ = strconv.Atoi(userInput)
	switch weekday {
	case 1:
		fmt.Println("星期一比较困！！！")
	case 2:
		fmt.Println("星期二好了一点！！！")
	case 3:
		fmt.Println("星期三感觉过的好快！！！")
	case 4:
		fmt.Println("明天就是星期五了！！！")
	case 5:
		fmt.Println("哇！今天就星期五了！！！")
	case 6:
		fmt.Println("周六又是晚起的一天！！！")
	case 7:
		fmt.Println("明天又要周一了！！！")
	default:
		fmt.Println("后续不足！！！")
	}
}
func mainTwo() {
	var age = 20
	switch {
	case age < 18:
		fmt.Println("还未成年！！！")
	case age >= 18 && age < 36:
		fmt.Println("青年人阶段！！！")
	case age >= 36 && age < 60:
		fmt.Println("中年人阶段！！！")
	case age >= 60 && age < 100:
		fmt.Println("老年人阶段！！！")
	default:
		fmt.Println("仙人阶段！！！")
	}
}
