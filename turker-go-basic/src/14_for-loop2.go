package main

import "fmt"

func gogodan() {
	for dan := 1; dan <= 9; dan++ {
		fmt.Printf("%d 단 \n", dan)

		for j := 1; j <= 9; j++ {
			fmt.Printf("%d * %d = %d\n", dan, j, j*dan)
		}

		fmt.Println("==========")
	}
}

func star1() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			// fmt.Print(j)
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var cnt int = 1

	for i := 0; i < 5; i++ {
		// fmt.Print(cnt)

		for j := 0; j < cnt; j++ {
			fmt.Print("*")
		}
		fmt.Println()

		if i < 2 {
			cnt++
		} else {
			cnt--
		}
	}
}
