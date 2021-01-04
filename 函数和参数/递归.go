package main

import (
	"fmt"
	"time"
)

/*求1+2+...+n的和----循环实现*/
func GetSum(n int) (sum int) {
	for i := 1; i < n+1; i++ {
		sum += i
	}
	return sum
}

/*求1+2+...+n的和---递归实现*/
func GetDi(n int) (sum int) {
	if n == 1 {
		return 1
	}
	sum = n + GetDi(n-1)
	return sum
}

/*递归与循环效率比较*/
func timeIt(GetDi func(int) int, arg int) (ret int) {
	startTime := time.Now()
	ret = GetDi(arg)
	endTime := time.Now()
	fmt.Println("执行耗时：", endTime.Sub(startTime))
	return
}
