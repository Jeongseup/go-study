package main

import "fmt"

func main() {

	// int : 4type => 10 length array : 4 * 10 = 40byte
	var A [10]int

	fmt.Println(A)

	for i := 0; i < len(A); i++ {
		A[i] = i * i
	}

	fmt.Println(A)

	// Go lang is accpeted UTF8 encoding
	// so String which has only English is that each rune has 1byte size

	s := "hello world"
	hanguel := "hello 월드"

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], ", ")
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), ", ")
	}
	fmt.Println()
	fmt.Println("hanguel len is", len(hanguel))
	for i := 0; i < len(hanguel); i++ {
		fmt.Print(hanguel[i], ", ")
	}
	fmt.Println()

	// 이럴 때 필요한게 UTF8 문자를 나타낼 수 있는 rune type
	// rune is one of go lang type for representing character
	// but, it has size from 1byte to 3type for UTF8
	hanguelRune := []rune(hanguel)
	fmt.Println("hanguelRune len is", len(hanguelRune))
	for i := 0; i < len(hanguelRune); i++ {
		fmt.Print(hanguelRune[i], ", ")
	}
	fmt.Println()
	for i := 0; i < len(hanguelRune); i++ {
		fmt.Print(string(hanguelRune[i]), ", ")
	}
}
