package main

import "fmt"

func main() {

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#4bf745",
		"green": "#ff0000",
	}

	delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
