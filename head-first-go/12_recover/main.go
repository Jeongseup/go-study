package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover() // recover를 호출하고 반환 값을 저장
	if p == nil {
		return
	}

	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		// 패닉값이 에러타입이 아닌 경우, 패닉을 다시 한번 더 일으킴
		panic(p)
	}
}

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// fmt.Printf("Returning error from scanDirectory(\"%s\" call\n", path)
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	defer reportPanic()
	// panic("some other issue")
	scanDirectory("go")
}
