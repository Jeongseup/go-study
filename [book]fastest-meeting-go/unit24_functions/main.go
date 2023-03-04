package main

import "fmt"

func SumAndDiff(a int, b int) (sum int, diff int) { // 리턴값을 각각 sum, diff로 이름을 정함
	sum = a + b  // 리턴값 변수 sum에 합 대입
	diff = a - b // 리턴값 변수 diff에 차 대입
	return
}

func mySum(n ...int) int { // int형 가변인자를 받는 함수 정의
	total := 0
	for _, value := range n { // range로 가변인자의 모든 값을 꺼냄
		total += value // 꺼낸 값을 모두 더함
	}

	return total
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	var hello func(a int, b int) (int, int) = SumAndDiff
	numbers := []int{1, 2, 3, 4, 5}
	r := mySum(numbers...)

	fmt.Println(hello(3, 4)) // 8 4
	fmt.Println(r)
	fmt.Println(factorial(5))

	// 익명함수
	func() {
		fmt.Println("Hello, Anonymous")
	}()

	func(s string) {
		fmt.Println(s)
	}("seup")

	myFun1 := func(a int, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(myFun1)

	// 클로저 closure

}
