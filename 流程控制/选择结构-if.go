package main

import "fmt"

/*单分支*/
func mainAa() {
	fmt.Println("请输入相亲者姓名：")

	// 声明一个字符串变量name
	var name string

	// 接收用户输入，存入name变量所在地址
	fmt.Scan(&name)

	if name == "张三" {
		fmt.Println("我是张三！！！")
		return
	}

}

/*多分支*/
func mainBb() {

	fmt.Println("请输入相亲者姓名：")

	var name string
	fmt.Scan(&name)
	if name == "张三" {
		fmt.Println("我是张三！！！")
	} else if name == "李四" {
		fmt.Println("我是李四！！！")
	} else {
		fmt.Println("我不是张三，也不是李四！！！")
	}
	fmt.Println("Game Over!!!")
}
