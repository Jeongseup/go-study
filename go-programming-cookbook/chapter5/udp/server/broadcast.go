package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type connections struct {
	addrs map[string]*net.UDPAddr
	// 맵의 수정을 위해 lock시킨다
	mu sync.Mutex
}

func broadcast(conn *net.UDPConn, conns *connections) {
	count := 0
	for {
		// 새로운 루프가 돌면 생기면 카운터를 올리고
		count++
		// 락
		conns.mu.Lock()
		// 알려진 주소에 대해 루프로 반복처리
		for _, retAddr := range conns.addrs {
			// 모두에게 메세지 전송 루프
			msg := fmt.Sprintf("Sent %d", count)
			if _, err := conn.WriteToUDP([]byte(msg), retAddr); err != nil {
				fmt.Printf("error encountered: %s\n", err.Error())
				continue
			}
		}
		conns.mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}
