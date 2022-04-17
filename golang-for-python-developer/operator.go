package main

import "fmt"

func sample() (int, string) {
	return 0, "something" // 리턴이 두 개 이상이어도 괜찮다.
}

// 데이터 타입을 뒤에 써준다.
func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func multiplied(a, b int) int {
	return a * b
}

func divided(a, b int) int {
	return a / b
}

func plusViod(a, b int) {
	// defer fmt.Println("plus viod ... ")
	defer func() {
		fmt.Println("plus function was ended")
		fmt.Println("plus function was ended2")
	}()

	fmt.Println(a + b)
}

func main() {
	a := 8
	b := 4

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a -= 2 // 8 - 2
	b += 3 // 4 + 3
	fmt.Println(a, b)

	c := 6
	d := 2

	c *= 3 // 6 * 3
	d /= 2 // 2 ÷ 2
	fmt.Println(c, d)

	// string
	s := "someth"
	s += "ing"
	fmt.Println(s)

	fmt.Println(sample())

	fmt.Println(plus(a, b))
	fmt.Println(minus(a, b))
	fmt.Println(multiplied(a, b))
	fmt.Println(divided(a, b))

	plusViod(a, b)
}
