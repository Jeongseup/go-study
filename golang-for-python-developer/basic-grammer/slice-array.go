package main

import (
	"fmt"
)

func main() {
	// slice 초기화
	// golang에는 슬라이스와 배열이 있다.
	// slice가 파이썬 리스트
	slice := []string{"something1"}
	// fmt.Println(slice)

	// slice 값 추가
	slice = append(slice, "something2")

	// fmt.Println(slice)
	slice = append(slice, "third")
	// fmt.Println(slice)

	// slice 값 제거
	slice = slice[1:]

	// fmt.Println(slice)
	// fmt.Println(slice[0])

	// array 초기화
	array := [3]string{"something1"}
	fmt.Println(array)
	fmt.Println(len(array))

	// array = append(array, "test") => append must be a slice

	// array 값 추가
	array[1] = "something2"

	// array 값 제거
	array[0] = ""

	// fmt.Println(array)
	// fmt.Println(array[1])

	var a []int
	b := make([]int, 0) // 이게 초기값을 할당해준거라는데.. 왜 print에는 출력이 안될까?
	c := []int{}
	d := make([]int, 0, 1)

	fmt.Println(a,
		a == nil,

		b,
		b == nil,

		c,
		d)

	// slice와 array nil 비교
	var foo1 []int
	// var foo2 []int

	fmt.Println(foo1 == nil)
	// fmt.Println(foo1 == foo2) => error : slicie can only be compared to nil not other slice

	// [3]string{"something1"}
	var foo3 = [2]int{}
	fmt.Println(foo3)

	var foo4 = [2]int{}
	fmt.Println(foo4)

	fmt.Println(foo3 == foo4) // slice가 아닌 array끼리는 == 비교연산자 가능.!

}
