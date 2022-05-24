package main

import (
	"fmt"

	"github.com/Jeongseup/go-study/personal/greeting"
)

func main() {
	msg := greeting.Hello("seup")
	fmt.Println(msg)
}
