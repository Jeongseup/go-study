package main

import (
	"fmt"
)

func f() {
	defer func() {
		s := recover() // 패닉이 발생해도 프로그램 종료 X
		fmt.Println(s)
	}()

	// panic("Error!!") // 패닉 발생
	a := [...]int{1, 2}
	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	f()
	fmt.Println("hello world")
}
