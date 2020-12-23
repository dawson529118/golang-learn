package main

import "fmt"

func mainWei() {

	fmt.Printf("12 & 10的结果：十进制%d，二进制%b\n", 12&10, 12&10)
	fmt.Printf("12 | 10的结果：十进制%d，二进制%b\n", 12|10, 12|10)
	fmt.Printf("12 ^ 10的结果：十进制%d，二进制%b\n", 12^10, 12^10)

	var a int8 = 12
	fmt.Println(a << 2)
	fmt.Println(a << 3)
	fmt.Println(a << 4)
}
