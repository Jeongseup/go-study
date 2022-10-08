package main

import (
	"fmt"
	"net/mail"
	"strings"
)

// printHeaderInfo 함수는 헤더 정보를 추출해 보기 좋게 출력
func printHeaderInfo(header mail.Header) {
	// 단일주소라는 것을 알기 때문에 다음 코드가 동작
	toAddress, err := mail.ParseAddress(header.Get("To"))
	if err == nil {
		fmt.Printf("To: %s <%s>\n", toAddress.Name, toAddress.Address)
	}
	fromAddress, err := mail.ParseAddress(header.Get("From"))
	if err == nil {
		fmt.Printf("From: %s <%s>\n", fromAddress.Name, fromAddress.Address)
	}

	fmt.Println("Subject:", header.Get("Subject"))

	// 다음 코드는 유효한 날짜에 동작
	if date, err := header.Date(); err == nil {
		fmt.Println("Date:", date)
	}

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println()
}
