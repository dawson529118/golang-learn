package main

import (
	"fmt"
	"time"
)

func main() {
	liuNeihe := ke("刘备", 0)
	fmt.Println(liuNeihe())
	go func() {
		for i := 0; i < 7; i++ {
			wujiang := liuNeihe()
			fmt.Println(wujiang)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(8 * time.Second)
}

var ones = []string{"关羽", "张飞", "赵云", "马超", "黄忠"}

func ke(zhugong string, start int) func() string {
	// 内核函数状态
	var suoyin int = start
	// 内核函数：得到主公的值班武将
	neiHe := func() string {
		one := ones[suoyin]
		suoyin++
		if suoyin > 4 {
			suoyin = 0
		}
		return zhugong + "麾下" + one
	}
	return neiHe
}
