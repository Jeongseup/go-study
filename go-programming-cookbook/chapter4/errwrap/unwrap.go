package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// Unwrap 함수는 오류의 래핑을 해제하고 오류의 타입 어설션을 수행
func Unwrap() {

	err := error(ErrorTyped{errors.New("an error occured")})
	err = errors.Wrap(err, "wrapped")
	fmt.Println("wrapped error: ", err)

	// 다양한 오류 타입을 처리할 수 있음
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred: ", err)
	default:
		fmt.Println("an unknown error occured")
	}
}

// StackTrace 함수는 오류에 대한 모든 스택을 출력
func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occured")})
	err = errors.Wrap(err, "wrapped")

	// fmt.Println(err)
	fmt.Printf("%+v\n", err)
}
