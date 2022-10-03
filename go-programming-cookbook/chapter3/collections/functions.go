package collections

import "strings"

// LowerCaseData 함수는 WorkWith의 Data 문자열에
// ToLower을 수행
func LowerCaseData(w WorkWith) WorkWith {
	w.Data = strings.ToLower(w.Data)
	return w
}

// IncrementVersion 함수는 WorkWiths의 버젼을 증가
func IncrementVersion(w WorkWith) WorkWith {
	w.Version++
	return w
}

// OldVersion 함수는 버전이 지정된 값보다 큰지 확인하는 클로저?를 반환
func OldVersion(v int) func(w WorkWith) bool {
	return func(w WorkWith) bool {
		return w.Version >= v
	}
}
