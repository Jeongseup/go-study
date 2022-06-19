package main

import (
	"fmt"
	"os"
)

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func HelloWorld() {
	defer func() {
		fmt.Println("seup") // HelloWorld() 함수가 끝나기 직전에
	}()

	hello()
}

func main() {
	// 함수의 호출이 지연됨.
	defer world() //현재 함수 main()이 끝나기 직전에 호출 // 맨위에 된게 제일 나중
	defer func() {
		fmt.Println("finally defer! ")
	}()

	hello()
	hello()
	HelloWorld()

	// 자료구조의 스택과 동일
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	ReadTxt()
}

func ReadTxt() {
	file, err := os.Open("hello.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return // file.Close() 호출
	}

	buf := make([]byte, 100)
	if _, err := file.Read(buf); err != nil {
		fmt.Println(err)
		return // file.Close() 호출
	}

	fmt.Println(string(buf))
	// file.Close() 호출
}
