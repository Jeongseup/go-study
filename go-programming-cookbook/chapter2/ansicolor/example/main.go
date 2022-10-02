package main

import (
	"ansicolor"
	"fmt"
)

func main() {
	r := ansicolor.ColorText{
		TextColor: ansicolor.Red,
		Text:      "I'm red!",
	}
	fmt.Println(r.String())

	r.TextColor = ansicolor.Green
	r.Text = "Now I'm green!"
	fmt.Println(r.String())

	r.TextColor = ansicolor.Black
	r.Text = "Now I'm black!"
	fmt.Println(r.String())

	r.TextColor = ansicolor.ColorNone
	r.Text = "Back to normal..."

	fmt.Println(r.String())
}
