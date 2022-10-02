package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// WorkCount 함수는 파일을 입력받아 각 단어를 키로 하
// 단어의 등장 수는 값ㅡ로 하는 map을 반환
func WordCount(f io.Reader) map[string]int {
	result := make(map[string]int)

	// 파일 io Reader인터페이스에서 동작할 스캐너
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return result
}

func main() {
	fmt.Printf("string: number_of_occurrences\n\n")
	for key, value := range WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}
}
