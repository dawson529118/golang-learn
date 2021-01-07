package main

import "fmt"

/*
判断一个接口实例的具体类型

类型断言的两种方式
方式一，判断接口实例是不是程序员
switch xxx.(type) {
case: *Coder
...
case: *Boos
}

方式二，判断接口实例是不是程序员
if coder, ok := xxx.(*Coder);ok {
	// 确实是程序员
	// 此时的Coker是程序员指针类型
}else {
	// xxx压根不是程序员
}
*/

type WorkerFS interface {
	// 每天工作多少小时，产出何种产品
	Work(hour int) (product string)
	// 休息
	Rest()
}

/*定义码农结构体*/
type CoderFS struct {
	skill string
}

/*码农指针实现worker接口*/
func (c *CoderFS) Work(hour int) (product string) {
	fmt.Printf("码农一天工作%d小时\n", hour)
	fmt.Printf("码农正在%s\n", c.skill)
	return "BUG"
}

func (c *CoderFS) Rest() {
	fmt.Println("休息是什么？？")
}

/*码农特有的方法*/
func (c *CoderFS) WorkHome() {
	fmt.Println("程序员在家工作")
}

/*定义产品经理结构体*/
type ProductManagerFS struct {
	skill string
}

/*PriductManger指针实现worker接口*/
func (pm *ProductManagerFS) Work(hour int) (product string) {
	fmt.Printf("产品经理一天工作%d小时\n", hour)
	fmt.Printf("产品正在%s\n", pm.skill)
	return "无逻辑需求"
}

func (pm *ProductManagerFS) Rest() {
	fmt.Println("产品经理休息")
}

/*定义老板结构体*/
type BoosFS struct {
	skill string
}

func (b *BoosFS) Work(hour int) (product string) {
	fmt.Printf("老板一天工作%d小时\n", hour)
	fmt.Printf("老板正在%s\n", b.skill)
	return "无逻辑需求"
}

func (b *BoosFS) Rest() {
	fmt.Println("老板休息")
}

func mainFS1() {
	// 创建由work组成的切片
	workers := make([]Worker, 0)

	// 向劳动大军中丢入worker的不同实现
	workers = append(workers, &Coder{"撸代码"})
	workers = append(workers, &ProductManager{"拍脑门"})
	workers = append(workers, &Boos{"喷空"})

	for _, worker := range workers {
		switch worker.(type) {
		case *Coder:
			fmt.Println("张三是个撸代码的")
		case *ProductManager:
			fmt.Println("李四是搞创意的，不要跟我聊逻辑")
		case *Boos:
			fmt.Println("我是给员工打工的")
		default:
			fmt.Println("不知道是什么鸟")
		}
	}
}

func mainFS2() {
	// 创建由work组成的切片
	workers := make([]Worker, 0)

	// 向劳动大军中丢入worker的不同实现
	workers = append(workers, &Coder{"撸代码"})
	workers = append(workers, &ProductManager{"拍脑门"})
	workers = append(workers, &Boos{"喷空"})

	for _, worker := range workers {
		if coder, ok := worker.(*Coder); ok {
			fmt.Println("发现一只程序员在", coder.skill)
			coder.WorkHome()
		} else {
			fmt.Println(worker, "不是程序员")
		}
	}
}
