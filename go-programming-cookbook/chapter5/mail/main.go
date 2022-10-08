package main

import (
	"io"
	"log"
	"net/mail"
	"os"
	"strings"
)

// 이메일 메세지의 예
// an example email message
const msg string = `Date: Thu, 24 Jul 2019 08:00:00 -0700
From: Aaron <fake_sender@example.com>
To: Reader <fake_receiver@example.com>
Subject: Gophercon 2019 is going to be awesome!

Feel free to share my book with others if you're attending.
This recipe can be used to process and parse email information.
`

func main() {
	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	printHeaderInfo(m.Header)

	// 헤더를 출력한 다음 표준 출력에 본문을 보낸다
	if _, err := io.Copy(os.Stdout, m.Body); err != nil {
		log.Fatal(err)
	}
}
