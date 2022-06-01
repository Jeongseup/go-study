package main

import (
	"fmt"
)

func Caculate(x int) (result int) {
	result = x + 2
	return result
}
func main() {
	fmt.Println("Go Testing")

	result := Caculate(2)
	fmt.Println(result)
}
