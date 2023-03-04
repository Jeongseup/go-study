package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5, 10) // 길이가 5이면 a[0], a[1], a[2], a[3], a[4]가 생성

	fmt.Println(len(a), cap(a))

	// fmt.Println(a[4]) // 0: make 함수를 사용하면 슬라이스의 요소는 모두 0으로 초기화
	// fmt.Println(a[5]) // 길이를 벗어난 인덱스에 접근했으므로 런타임 에러 발생
	// fmt.Println(a[8]) // 길이를 벗어난 인덱스에 접근했으므로 런타임 에러 발생

	a = append(a, 4, 5, 6)
	fmt.Println(a)

	a2 := []int{1, 2, 3, 4, 5}
	b2 := make([]int, 3) // make 함수로 공간을 할당

	copy(b2, a2) // 슬라이스 a의 요소를 슬라이스 b에 복사
	b2[0] = 9999

	fmt.Println(a2) // [1 2 3 4 5]
	fmt.Println(b2) // [1 2 3]: 슬라이스 b의 길이는 3이므로 a의 요소 3개만 복사됨
}
