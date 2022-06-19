package main

import (
	"fmt"
	"strconv"
)

type hello interface {
}
type Rectangle struct { // 사각형 구조체 정의
	width, height int
}

func (r Rectangle) Print() { // Rectangle에 Print 메서드를 연결
	fmt.Println(r.width, r.height)
}

type MyInt int // int형을 MyInt로 정의

func (i MyInt) Print() { // MyInt에 Print 메서드를 연결
	fmt.Println(i)
}

type Printer interface { // Print 메서드를 가지는 인터페이스 정의
	Print()
}

type Any interface{}

func formatString(arg Any) string {
	//       ↓ 인터페이스에 저장된 타입에 따라 case 실행
	fmt.Println(arg)

	switch arg.(type) { // 단 이 방법은 switch 분기문 안에서만 사용할 수 있고, 일반적인 방법으로는 사용할 수 없습니다.
	case int: // arg가 int형이라면
		i := arg.(int)         // int형으로 값을 가져옴
		return strconv.Itoa(i) // strconv.Itoa 함수를 사용하여 i를 문자열로 변환
	case float32: // arg가 float32형이라면
		f := arg.(float32) // float32형으로 값을 가져옴
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
		// strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case float64: // arg가 float64형이라면
		f := arg.(float64) // float64형으로 값을 가져옴
		return strconv.FormatFloat(f, 'f', -1, 64)
		// strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case string: // arg가 string이라면
		s := arg.(string) // string으로 값을 가져옴
		return s          // string이므로 그대로 리턴
	default:
		return "Error"
	}
}

func main() {
	var h hello
	fmt.Println(h)

	var i MyInt = 5
	r := Rectangle{10, 20}

	pArr := []Printer{i, r} // 슬라이스 형태로 인터페이스 초기화

	for _, value := range pArr {
		value.Print() // 슬라이스를 순회하면서 Print 메서드 호출
	}

	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello, world!"))

	// 특정 타입인지 검사
	var t interface{}
	t = Rectangle{10, 20}
	if v, ok := t.(Rectangle); ok {
		fmt.Println(v, ok)
	}
}
