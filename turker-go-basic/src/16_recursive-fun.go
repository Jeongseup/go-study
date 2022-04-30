package main

import "fmt"

func main() {
	// s := sum(10, 0)
	s := fibo(10)
	fmt.Println(s)
}

func fibo(x int) int {
	if x == 0 || x == 1 {
		return 1
	}

	return fibo(x-1) + fibo(x-2)
}

func sum(x, s int) int {
	if x == 0 {
		return s
	}

	s += x
	fmt.Println(s)
	return sum(x-1, s)
}
