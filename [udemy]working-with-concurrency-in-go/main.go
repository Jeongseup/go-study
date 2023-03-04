package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
	}

	wg.Add(len(words))
	for i, v := range words {

		go printSomething(fmt.Sprintf("%d: %s", i, v), &wg)
	}
	wg.Wait()
	// printSomething("the end")
	fmt.Println("the end")
}
