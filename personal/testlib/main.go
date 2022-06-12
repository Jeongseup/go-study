package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("")
	myJsonString := `{"some":"json"}`

	fmt.Println(myJsonString)
	fmt.Println([]byte(myJsonString))

	var myStoredVar string
	json.Unmarshal([]byte(myJsonString), &myStoredVar)
	fmt.Println(myStoredVar)
}
