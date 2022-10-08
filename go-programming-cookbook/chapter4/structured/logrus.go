package structured

import "github.com/sirupsen/logrus"

// Hook 구조체는 logrus의 hook 인터페이스를 구현
type Hook struct {
	id string
}

// Fire함수는 로그를 기록할 때마다 호출
func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.id
	entry.Data["test"] = "test"
	return nil
}

// Levels 함수는 전달된 hook 이 실행될 레벨
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Logrus 함수는 몇 가지 기본적인 logrus 기능을 보여줌
func Logrus() {
	// json 포맷 로그를 기록
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happend", "Just now"}

	x := logrus.WithFields(fields)
	x.Warn("warning!")
	x.Error("error!")
	x.Infof("this is info")
}
