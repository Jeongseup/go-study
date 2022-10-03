package mymath

import "math/big"

// 피보나치 수열을 저장하기 위한 글로벌 변수
var memorize map[int]*big.Int

func init() {
	// 맵 초기화
	memorize = make(map[int]*big.Int)
}

// Fib 함수는 피보나치 수열의 n번째 수를 출력
// 0보다 작은 경우 1을 반환
func Fib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}

	// 기본 케이스
	if n < 2 {
		memorize[n] = big.NewInt(1)
	}

	// 이전에 저장했는지 확인
	// 이미 저장한 경우, 계산하지 않고 저장된 값을 반환
	if val, ok := memorize[n]; ok {
		return val
	}

	// 맵을 초기화한 다음, 이전 두 개의 피보나치 값을 더한다
	memorize[n] = big.NewInt(0) // 초기화
	memorize[n].Add(memorize[n], Fib(n-1))
	memorize[n].Add(memorize[n], Fib(n-2))

	return memorize[n]
}
