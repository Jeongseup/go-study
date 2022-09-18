package flags

import (
	"flag"
	"fmt"
)

type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

// Setup 함수는 전달되는 플래그 값으로 config 초기화
func (c *Config) Setup() {
	// 다음과 같이 flag를 직접 설정 가능
	// var someVar = flag.Sring("flag_name", "default_val", "description")
	// 하지만!? 구조체에 넣는게 더 낫다

	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty(shortterm)") // 단축어로도 설정가능한듯?
	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome or what?")
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awsome out of 10?")
	// 커스텀 변수 타입
	flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

// GetMessage 함순느 모든 내부 config 변수를 사용해 문장을 반환
func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg += " is awesome"
	} else {
		msg += " is NOT awesome"
	}
	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s",
		msg,
		c.howAwesome,
		c.countTheWays.String(),
	)

	return msg
}
