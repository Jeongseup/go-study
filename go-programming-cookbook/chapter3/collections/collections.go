package collections

// WorkWith 구조체는 컬렉션을 구현하는데 사용
type WorkWith struct {
	Data    string
	Version int
}

// Filter 함수는 함수형 필터
// WorkWith 리스트와 각 true 요소에 대한 bool을 리턴하는 WorkWith 함수를 입력 받아 리스트 반환
func Filter(ws []WorkWith, f func(w WorkWith) bool) []WorkWith {
	// 결과에 따라 result 길이가 0인 경우 크기를 줄인다
	result := make([]WorkWith, 0)
	for _, w := range ws {
		if f(w) {
			result = append(result, w)
		}
	}
	return result
}

// Map함수는 함수형 맵.
// WorkWith의 리스트와
// 수정된 WorkWith를 반환하는 함수를 인자로 받음.
func Map(ws []WorkWith, f func(w WorkWith) WorkWith) []WorkWith {
	// result는 항상 동일한 길이를 가져야 한다
	result := make([]WorkWith, len(ws))

	for pos, w := range ws {
		newW := f(w)
		result[pos] = newW
	}

	return result
}
