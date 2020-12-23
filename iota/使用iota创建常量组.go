package main

import "fmt"

/*
定义所有星期日
*/
/*const mondey  = 1
const tuesday = 2
const wednesday = 3
const thursday = 4
const friday = 5
const saturday = 6
const sunday  = 7*/
const (
	Sunday = iota
	/*自动延用排头兵的表达式，但iota逐一递增*/
	Mondey
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Mondey)
}
func main011() {
	fmt.Println(Mondey)
}
