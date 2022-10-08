package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:8888"

func echoBackCapitalized(conn net.Conn) {
	// conn reader execution config
	reader := bufio.NewReader(conn)

	// 읽어온 첫 줄을 가져옴
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data: %s\n", err.Error())
		return
	}

	// 출력한 다음 데이터를 다시 보낸다
	fmt.Printf("Received: %s", data)
	conn.Write([]byte(strings.ToUpper(data)))
	// 완료된 연결 종료
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Printf("listening on: %s\n", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("encountered an error accepting connection: %s\n", err.Error())
			// 오류가 있으면 재시도
			continue
		}
		// 이 작업을 비동기로 처리하면 잠재적으로 워커풀을ㅇ ㅟ해 좋은 사용사례가 될것?
		go echoBackCapitalized(conn)
	}

}
