package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
	a := [5]int{32, 29, 78, 16, 81}

	for i, value := range a { // i에는 인덱스, value에는 배열 요소의 값이 들어감
		fmt.Println(i, value)
	}

	b := a
	b[2] = 0
	fmt.Println(a, b)
}
