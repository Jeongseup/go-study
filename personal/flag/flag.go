package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("")

	// word라는 문자열 플래그 선언
	// foo라는 기본값을 가짐
	// a string이 설명?
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPrt := flag.Bool("fork", false, "a bool")

	// 다른 곳에서 선언된 변수 참조
	var stringValue string
	// fmt.Println("stringValue", stringValue) => 공백

	flag.StringVar(&stringValue, "stringValue", "bar", "a string var")

	// 모든 플래그가 선언되면 커맨드라인 파싱을 실행하기 위해 flag.Parse()호출
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPrt)
	fmt.Println("stringValue", stringValue)
	fmt.Println("tail", flag.Args())

	/*
		example) $ ./flag -h

			Usage of ./flaasdag:
			  -fork
			        a bool
			  -numb int
			        an int (default 42)
			  -stringValue string
			        a string var (default "bar")
			  -word string
			        a string (default "foo")
	*/
}
