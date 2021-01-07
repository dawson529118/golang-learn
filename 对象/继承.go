package main

import "fmt"

type PersonJC struct {
	// 封装结构体的属性
	name  string
	age   int
	sex   bool
	hobbl []string
}

func (p *PersonJC) Eat() {
	fmt.Printf("%s爱饕餮\n", p.name)
}

func (p *PersonJC) Drank() {
	fmt.Printf("%s爱喝酒\n", p.name)
}

/*死堆代码的*/
type coder struct {
	// 持有一个父类声明--继承了person
	Person
	// 会的语言
	langs []string
}

func (c *coder) Code() {
	fmt.Printf("%s会%v，正在堆代码", c.name, c.langs)
}

type driver struct {
	Person
	jiazhaoId string
	isDriving bool
}

/*司机的特有方法*/
func (d *driver) drive() {
	fmt.Printf("%s一言不合就开车\n", d.name)
}

/*覆写父类方法*/
func (d *driver) Drank() {
	if !d.isDriving {
		fmt.Printf("%s爱喝酒\n", d.name)
	} else {
		fmt.Println("fuckoff，司机一滴酒亲人两行泪")
	}
}

func mainJC() {
	/*c := new(coder)
	c.name = "张三"
	c.langs = []string{"Go", "Java"}
	c.Drank()
	c.Code()*/
	dPtr := new(driver)
	dPtr.name = "李四"

	dPtr.jiazhaoId = "京·34NKL6"
	dPtr.drive()
	dPtr.isDriving = false
	dPtr.Drank()
}
