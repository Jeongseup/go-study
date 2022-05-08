package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var myLogger *log.Logger

func main() {
	fmt.Println("")
	// logger setting
	// 로그파일 오픈 -> 없으면 자동으로 생성해주네?
	// param : (파일포인터 지정 / 로그 플래그..? 자동으로 생성, )
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	// 파일과 화면에 같이 출력
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	// 표준로거를 파일로그로 변경
	log.SetOutput(multiWriter)

	// param : (io.Writer 인터페이스, 출력 인터페이스 / prefix / 표준 플래그, 날짜 플래그, 시간 플래그 등.. )
	// myLogger = log.New(os.Stdout, "INFO : ", log.LstdFlags)
	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	run()

	// 개발자가 별도의 logger를 생성하지 않으면 표준 logger를 자동으로 사용
	log.Println("Logging")
	myLogger.Printf("End of Program\n\n")
}

func run() {
	log.Println("Before Test")
	myLogger.Print("Test")
	log.Println("After Test")
}
