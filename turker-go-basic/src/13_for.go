package main

import "fmt"

func main() {

	var i int // defaul 0이 할당
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("final i : ", i)

	var j int
	for {
		j++
		fmt.Println(j)

		if j == 100 {
			break
		}
	}
}
