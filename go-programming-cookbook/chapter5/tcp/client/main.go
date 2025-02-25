package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "localhost:8888"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// 클라이언트로부터 문자열 입력을 가져온다
		fmt.Println("Enter some text: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("encounterd an error reading input: %s\n", err.Error())
			continue
		}
		// addr에 연결한다
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("encountered an error connecting: %s\n", err.Error())
		}
		// 수립된 연결에 데이터를 쓴다?
		fmt.Fprintf(conn, data)
		// 응답을 다시 읽어옴
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reading response: %s\n", err.Error())
		}
		fmt.Printf("Received back: %s\n", status)
		// 완료된 연결 종료
		conn.Close()
	}
}
