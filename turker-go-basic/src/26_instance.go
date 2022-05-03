package main

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	grade int
}

func SetNameByFunction(t *Student, newName string) {
	t.name = newName
}

func (t *Student) SetNameByMethod(newName string) {
	t.name = newName
}

func main() {
	fmt.Println("")

	a := Student{"seup", 27, 10}
	fmt.Println("before a is", a)

	var b *Student
	b = &a

	b.age = 30

	fmt.Println(a)
	fmt.Println(b)

	// 그냥 넣으면 안 바뀌고, a의 주소를 대입해줘야 set함수가 실행될 수 있다.
	// 주소값이 복사되기 때문에 참조가 가능해진다.
	// 기능 단위로 볼 것이냐?
	SetNameByFunction(&a, "bbb")
	fmt.Println(a)

	// 관계가 있는 단위로 볼 것이냐?
	// OOP에서 ER(Entity Relationship)이 중요함.
	a.SetNameByMethod("ccc")
	fmt.Println(a)
}
