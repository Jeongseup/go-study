package main

import (
	"fmt"
	"something"

	"github.com/fatih/color"
)

func main() {
	something.Something()
	fmt.Println("그냥 글씨")
	color.Green("초록색 글씨")
}
