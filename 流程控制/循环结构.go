package main

import (
	"fmt"
	"time"
)

/*有限次循环*/
func mainIf() {
	/*for i := 1; i < 10 ;i++  {
		fmt.Println(i)
	}*/

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
	fmt.Println("GAME OVER！")
}

/*无限次循环*/
func mainInfinite() {
	for {
		fmt.Println("我爱撸代码")
		// 一秒执行一次
		time.Sleep(1 * time.Second)
	}
}

/*退出循环*/
func mainExit() {
	var count = 0
	for {
		fmt.Println("我爱撸代码")
		count++
		time.Sleep(500 * time.Millisecond)

		if count > 5 {
			break
		}
	}
	fmt.Println("game over")
}

func mainDemo() {
	var count = 1
	for {
		fmt.Println("我爱撸代码", count)
		time.Sleep(5 * time.Millisecond)
		count++
		if count >= 19 {
			// 结束循环
			break
		}
		if count%5 == 0 {
			// 跳过当前循环，继续下一循环
			continue
		}
		fmt.Println("GAME OVER!")
	}
}
