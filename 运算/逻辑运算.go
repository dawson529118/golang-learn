package main

import "fmt"

func mainLuo() {
	a1 := (1 + 1) == 2
	a0 := (1 + 1) != 2

	b1 := (2 + 3) > 4
	b0 := (2 + 3) <= 4

	// 逻辑与：双方都成立，结果才成立
	fmt.Println(a1 && b1)
	fmt.Println(a1 && b0)
	fmt.Println(a0 && b1)
	fmt.Println(a0 && b0)

	// 逻辑或：有一个成立，就成立
	fmt.Println(a1 || b1)
	fmt.Println(a1 || b0)
	fmt.Println(a0 || b1)
	fmt.Println(a0 || b0)

	// 逻辑非：运算结果与原条件相反
	fmt.Println(!a1)
	fmt.Println(!b1)
	fmt.Println(!a0)
	fmt.Println(!b0)
}
