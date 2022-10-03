package main

import (
	"fmt"
	"nulls"
)

func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := nulls.NullEncofing(); err != nil {
		panic(err)
	}
	fmt.Println()
}
