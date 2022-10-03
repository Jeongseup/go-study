package currency

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ConvertStringDollarsToPennies 함수는 문자열로 달러의 양을 입력 받아 int64로 변환
func ConvertStringDollarsToPennies(amount string) (int64, error) {
	// amount 인자가 float 변환이 가능한지 체크
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	// . 문자를 기준으로 값을 나눔
	groups := strings.Split(amount, ".")
	fmt.Println(groups)
	// . 문자가 없는 경우 다음 코드에서 값을 result에 저장
	result := groups[0]
	// 기본 문자열?
	r := ""
	// "." 다음의 데이터 처리
	if len(groups) == 2 {
		if len(groups[1]) != 2 {
			return 0, errors.New("invalid cents")
		}
		r = groups[1]
	}

	// 0으로 채운다 . 문자가 없는경우
	// 두 개의 0이 채워질 것
	for len(r) < 2 {
		r += "0"
	}

	result += r

	// int로 변환
	return strconv.ParseInt(result, 10, 64) // (int, err)
}
