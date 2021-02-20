package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*使用缓冲1K的缓冲区配合缓冲读写器进行图片拷贝*/
func mainBuffer() {
	// 打开源文件
	srcFile, _ := os.OpenFile("D:/gopath/src/myProject/images/test.jpg", os.O_RDONLY, 0666)
	// 打开目标文件
	dstFile, _ := os.OpenFile("D:/gopath/src/myProject/images/test111.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	// 延时关闭
	defer func() {
		srcFile.Close()
		dstFile.Close()
		fmt.Println("文件全部关闭")
	}()
	// 创建缓冲读取器
	reader := bufio.NewReader(srcFile)
	// 创建目标文件的写出器
	writer := bufio.NewWriter(dstFile)

	// 创建小水桶
	buffer := make([]byte, 1024)
	// 一桶一桶的读入数据，直到io.EOF
	for {
		_, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("源文件读取完毕！")
				break
			} else {
				fmt.Println("读取源文件错误，err=", err)
				return
			}
		} else {
			// 将小桶数据写出到目标文件
			_, err := writer.Write(buffer)
			if err != nil {
				fmt.Println("写出文件错误，err=", err)
				return
			}
		}
	}
	fmt.Println("拷贝完毕")
	// 将每桶数据写出到目标文件
}
