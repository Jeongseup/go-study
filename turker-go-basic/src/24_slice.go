package main

import (
	"fmt"
)

func main() {
	fmt.Println("")

	// 빈 공간에 동적배열
	a := []int{1, 2, 3, 4, 5}

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("len(a) = %d\n", cap(a))

	fmt.Printf("Address of slice = %v: %p\n", a, &a)

	a = append(a, 6)
	fmt.Printf("Address of slice = %v: %p\n", a, &a)
	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("len(a) = %d\n", cap(a))

	b := make([]int, 0, 8)
	fmt.Printf("Address of slice = %v: %p\n", b, &b)

	// 이건 길이
	fmt.Printf("len(a) = %d\n", len(b))
	// 이건 캐파(복사하여 메모리 위치를 바꾸지 않고 appending 가능한 길이)
	fmt.Printf("len(a) = %d\n", cap(b))

	for i := 0; i < len(a); i++ {
		fmt.Println(&a[i])
	}

	b = append(b, 1)

	fmt.Printf("len(a) = %d\n", len(b))
	fmt.Printf("len(a) = %d\n", cap(b))

	fmt.Printf("Address of slice = %v: %p\n", b, &b)

	beforeC := make([]int, 2, 4)
	beforeC[0] = 1
	beforeC[1] = 2

	afterC := make([]int, len(beforeC))

	for i := 0; i < len(beforeC); i++ {
		afterC[i] = beforeC[i]
	}

	afterC = append(afterC, 3)

	fmt.Printf("%p %p \n", beforeC, afterC)
}
