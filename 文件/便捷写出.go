package main

import (
	"fmt"
	"io/ioutil"
)

// 使用ioUtil包进行简易文件写出
func mainGoWriter() {
	dataStr := `忘
			恩
				负
					义`
	fmt.Printf("data type=%T, value=%v\n", dataStr, dataStr)
	dataBytes := []byte(dataStr)
	err := ioutil.WriteFile("D:/aa.txt", dataBytes, 0666)
	if err != nil {
		fmt.Println("写出文件失败")
	} else {
		fmt.Println("写出文件成功")
	}
}
