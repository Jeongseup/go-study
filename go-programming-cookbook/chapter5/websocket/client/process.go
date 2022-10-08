package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func process(c *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter some text: ")
		// ctrl-c를 막기 때문에 종료하면 ctrl-c를 누르고 ..
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read stdin", err)
		}

		// 읽어온 문자열에서 공백제거
		data = strings.TrimSpace(data)

		// 웹소켓에 메시지를 바이트로 작성
		err = c.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("failed to write message:", err)
			return
		}

		// 에코 서버이므로 서버에 요청했다가 받은 데이터를 작성한 후에 읽어올 수 있음
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("failed to read:", err)
			return
		}
		log.Printf("received back from server: %#v\n", string(message))
	}
}
