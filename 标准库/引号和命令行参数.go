package main

import (
	"flag"
	"fmt"
	"os"
)

/*引号差异*/
func mainYin() {
	// 字符
	fmt.Println('我')
	// 字符串
	fmt.Println("字符串")
	// 原样输出
	fmt.Println(`
闭包就像一层壳
	内层函数是内核
	`)
}

/*获取命令行参数：简易版*/
func mainSimple() {
	fmt.Println(os.Args)
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}

/*获得命令行参数*/
func mainFlag() {

	// 定义参数
	namePtr := flag.String("name", "无名氏", "姓甚名谁")
	agePtr := flag.Int("age", -1, "阁下的年龄")
	rmbPtr := flag.Float64("rmb", 1.0, "你的资产")
	alivePtr := flag.Bool("alive", true, "是否健在")

	// 解析获取参数，丢入参数的指针中
	flag.Parse()
	fmt.Println(*namePtr, *agePtr, *rmbPtr, *alivePtr)
}

/*引用地址方法实现*/
func mainAddre() {

	var name string
	var age int
	var rmb float64
	var alive bool

	// 定义参数
	flag.StringVar(&name, "name", "无名氏", "姓甚名谁")
	flag.IntVar(&age, "age", -1, "阁下的年龄")
	flag.Float64Var(&rmb, "rmb", 1.0, "你的资产")
	flag.BoolVar(&alive, "alive", true, "是否健在")

	// 解析获取参数，丢入参数的指针中
	flag.Parse()
	fmt.Println(name, age, rmb, alive)
}
