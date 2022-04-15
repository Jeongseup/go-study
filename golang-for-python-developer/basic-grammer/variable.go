package main

import (
	"fmt"
	"reflect"
)

// const 전역변수
const a, b int = 1, 2

// 전역변수라서 할당할 필요가 없는 건가?

func main() {
	// 전역변수 사용
	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

	// 로컬변수
	var c = 3
	d := 4

	// zero value
	// 할당만 하면 그냥 default 값이 아래와 같이 잡힌다.

	var (
		e int
		f bool
		s string // ''
	)

	// 로컬변수 사용
	fmt.Println(c, d, e, f, s)
	// fmt.Println(s)
}
