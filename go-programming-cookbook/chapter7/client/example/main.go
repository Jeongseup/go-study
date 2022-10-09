package main

import (
	"client"
)

func main() {
	cli := client.Setup(true, false)
	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}
	// fmt.Println(&cli)

	if err := client.DoOps(cli); err != nil {
		panic(err)
	}

	c := client.Controller{Client: cli}
	if err := c.DoOps(); err != nil {
		panic(err)
	}

	// Setup 함수의 첫 번째 인자를 true, true tjfwjd
	client.Setup(true, true)
	// fmt.Println(&cli)

	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}
}
