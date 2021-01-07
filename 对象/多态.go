package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
劳动者父类接口
内含两个抽象方法：工作、休息
*/
type Worker interface {
	// 每天工作多少小时，产出何种产品
	Work(hour int) (product string)
	// 休息
	Rest()
}

/*定义码农结构体*/
type Coder struct {
	skill string
}

/*码农指针实现worker接口*/
func (c *Coder) Work(hour int) (product string) {
	fmt.Printf("码农一天工作%d小时\n", hour)
	fmt.Printf("码农正在%s\n", c.skill)
	return "BUG"
}

func (c *Coder) Rest() {
	fmt.Println("休息是什么？？")
}

/*码农特有的方法*/
func (c *Coder) WorkHome() {
	fmt.Println("程序员在家工作")
}

/*定义产品经理结构体*/
type ProductManager struct {
	skill string
}

/*PriductManger指针实现worker接口*/
func (pm *ProductManager) Work(hour int) (product string) {
	fmt.Printf("产品经理一天工作%d小时\n", hour)
	fmt.Printf("产品正在%s\n", pm.skill)
	return "无逻辑需求"
}

func (pm *ProductManager) Rest() {
	fmt.Println("产品经理休息")
}

/*定义老板结构体*/
type Boos struct {
	skill string
}

func (b *Boos) Work(hour int) (product string) {
	fmt.Printf("老板一天工作%d小时\n", hour)
	fmt.Printf("老板正在%s\n", b.skill)
	return "无逻辑需求"
}

func (b *Boos) Rest() {
	fmt.Println("老板休息")
}

func mainDT() {
	// 创建由work组成的切片
	workers := make([]Worker, 0)

	// 向劳动大军中丢入worker的不同实现
	workers = append(workers, &Coder{"撸代码"})
	workers = append(workers, &ProductManager{"拍脑门"})
	workers = append(workers, &Boos{"喷控"})

	/*随机生成今天星期几*/
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	weekDay := r.Intn(7)
	fmt.Printf("今天星期%d\n", weekDay)

	// 工作日全体工作
	if weekDay > 0 && weekDay < 6 {
		// 全体工作
		for _, worker := range workers {
			worker.Work(8)
		}
	}

}
