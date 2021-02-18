package main

import (
	"fmt"
	"io/ioutil"
)

func mainGoCopy() {
	bytes, err := ioutil.ReadFile("D:/fu.png")
	if err != nil {
		fmt.Println("读取文件失败！")
		return
	}
	err = ioutil.WriteFile("D:/fu2222.png", bytes, 0666)
	if err != nil {
		fmt.Println("拷贝失败！")
	} else {
		fmt.Println("拷贝成功！")
	}
	fmt.Println("执行完毕！")
}
