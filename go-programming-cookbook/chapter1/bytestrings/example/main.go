package main

import (
	"bytestrings"
)

func main() {
	// fmt.Println("")
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
