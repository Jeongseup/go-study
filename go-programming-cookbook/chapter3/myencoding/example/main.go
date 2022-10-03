package main

import (
	"myencoding"
)

func main() {
	if err := myencoding.Base64Example(); err != nil {
		panic(err)
	}

	if err := myencoding.Base64ExampleEncoder(); err != nil {
		panic(err)
	}

	if err := myencoding.GobExample(); err != nil {
		panic(err)
	}
}
