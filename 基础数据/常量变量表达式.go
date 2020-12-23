package main

import "fmt"

func main021() {
	// 声明常量（const关键字）
	const pi = 3.14

	// 声明变量（variable关键字）
	var r = 1.0

	// 声明变量area
	area := pi * r * r
	fmt.Println("圆的面积是：", area)

	// pi = 3.1415926
	r = 2.0
	area = pi * r * r
	fmt.Println("圆的面积是：", area)
}

/*
一次性声明多个常量和变量
*/
func main() {
	fmt.Println("我又活过来了")

	//  一次性声明多个常量
	const (
		lightSpeed  = 30 * 10000
		earthCircle = 40000
	)

	fmt.Println("光速是：", lightSpeed, "地球周长是：", earthCircle)

	// 一次性声明多个变量
	var (
		age    = 20
		height = 175
	)
	fmt.Println("年龄是：", age, "身高是：", height)
}
