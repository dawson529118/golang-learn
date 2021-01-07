package main

import (
	"fmt"
	"os"
	"time"
)

func mainFileOpenAndClose() {
	file, err := os.Open("D:/aa.txt")
	if err != nil {
		fmt.Println("文件打开失败，err=", err)
		return
	} else {
		fmt.Println("文件打开成功！！！")
	}

	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()
	fmt.Println("拿着文件进行操作", file)
	time.Sleep(1 * time.Second)
}
