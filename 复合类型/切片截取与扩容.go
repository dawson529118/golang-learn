package main

import "fmt"

/*
扩容策略：容量不够时，容量*2（有时+2），找到一片新的连续内存，将原有元素拷贝进去
*/
func mainStr() {

	// 创建一个0长度的整型切片
	slice1 := make([]int, 0)
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 1)
	fmt.Printf("长度是：%d，容量是：%d，值是：%v\n", len(slice1), cap(slice1), slice1)
	fmt.Println("首元素地址值：", &slice1[0])
	slice1 = append(slice1, 2)
	fmt.Printf("长度是：%d，容量是：%d，值是：%v\n", len(slice1), cap(slice1), slice1)
	fmt.Println("首元素地址值：", &slice1[0])
	slice1 = append(slice1, 3)
	fmt.Printf("长度是：%d，容量是：%d，值是：%v\n", len(slice1), cap(slice1), slice1)
	fmt.Println("首元素地址值：", &slice1[0])
	slice1 = append(slice1, 4)
	fmt.Printf("长度是：%d，容量是：%d，值是：%v\n", len(slice1), cap(slice1), slice1)
	fmt.Println("首元素地址值：", &slice1[0])
}

/*
切片扩容的一刹那，所有元素的地址都重新拷贝到一片新的连续内存中
*/
func mainStr1() {
	var array = [5]int{1, 2, 3, 4, 5}
	slice1 := array[:]
	slice2 := array[:]

	// array,slice1,slice2的具有完全相同的元素初始地址
	fmt.Printf("%p, %p, %p\n", &array, &slice1, &slice2)          // 各不相同
	fmt.Printf("%p, %p, %p\n", &array[0], &slice1[0], &slice2[0]) // 相同
	fmt.Printf("%p, %p, %p\n", &array[1], &slice1[1], &slice2[1]) // 相同

	// 牵一发而动全身
	array[1] = 111
	slice1[2] = 222
	slice2[3] = 333
	fmt.Println(array)
	fmt.Println(slice1)
	fmt.Println(slice2)

	// slice2扩容之后，所有元素的地址都会发生变化
	slice2 = append(slice2, 5, 6, 7)
	array[0] = 10
	slice1[1] = 111
	slice2[0] = 1000
	fmt.Println(array)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Printf("addr2=%p, cap2=%d", &slice2, cap(slice2))
}
