package basicerrors

import (
	"errors"
	"fmt"
)

// ErrorValue는 확인을 위한 패키지 수준 오류를 만드는 방법이다
// 즉, if err == ErrorValue
var ErrorValue = errors.New("this is a typed error")

// Typed Error 는 err.(type) == ErrorValue를 수행할 수 있는 오류 타입을 만드는 방법
type TypedError struct {
	error
}

func BasicErrors() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occurred: %s", "something")
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorValue
	fmt.Println("value error: ", err)

	err = TypedError{
		error: errors.New("typed error"),
	}
	fmt.Println("typed error: ", err)
}
