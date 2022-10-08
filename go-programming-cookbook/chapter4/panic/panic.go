package panic

import (
	"fmt"
	"strconv"
)

// 패닉 함수는 0으로 나누기로 인해 패닉 발생
func Panic() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}
	a := 1 / zero
	fmt.Println("we'll never get here", a)
}

// Catcher 는 패닉함수 호출
func Catcher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred", r)
		}
	}()

	Panic()
}
