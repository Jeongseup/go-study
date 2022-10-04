package mylog

import (
	"log"

	"github.com/pkg/errors"
)

// OriginalError함수는 원래 의 오류를 반환
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError 함수는
func PassThroughError() error {
	err := OriginalError()
	// nil 에도 동작하기 때문에 오류를 확인할 필요가 없다?
	return errors.Wrap(err, "in passthrougherror")

}

func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// 예상치 못한 오류가 발생해 로그를 기록
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}
