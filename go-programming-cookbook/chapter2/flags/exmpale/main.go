package main

import (
	"flag"
	"flags"
	"fmt"
)

func main() {
	c := flags.Config{}
	c.Setup()

	flag.Parse()
	fmt.Println(c.GetMessage())
}
