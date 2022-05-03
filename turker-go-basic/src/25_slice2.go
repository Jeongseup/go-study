package main

import (
	"fmt"
)

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}

func main() {
	fmt.Println("")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 4 ~ 8 - 1까지
	b := a[4:8]
	// 인덱스 4부터
	c := a[4:]

	// b를 바꾸더라도 b가 메모리를 참조해서 값이 바뀐다.
	b[0] = 0
	b[1] = 1

	fmt.Println(a, len(a))
	fmt.Println(b, len(b))
	fmt.Println(c, len(c))

	for i := 0; i < 5; i++ {
		var back int

		a, back = RemoveBack(a)
		fmt.Printf("%d\n", back)
	}
	fmt.Println(a)
}
