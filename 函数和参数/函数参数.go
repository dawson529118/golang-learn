package main

import "fmt"

func mainParam() {
	//sayLove("拉芳", 4)
	sayLoveToThem("拉芳", "拉草", "潘婷")
}

func sayLove(name string, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("我爱%s\n", name)
	}
}

func sayLoveToThem(names ...string) {
	fmt.Println("人生苦短，必须性感")
	fmt.Println("我爱", names)
}

func sayLoveToThem1(names ...string) {
	for i, name := range names {
		fmt.Printf("我爱No.%d：%s\n", i, name)
	}
}
