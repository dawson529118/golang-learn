/*
%T  类型占位符
%v  值占位符
%d  整型占位符
%f  浮点型占位符
%c  字符占位符
%s  字符串占位符
*/
package main

import "fmt"

func main() {
	var v1 = 123
	fmt.Printf("v1的类型是%T\n", v1)

	var v2 int = 123
	fmt.Printf("v2的类型是%T\n", v2)

	var v3 = 123.0
	fmt.Printf("v3的类型是%T\n", v3)

	var v4 float32 = 123.0
	fmt.Printf("v4的类型是%T\n", v4)

	var v5 = "hello"
	fmt.Printf("v5的类型是%T\n", v5)

	var v6 = '俊'
	fmt.Printf("v6的类型是%T,v6的值是%v\n", v6, v6)
	fmt.Printf("v6的字符是%c\n", v6)

	var v7 = (100 == (40 + 60))
	fmt.Printf("v7的类型是%T, 值是%v\n", v7, v7)
}
