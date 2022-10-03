package main

import (
	"collections"
	"fmt"
)

func main() {
	ws := []collections.WorkWith{
		collections.WorkWith{"Example", 1},
		collections.WorkWith{"Example2", 2},
	}

	fmt.Printf("Initial list: %#v\n", ws)

	// 먼저 리스트의 값을 소문자로 만든다
	ws = collections.Map(ws, collections.LowerCaseData)
	fmt.Printf("After LowerCaseData Map: %#v\n", ws)

	// 그 다음 모든 버젼을 증가시킨다
	ws = collections.Map(ws, collections.IncrementVersion)
	fmt.Printf("After IncrementVersion Map: %#v\n", ws)

	// 마지막으로 3보다 작은 모든 버전을 제거
	ws = collections.Filter(ws, collections.OldVersion(3))
	fmt.Printf("After OldVersion Filter: %#v\n", ws)
}
