package main

import (
	"fmt"
	"os"
)

func mainOs() {
	// 获得当前工作路径，默认当前工程根目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作路径失败")
	}
	fmt.Println("获取当前工作路径是：", dir)

	os.Environ()
	fmt.Println("临时文件夹：", os.TempDir())

	info, err := os.Stat("D:/qingyuan/record.txt")
	if err == nil {
		// 是否文件夹
		fmt.Println(info.IsDir())
		// 模式：读写执行权限
		fmt.Println(info.Mode())
		// 修改时间
		fmt.Println(info.ModTime())
		// 文件大小
		fmt.Println(info.Size())
	}
}
