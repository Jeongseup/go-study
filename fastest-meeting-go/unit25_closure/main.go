package main

import (
	"fmt"
)

func main() {
	//      ↓ 익명 함수를 선언 및 정의
	sum := func(a, b int) int {
		return a + b
	}

	r := sum(1, 2) // 익명함수 사용

	fmt.Println(r) // 3

	a, b := 3, 5
	f := func(x int) int {
		return a*x + b
	}

	y := f(5)
	fmt.Println(y)

	myFun := calc()
	fmt.Println(myFun(1))
}

//              ↓ 리턴 값이 익명 함수
func calc() func(x int) int {
	a, b := 3, 5 // 지역 변수는 함수가 끝나면 소멸되지만
	return func(x int) int {
		return a*x + b // 클로저이므로 함수를 호출 할 때마다 변수 a와 b의 값을 사용할 수 있음
	}
	// ↑ 익명 함수를 리턴
}
