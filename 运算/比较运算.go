package main

import "fmt"

func mainBi() {

	// true false
	fmt.Println((1+1 == 2))
	fmt.Println((1+1 == 3))

	// false true
	fmt.Println((1+1 != 2))
	fmt.Println((1+1 != 3))

	// true false
	fmt.Println((2 + 3) > 4)
	fmt.Println((2 + 3) < 4)

	// true true
	fmt.Println((2 + 3) >= 5)
	fmt.Println((2 + 3) <= 5)
}
