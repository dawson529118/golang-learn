package main

import "fmt"

func mainMap() {
	//var scoreMap map[string]int = map[string]int{}
	//var scoreMap = map[string]int{}
	//scoreMap := map[string]int{}

	// 没有指定长度，长度为0
	scoreMap := make(map[string]int)

	// 添加键值对
	scoreMap["张三"] = 59
	scoreMap["李四"] = 63
	scoreMap["王五"] = 89

	// 根据键访问值
	fmt.Printf("张三的分数是：%d\n", scoreMap["张三"])
	// 修改值
	scoreMap["王五"] = 75
	fmt.Printf("王五修改后的分数是：%d\n", scoreMap["王五"])
	// 带校验地访问键值
	score, ok := scoreMap["李四"]
	fmt.Println(score, ok)

	score1, ok := scoreMap["赵六"]
	if ok {
		fmt.Println("赵六的成绩是：", score1)
	} else {
		fmt.Println("查无此人")
	}

}

func mainMap1() {
	scoreMap := make(map[string]int)

	scoreMap["张三"] = 59
	scoreMap["李四"] = 63
	scoreMap["王五"] = 89

	// 遍历key和value
	for key, value := range scoreMap {
		fmt.Printf("scoreMap[%s]=%d\n", key, value)
	}
}
