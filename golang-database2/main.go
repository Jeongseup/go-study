package main

import (
	"bufio"
	"golang-database2/print"
	"golang-database2/read"

	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		print.PrintPrompt()
		read.ReaInput(reader)
	}
}
