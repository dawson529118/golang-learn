package main

import "fmt"

type Person struct {
	// 封装结构体的属性
	name  string
	age   int
	sex   bool
	hobbl []string
}

/*
封装结构体的方法
- 无论方法的主语定义为值类型还是指针类型，对象值和对象指针都能正常访问
- 通常会将主语定义为指针类型
*/
func (p *Person) Eat() {
	fmt.Printf("%s爱饕餮\n", p.name)
}

func (p *Person) Drank() {
	fmt.Printf("%s爱喝酒\n", p.name)
}

func (p *Person) SelfInt() {
	fmt.Printf("我是：%s，今年%d\n", p.name, p.age)
}

func MakeHimEat(p Person) {
	fmt.Printf("person的真身地址是：%p\n", &p)
	p.Eat()
	p.age -= 1
}

func MakeHimEatEat(p *Person) {
	p.Eat()
	p.age -= 1
}

func mainEnca() {

	// 创建对象/实例
	person := Person{}
	// 设置属性
	person.name = "张三"
	// 访问其方法
	person.Eat()
	person.Drank()
}

func mainEncaOne() {
	// 创建对象时，给指定属性赋值
	//person := &Person{name: "张三", age: 32}
	person := &Person{"张三", 32, true, []string{"撸代码", "面向对象"}}
	person.Eat()
	person.Drank()
	person.SelfInt()
}

func mainTwo() {
	person := Person{"张三", 24, true, []string{}}

	// 根据要求进行传值：1、值传递；2、指针传递
	//MakeHimEat(person)
	//MakeHimEatEat(&person)

	// 值传递只是传递的是副本；指针传递传递的才是真身
	for i := 0; i < 7; i++ {
		MakeHimEatEat(&person)
	}
	fmt.Printf("吃过之后，年龄是：%d\n", person.age)
}
