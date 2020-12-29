package main

import "fmt"

/*
array[start:end]从数组身上截取下标为[start,end]片段，形成片段
start代表开始下标，不写默认从头开始
end代表结束下标（本身不包含），不写默认截取到末尾
*/
func mainSlice() {
	var array = [5]int{3, 6, 4, 8, 2}

	// 半闭半开，含头不含尾
	slice := array[0:3]

	fmt.Printf("array的类型是：%T，值是：%v\n", array, array)
	fmt.Printf("slice的类型是：%T，值是：%v\n", slice, slice)

	// 如果不写尾，默认直接截到最后
	slice = array[0:]
	fmt.Printf("slice的类型是：%T，值是：%v\n", slice, slice)

	// 头尾都不写，默认从头截到尾
	slice = array[:]
	fmt.Printf("slice的类型是：%T，值是：%v\n", slice, slice)

	// 对切片进行截取，得到切片
	sonSlice := slice[0:3]
	fmt.Printf("sonSlice的类型是：%T，值是：%v\n", sonSlice, sonSlice)
}

func mainSlice1() {
	// 初始化一个没有元素的整型切片
	var slice []int = []int{}
	fmt.Printf("类型是：%T，值是%v\n", slice, slice)
	fmt.Printf("切片的长度是：%d，值是%v\n", len(slice), slice)

	slice = append(slice, 0)
	fmt.Printf("切片的长度是：%d，值是%v\n", len(slice), slice)

	slice = append(slice, 3, 4, 5, 6)
	fmt.Printf("切片的长度是：%d，值是%v\n", len(slice), slice)

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice第%d个元素，值是%d\n", i, slice[i])
	}

	for index, value := range slice {
		fmt.Printf("slice第%d个元素，值是%d\n", index, value)
	}

}

func mainSlice2() {
	var slice1 = []int{1, 2, 3}
	slice2 := []int{5, 6, 7}

	// slice1兼容slice2
	// 将另一个切片整体进行追加时
	// 容量先扩大一倍，再两个两个的递增
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

	fmt.Println(len(slice1), cap(slice1))
}

/*
使用make([]int,len)创建指定初始长度的切片，容量默认为0
使用make([]int, len, cap)创建指定初始长度和容量的切片
*/
func mainSlice3() {
	// 创建长度为3的切片
	slice := make([]int, 3)
	// 创建长度为3，容量为5的切片
	//slice := make([]int, 3, 5)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
}
