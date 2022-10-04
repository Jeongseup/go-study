package basicerrors

import "fmt"

// CustomErrors는 Error() 인터페이스를 구현할 구조체
type CustomError struct {
	Result string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("there was an error; %s was the result", c.Result)
}

// SomeFunc 함수는 오류를 반ㅇ환
func SomeFunc() error {
	c := CustomError{Result: "this"}
	return c
}
