package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

// 아... 이거 그냥 input?
// inputName inputType
func (human Person) printName() {
	fmt.Println(human.name)
}

type Student struct {
	name  string
	class int
	score Score
}

func (s Student) ViewScore() {
	fmt.Println(s.score)
}

func (s Student) InsertScore(name string, grade string) {
	s.score.name = name
	s.score.grade = grade
}

// func InsertScore(student Student, name string, grade string) {
// 	student.score.name = name
// 	student.score.grade = grade
// }

type Score struct {
	name  string
	grade string
}

func main() {
	// var p Person
	// p1 := Person{"John", 17}
	// p2 := Person{name: "David", age: 20}
	// p3 := Person{age: 24}

	// fmt.Println(p1, p2, p3)
	// p.name = "Jeongseup"

	// p.printName()
	// p2.printName()

	var student Student

	student.class = 1
	student.name = "Jeongseup"
	student.score.name = "수학"
	student.score.grade = "A"

	student.ViewScore()
	student.InsertScore("과학", "B")

	// golang에서는 함수를 호출하는 원리가 무조건 복사이기 때문에
	// 일반적인 set함수를 쓰더라도 실제 그 값이 넣어지는 struct는 내가 생각했던 그 구조체가 아닌 복사된 또 다른 구조체의 값을 변경하는 것 뿐이다.
	// 그래서 원래 구조체의 값이 변경되지 않는다.
	student.ViewScore()

}
