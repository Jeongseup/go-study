package main

import "fmt"

func sample() (int, string) {
	return 0, "something" // 리턴이 두 개 이상이어도 괜찮다.
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
}
