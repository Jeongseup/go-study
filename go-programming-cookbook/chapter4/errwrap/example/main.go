package main

import (
	"errwrap"
	"fmt"
)

func main() {
	errwrap.Wrap()
	fmt.Println()

	errwrap.Unwrap()
	fmt.Println()

	errwrap.StackTrace()
}
