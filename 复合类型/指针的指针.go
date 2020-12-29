package main

import "fmt"

func mainPrt() {
	var x = 123
	var mPrt *int = &x

	mmPrt := &mPrt
	fmt.Println(mmPrt)
	fmt.Printf("mmPrt的数据类型：%T\n", mmPrt)

	fmt.Println(*mPrt)
	fmt.Println(*(*mmPrt))
	fmt.Println(**mmPrt)
}
