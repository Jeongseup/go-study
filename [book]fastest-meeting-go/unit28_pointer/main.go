package main

import (
	"fmt"
)

func hello(n *int) {
	*n = 2
}

func main() {
	var numPointer *int = new(int) // new 함수로 공간 할당
	*numPointer = 1                // 역참조로 포인터형 변수에 값을 대입
	fmt.Println(*numPointer)       // 포인터형 변수에서 값을 가져오기

	var n int = 1
	hello(&n)
	fmt.Println(n)
}
