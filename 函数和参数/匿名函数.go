package main

import (
	"fmt"
	"time"
)

/*匿名典型场景一：延时执行一段代码*/
func mainDefer() {
	defer fmt.Println("我是倒数第一")

	defer func() {
		fmt.Println("我是倒数第二")
		fmt.Println("我是倒数第三")
		fmt.Println("我是倒数第四")
	}()
	fmt.Println("我是正数第一")
}

/*典型应用场景二：并发执行一段代码*/
func mainBing() {
	// go 关键字跑在子协程
	go func() {
		fmt.Println("我是倒数第二")
		time.Sleep(1 * time.Second)
		fmt.Println("我比较能睡")
	}()
	// 跑在主协程
	fmt.Println("我是正数第一")
	time.Sleep(1 * time.Second)
	fmt.Println("我是正数第二")
	time.Sleep(1 * time.Second)
	fmt.Println("我是正数第三")
}

/*带有参数和返回值的匿名函数*/
func them(name string, age int) (character int) {
	fmt.Printf("姓名：%s\n", name)
	fmt.Printf("年龄：%d\n", age)
	character = 100
	return
}

func mainThem() {
	//defer them("张三", 24)
	fmt.Println("执行")
	defer func(name string, age int) (character int) {
		fmt.Printf("姓名：%s\n", name)
		fmt.Printf("年龄：%d\n", age)
		character = 100
		return
	}("张三", 24)
}
