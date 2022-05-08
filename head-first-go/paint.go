package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := errors.New("height can't be negative")
	// fmt.Println(err.Error())
	log.Fatal(err)
	fmt.Println(err)
}
