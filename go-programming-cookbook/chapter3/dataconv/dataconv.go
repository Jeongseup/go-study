package dataconv

import "fmt"

// ShowConv함수는 일부 타입의 변환 예제를 보여줌
func ShowConv() {
	//	int
	var a = 24
	// float
	var b = 2.0

	c := float64(a) * b
	fmt.Println(c)
	// fmt.Sprintf는 문자열로 변환하는 좋은 방법이다
	precision := fmt.Sprintf("%.2f", b)

	// 값과 타입 출력
	fmt.Printf("%s - %T\n", precision, precision)
}
