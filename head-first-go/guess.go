package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	target := rand.Intn(100) + 1
	// fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	success := false
	for count := 0; count < 10; count++ {
		fmt.Println("You have", 10-count, "counts left")

		fmt.Print("Make a guess :")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("your guess was LOW")
		} else if guess > target {
			fmt.Println("your guess was HIGH")
		} else {
			success = true
			fmt.Println("Good, You guessed it")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was :", target)
	}
}
