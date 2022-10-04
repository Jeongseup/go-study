package mylog

import (
	"bytes"
	"fmt"
	"log"
)

// Log함수는 로거 설정을 사용한다
func Log() {
	// bytes.Buffer에 기록하기 위해 로거를 설정
	buf := new(bytes.Buffer)

	// 두번째 인자는 접두사이며 마지막 인자는 논리 연산자 or를 결합해 설정
	logger := log.New(buf, "logger: ", log.Lshortfile|log.Ldate)
	logger.Println("test")

	logger.SetPrefix("new logger : ")

	logger.Printf("you can also add args(%v) and use Fatalln to log and crash", true)

	fmt.Println(buf.String())
}
