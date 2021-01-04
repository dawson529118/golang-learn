package main

import "fmt"

/*
1 1 2 3 5 8 13 21...
*/
func mainFibonacci() {
	/*sum := GetFibonacci(4)
	fmt.Printf("获取到的数值：%d", sum)*/
	for i := 0; i < 10; i++ {
		fmt.Println(i, "\t", GetFibonacci(i))
	}
}

/*递归求斐波那契数列第N项的值*/
func GetFibonacci(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return GetFibonacci(n-1) + GetFibonacci(n-2)
}
