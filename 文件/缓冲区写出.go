package main

import (
	"bufio"
	"fmt"
	"os"
)

/*以【创-写-追加】或【创-写-覆盖】方式打开一个文件，缓冲式写出几行数据，倒干缓冲区后退出*/
func mainWriter() {
	file, err := os.OpenFile("D:/aa.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败，err=", err)
		return
	}

	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()
	writer := bufio.NewWriter(file)
	writer.WriteString("新\n")
	writer.WriteString("年\n")
	writer.WriteString("快\n")
	writer.WriteString("乐\n")
	writer.Flush()
	fmt.Println("写出完毕")
}
