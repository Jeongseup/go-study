package main

import (
	"fmt"
)

type Rectangle struct {
	width, height int // 자료형이 같은 필드는 한 줄로 나열
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height} // 구조체 인스턴스 생성 뒤 포인터를 리턴
}

//           ↓ 포인터 방식
func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor   // 포인터이므로 원래의 값이 변경됨
	rect.height = rect.height * factor // 포인터이므로 원래의 값이 변경됨
}

//          ↓ 일반 구조체 방식
func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor   // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
	rect.height = rect.height * factor // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
}

//          ↓ 리시버 변수 정의(연결할 구조체 지정)
func (rect *Rectangle) area() int {
	return rect.width * rect.height
	//     ↑  리시버 변수를 사용하여 현재 인스턴스에 접근할 수 있음
}

func main() {
	fmt.Println("")

	var rect1 *Rectangle   // 구조체 포인터 선언
	rect1 = new(Rectangle) // 구조체 포인터에 메모리 할당

	// rect2 := new(Rectangle)               // 구조체 포인터 선언과 동시에 메모리 할당
	var rect2 *Rectangle = new(Rectangle) // 구조체 포인터 선언 후 메모리 할당

	rect1.height = 20 // 구조체 인스턴의 필드에 접근할 때 .을 사용
	rect2.height = 62 // 구초체 포인터에 접근할 때도 .을 사용

	fmt.Println(rect1) // {0 20}: 구조체 인스턴스의 값
	fmt.Println(rect2) // &{0 62}: 구조체 포인터이므로 앞에 &가 붙음

	// 생성자(Constuctor) 흉내
	myRect := NewRectangle(20, 10)
	fmt.Println(myRect)
	fmt.Println(myRect.area())

	myRect.scaleA(10)
	fmt.Println(myRect)
}
