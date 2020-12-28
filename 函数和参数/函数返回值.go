package main

import "fmt"

func mainReturn() {
	result := sum(3, 4)
	fmt.Println(result)
	resultSum := getSum(4, 5)
	fmt.Println(resultSum)
	sum, cha := calc(6, 3)
	fmt.Printf("和：%d, 差：%d\n", sum, cha)
	sum1, cha1 := calc1(8, 2)
	fmt.Printf("和：%d, 差：%d", sum1, cha1)
}

/*一个返回值*/
func sum(a, b int) int {
	result := a + b
	return result
}

/*预定义返回值的名字*/
func getSum(a, b int) (sum int) {
	sum = a + b
	return
}

/*多个返回值*/
func calc(a, b int) (int, int) {
	sum := a + b
	cha := a - b
	return sum, cha
}

/*多个返回值+预定义返回值名称*/
func calc1(a, b int) (sum int, cha int) {
	sum = a + b
	cha = a - b
	return
}
