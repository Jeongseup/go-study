package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go!")

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// clone := [10]int{}
	// temp := [10]int{}

	// for i := 0; i < len(arr); i++ {
	// 	clone[i] = arr[i]
	// 	temp[i] = arr[len(arr)-1-i]

	// }

	// fmt.Println(clone)
	// fmt.Println(temp)

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println(arr)

	arr2 := [10]int{0, 2, 3, 4, 1, 3, 1, 7, 5, 4}
	temp := [10]int{}

	fmt.Println("before", arr2)

	for i := 0; i < len(arr2); i++ {
		idx := arr2[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr2[idx] = i
			idx++
		}
	}

	fmt.Println("after", arr2)
}
