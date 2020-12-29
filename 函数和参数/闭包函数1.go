package main

import "fmt"

func main011() {
	toTack("武松", 24)
	fmt.Printf("武松的进度是%d", processWusong)

	getTackFunc()
}

var processWusong int = 0
var processHeShang int = 0

func toTack(name string, hours int) {
	fmt.Printf("%s头领带队行军%d小时\n", name, hours)
	if name == "武松" {
		processWusong += hours
	} else {
		processHeShang += hours
	}
}

/*函数也是一种复合类型*/
func mainFuhe() {
	f := func(name string, hours int) {
		fmt.Printf("%s头领带队行军%d小时\n", name, hours)
		processWusong += hours
	}
	fmt.Printf("数据类型：%T", f, f)
}

func main() {
	var process, processLu int

	// 闭包函数的作用是保存【各自的内层函数状态】
	f1 := getTackFunc()
	f2 := getTackFunc()

	process = f1("武松", 12)
	fmt.Printf("二哥的进度：%d\n", process)

	processLu = f2("鲁智深", 24)
	fmt.Printf("二哥的进度：%d\n", processLu)
}

/*
闭包函数
此处返回二参无返回值的函数：func(name string, hours int)
*/
func getTackFunc() func(name string, hours int) (process int) {
	var process int = 0
	// 定义匿名函数
	toTack := func(name string, hours int) int {
		fmt.Printf("%s头领带队行军%d\n", name, hours)
		process += hours
		return process
	}
	return toTack
}
