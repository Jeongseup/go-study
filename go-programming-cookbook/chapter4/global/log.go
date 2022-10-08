package global

// UseLog함수는 전역로그를 사용하는 방법을 보여준다
func UseLog() error {
	if err := Init(); err != nil {
		return err
	}

	// 다른 패키지 안에 있는 경우?
	// 이 값은 global.WithField와 Debug를 통해 가져올 수 ㅣ있음
	WithField("key", "value").Debug("hello")
	Debug("test")

	return nil
}
