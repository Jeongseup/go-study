package currency

import "strconv"

// ConvertPenniesToDollarString
func ConvertPenniesToDollarString(amount int64) string {
	// 입력받은 페니의 양을 10진수로 변환
	result := strconv.FormatInt(amount, 10)

	// 음수인 경우 나중에 설정?
	negative := false
	if result[0] == '-' {
		result = result[1:]
		negative = true
	}

	// 값이 100보다 작으면 왼쪽에 0을 하나 추가
	for len(result) < 3 {
		result = "0" + result
	}

	length := len(result)

	// 10 진수 덧셈
	result = result[0:length-2] + "." + result[length-2:]

	// 음수인 경우 -를 붙여준다
	if negative {
		result = "-" + result
	}

	return result

}
