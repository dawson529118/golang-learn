package main

import "fmt"

func maindefer() {
	//defer lastPassage()
	defer func() {
		fmt.Println("疑是银河落九天")
		fmt.Println("一一李白")
	}()
	fmt.Println("日照香炉生紫烟")
	fmt.Println("遥看瀑布挂前川")
	fmt.Println("飞流直下三千尺")
}

func lastPassage() {
	fmt.Println("疑是银河落九天")
	fmt.Println("一一李白")
}
