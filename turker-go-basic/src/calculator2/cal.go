package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Enter a number\n")

	// input value at the 'reader' var
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	// fmt.Printf("입력하신 숫자는 %d, %d.", n1, n2)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1 + n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1 + n2)
	} else if line == "*" {
		fmt.Printf("%d * %d = %d", n1, n2, n1 + n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1 + n2)
	} else {
		fmt.Printf("잘못 입력")
		
	}
}