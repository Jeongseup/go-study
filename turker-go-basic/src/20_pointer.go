package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	name    string
	subject string
	grade   string
}

func (s *Student) ViewScore() {
	fmt.Println(s.subject, ":", s.grade)
}

// func (s *Student) GetScore() {
// 	// return s
// 	fmt.Println(reflect.Type(s))
// }

func (s Student) GetScoreByCopy() Student {
	return s
}

func (s *Student) InsertScore(subject string, grade string) {
	s.subject = subject
	s.grade = grade
}

// go에서는 복사밖에 안되니까 주소를 복사해서 주면? 그 값을 실제로 변경할 수 있게 된다.
func Increase(x *int) {
	// x++
	*x = *x + 1
}

// 포인트의 개념은 변수와 똑같은데
// 단, 실질적인 값이 아닌 그 밸류의 주소를 가리킨다.
func main() {
	// fmt.Println("Hello Go!")

	var a int
	var p *int

	a = 3
	p = &a

	// pointer의 주소 출력
	fmt.Println(a, p)
	// pointer의 주소에 있는 값 출력
	fmt.Println(a, *p)

	Increase(&a)
	fmt.Println(a)

	s := Student{"Jeongseup", "수학", "A"}
	s.ViewScore()

	s.InsertScore("과학", "B+")
	s.ViewScore()

	// memory size test between copy and reference
	fmt.Printf("a: %T, %d\n", s, unsafe.Sizeof(s))
	fmt.Printf("b: %T, %d\n", s, unsafe.Sizeof(s.GetScoreByCopy()))
}
