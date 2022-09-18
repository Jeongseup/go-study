package main

import (
	"csvformat"
	"fmt"
	"log"
)

func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		log.Fatalln(err)
	}

	if err := csvformat.WriteCSVOutput(); err != nil {
		panic(err)
	}

	buffer, err := csvformat.WriteCSVBuffer()
	if err != nil {
		panic(err)
	}

	// 만약 여기서 반환된 버퍼를 저장하면 파일이 생성되는 거
	// 여기선 그냥 다시 위에 처럼 Stdout으로 출력

	fmt.Println("Buffer =  ", buffer.String())
}
