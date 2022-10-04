package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// WrappedError 함수는 오류 래핑과 오류에 주석을 표시하는 방법을 보여줌
func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred int WrappedError")
}

// ErrorTyped 구조체는 예제에서 확인할 오류
type ErrorTyped struct {
	error
}

// Wrap 함수는 오류를 래핑할 때 호출
func Wrap() {
	e := errors.New("standard error")
	fmt.Println("Regular Error - ", WrappedError(e))

	fmt.Println("Typed Error - ", WrappedError(ErrorTyped{errors.New("typed error")}))

	fmt.Println("Nil - ", WrappedError(nil))
}
