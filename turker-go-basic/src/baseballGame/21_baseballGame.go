package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

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

// 키보드로부터 0~9 사의 겹치지 않는 숫자 3개 입력받아 반환
func InputNumbers() [3]int {
	var input [3]int

	for {
		var num int

		_, err := fmt.Scanf("%d\n", &num)
		if err != nil { // 에러가 아무것도 아닌게 아닌 경우 => 에러가 존재하는 경우
			fmt.Println("error occur")
			continue

		}

		// fmt.Println(&num, num)
		if num > 999 {
			fmt.Println("네 자리 이상 입력")
			continue
		}

		success := true

		idx := 0
		for num > 0 {
			n := num % 10
			num = num / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if input[j] == n {
					//	겹침 다시 뽑
					duplicated = true
					break

				}
			}

			if duplicated {
				fmt.Println("다시 뽑으세여")
				success = false
				break
			}

			input[idx] = n
			idx++
		}
		if !success {
			continue
		}

		break
	}

	input[0], input[2] = input[2], input[0]
	fmt.Println(input)
	return input
}

func CompareNumbers(numbers, inputNumbers [3]int) Result {
	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
			}
			break
		}
	}

	return Result{strikes, balls}
}

func PrintResult(result Result) {
	fmt.Printf("%dS %dB", result.strikes, result.balls)
}

func IsEndGame(result Result) bool {
	return result.strikes == 3
}
