package global

import (
	"errors"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// 전역 패키지 수준 변수를 소문자로 만든다
var (
	log     *logrus.Logger
	initLog sync.Once
)

// Init 함수는 처음에 로거를 설정하고 여러번 실행되면 오류를 반환한다
func Init() error {
	err := errors.New("already initialized")
	initLog.Do(func() {
		err = nil
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel
	})

	// fmt.Printf("initialized logger : %v", &log)
	return err
}

// SetLog 함수는 로그를 설정
func SetLog(l *logrus.Logger) {
	log = l
}

// WithField함수는 전역 로그에 WithField함수가 연결된 로그를 내보낸다
func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

// Debug 함수는 전역 로그에 debug함수가 연결된 로그를 내보낸다
func Debug(args ...interface{}) {
	log.Debug(args...)
}
