package main

import (
	"fmt"
	"reflect"
)

func printStudents(students ...string) {
	fmt.Println(students)
	fmt.Println(reflect.TypeOf(students))
}

func main() {
	printStudents("김박사", "오박사", "손정습")
}
