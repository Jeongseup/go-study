package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cnt := 0
	for {
		cnt++

		// 무작위 숫자 3개를 만든다.
		numbers := MakeNumbers()
		// 사용자로부터 숫자 3개를 받는다.
		inputNumbers := InputNumbers()
		// 값을 비교한다.
		result := CompareNumbers(numbers, inputNumbers)
		// 결과 출력
		PrintResult(result)

		if IsEndGame(result) {
			// 게임 끝
			break
		}
	}

	fmt.Printf("%d번만에 맞췄습니다", cnt)
}

// 0~9 사이의 곂치지 않는 숫자 반환
func MakeNumbers() [3]int {
	// set으로 중복방지
	tempSet := map[int]bool{}
	for {
		n := rand.Intn(10)
		_, present := tempSet[n] // check for existence

		// 없으면 skip 있으면 추가
		if !present {
			// add
			tempSet[n] = true
			// 추가하고 나서 길이가 3이면 스킾
			if len(tempSet) == 3 {
				break
			}
		}
	}

	// fmt.Print(tempSet)
	rst := [3]int{}

	idx := 0
	for k := range tempSet {
		// fmt.Println(i, k)
		rst[idx] = k
		idx++
	}

	fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {

	var input [3]int
	return input
}

func CompareNumbers(numbers, inputNumbers [3]int) bool {
	return true
}

func PrintResult(result bool) {
	fmt.Println(result)
}

func IsEndGame(result bool) bool {
	return true
}
