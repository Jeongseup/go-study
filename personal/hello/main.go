package main

import (
  "fmt"

  "github.com/Jeongseup/go-study/personal/greeting"
)

func main() {
  message := greeting.Hello("John")
  fmt.Println(message)
}
