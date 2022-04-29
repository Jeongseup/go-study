package main

import "fmt"

func main() {
	a := 4
	b := 2

	// var c int
	// c = 5
	// var d = 6

	// var e = 3.14
	// var f int = 8

	fmt.Printf("4와 2의 AND연산 : %v\n", a&b)
	fmt.Printf("4와 2의 OR연산 %v\n", a|b)
	fmt.Printf("4와 2의 XOR연산 : %v\n", a^b)

	////////////////////////////////////
	a2 := 21
	c := a2%10
	a2 = a2 / 10
	d := a2%10
	
	fmt.Printf("첫번째 수 : %v, 두번째 수 : %v\n", c, d)
	// a&^a => clear 연산 => a와 not b의 AND 연산 => 다 0이 됨.
	// shift 연산
	// >> , <<
}