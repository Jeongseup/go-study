package main

import (
	"fmt"
	"net"
)

const addr = "localhost:8888"

func main() {
	conns := &connections{
		addrs: make(map[string]*net.UDPAddr),
	}

	fmt.Printf("serving on %s\n", addr)

	// udp addr 생성
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}

	// 지정된 주소로 요청 대기
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	go broadcast(conn, conns)
	msg := make([]byte, 1024)
	for {
		// 메시지를 다시 보내기 위한 주소와 포트 정보의 수집을 위해
		// 메시지 수신
		_, retAddr, err := conn.ReadFromUDP(msg)
		if err != nil {
			continue
		}
		// 수신한 메시지를 맵에 저장
		conns.mu.Lock()
		conns.addrs[retAddr.String()] = retAddr
		conns.mu.Unlock()
		fmt.Printf("%s connected\n", retAddr)
	}
}
