package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
以只读的方式打开一个文件，创建其带缓冲的读取器，逐渐读取到末尾
4=readable,2=writable,1=executable
6=4+2
*/
func mainHCFile() {
	// 以只读模式打开文件（0666代表所有人都有读写权限）
	file, err := os.OpenFile("D:/aa.txt", os.O_RDONLY, 0666)

	// 判断文件打开是否成功
	if err != nil {
		fmt.Println("打开文件失败, err=", err)
		return
	} else {
		fmt.Println("打开文件成功")
	}

	// 延时（函数返回前）关闭文件
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()

	// 创建该文件的缓冲读取器
	reader := bufio.NewReader(file)

	// 循环读入数据
	for {
		// 每次读入一行
		str, err := reader.ReadString('\n')

		// 判断读入是否成功
		if err == nil {
			fmt.Println(str)
		} else {
			if err == io.EOF {
				fmt.Println("文件已到末尾")
				break
			} else {
				fmt.Println("读取失败，err=", err)
				return
			}
		}
	}
	fmt.Println("文件读取完毕")
}
