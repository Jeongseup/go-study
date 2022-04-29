package main

import "fmt"

func exchange(x, y int) (int, int) {
	return y, x
}

func main() {
	x, y := 5, 10
	fmt.Println(exchange(x, y))
}
