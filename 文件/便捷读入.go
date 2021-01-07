package main

import (
	"fmt"
	"io/ioutil"
)

/*使用ioutil包对文件进行简易读取*/
func mainBJRead() {

	// 读入指定文件的全部数据，返回[]byte类型的原始数据
	file, err := ioutil.ReadFile("D:/aa.txt")
	if err == nil {
		str := string(file)
		fmt.Println(str)
	} else {
		fmt.Println("读取失败，err=", err)
	}
}
