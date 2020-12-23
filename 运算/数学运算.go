package main

import (
	"fmt"
	"math"
)

/*四舍五入、绝对值、乘方、开方*/
func mainMath() {
	fmt.Println(math.Round(3.4))
	fmt.Println(math.Round(3.5))
	/*先对绝对值四舍五入，再加负号*/
	fmt.Println(math.Round(-3.4))
	fmt.Println(math.Round(-3.5))

	//  绝对值
	fmt.Println(math.Abs(-3.4))

	// 乘方
	fmt.Println(math.Pow(2, 3))
	// 开方
	fmt.Println(math.Sqrt(9))

}
