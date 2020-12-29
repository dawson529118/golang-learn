package main

import "fmt"

func mainArray() {
	// 创建数组
	//array := [5]int{1,2,3,4,5}
	array := [...]int{3, 4, 5, 6}
	fmt.Println(array)

	// 获得数组的长度
	// 通过内建函数len(x)得到数组长度
	fmt.Printf("数组的长度是：%d\n", len(array))

	// 根据下标对元素进行访问
	fmt.Printf("数组中第一个元素：%d\n", array[0])
	array[0] = 333
	fmt.Printf("数组中第一个元素：%d\n", array[0])

	// 数组越界异常 index out of range
	//myLen := 20 + 30
	//fmt.Println("第50个元素：", array[myLen])

	// 遍历数组方式一
	/*for i:=0;i<len(array);i++ {
		fmt.Printf("数组第%d个元素是%d\n", i, array[i])
	}*/

	// 遍历数组方式二
	for index, value := range array {
		fmt.Printf("数组第%d个元素是%d\n", index, value)
	}
}
