package tags

import (
	"fmt"
	"reflect"
)

// SerializeStructStrings 함수는 구조체를 사용자 정의 직렬화? 포맷으로 변환
// 문자열 타입에 대한 구조체 태그를 직렬화하는 방식을 사용
func SerializeStructStrings(s interface{}) (string, error) {
	result := ""

	// 인터페이스를 특정 타입으로 반영
	r := reflect.TypeOf(s)
	// fmt.Println("here", r)

	value := reflect.ValueOf(s)
	// fmt.Println("here", value)

	// fmt.Println("here", r.Kind())
	// fmt.Println("here", r, value)

	// 구조체에 대한 포인터가 전달된 경우?
	// 적절히 처리
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		value = value.Elem()
	}

	fmt.Println("here", r, value)
	// 모든 필드에 대해 루프 처리
	for i := 0; i < r.NumField(); i++ {

		field := r.Field(i)
		// fmt.Println("current field :", field)

		// 구조체 태그가 발견된 경우
		key := field.Name
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// "-"를 무시
			// 그렇지 않으면 전체 값이 직렬화 키가 된다..?
			if serialize == "-" {
				continue
			}
			key = serialize
		}

		switch value.Field(i).Kind() {
		// 이 예제에서는 문자열만 지원?
		case reflect.String:
			result += key + ":" + value.Field(i).String() + ";"
			// 그외 타입은 건너뜀
		default:
			continue
		}
	}

	return result, nil
}
