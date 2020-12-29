package main

import "fmt"

func main01() {

	//  声明变量a时，系统开辟一块内存【地址】，里面存的值是【123】
	var a int = 123
	fmt.Printf("a的类型是：%T\n", a)
	fmt.Printf("a的值是：%v\n", a)
	fmt.Printf("a的地址是：%p\n", &a)

	// &a取值变量a的地址
	aPointer := &a
	fmt.Printf("aPointer的类型是：%T\n", aPointer)

	// 将aPointer指向的地址的值修改为456
	*aPointer = 456
	fmt.Println("*aPointer=", aPointer)
	fmt.Println("a=", a)
}

func main02() {
	var x = 456
	fmt.Println(x)

	var xPtr *int
	fmt.Println("xPtr=", xPtr)

	xPtr = &x

	// 将xPtr指向的地址的值修改为【789】
	*xPtr = 789

	fmt.Println(x)
	fmt.Println(xPtr)
	fmt.Println(&x)
	fmt.Println(&xPtr)
	fmt.Println(*xPtr == x)
	fmt.Println(xPtr == &x)

	var y = 456
	*xPtr = y
	fmt.Println(x)
	fmt.Println(xPtr)
	fmt.Println(&x)
	fmt.Println(*xPtr)
	fmt.Println(*xPtr == y)
	fmt.Println(x == y)
	fmt.Println(&x == &y)
	fmt.Println(xPtr == &y)
}
