package main

import (
	"fmt"
	"mymath"
)

func main() {
	// mymath.Examples()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", mymath.Fib(i))
	}
	fmt.Println()
}
