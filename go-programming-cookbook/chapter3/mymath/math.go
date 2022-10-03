package mymath

import (
	"fmt"
	"math"
)

// Examples 함수는 math 패키지에서 제공하는 일부 함수의 사용법을 보여준다
func Examples() {
	// sqrt 예제
	i := 25

	result := math.Sqrt(float64(i))
	fmt.Printf("%v의 제곱근 %v\n", i, result)

	result = math.Ceil(9.5)
	fmt.Println(result)

	result = math.Floor(9.5)
	fmt.Println(result)

	fmt.Println("Pi:", math.Pi, "E:", math.E)
}
