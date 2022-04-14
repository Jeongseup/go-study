package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// 문자열 타입
	str := "123"
	fmt.Printf("%s is string? \n", str)
	fmt.Println("Yes!", reflect.TypeOf(str))

	// string to int
	// 타입 변환 이후 err인지 아닌 디텍트 할 수 있음
	i, err := strconv.Atoi(str) // ParseInt랑 다른건가? ASCII to Int
	checkErr(nil)
	fmt.Println(i, err)
	fmt.Println(reflect.TypeOf(i))

	// string to float64
	f, err := strconv.ParseFloat(str, 64)
	checkErr(err)
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

	// string to bool
	b, err := strconv.ParseBool("true")
	checkErr(err)
	fmt.Println(b, reflect.TypeOf(b))

	// string to []rune
	// []는 slice인데 좀 있다가 배울 겁니다.
	rs := []rune(str)
	fmt.Println(rs, rune('a'))

	///// other types to string. /////
	// int to string
	s := strconv.Itoa(1) // Int to ASCII?
	fmt.Println(s, reflect.TypeOf(s))

	// bool to string
	s = strconv.FormatBool(true)
	fmt.Println(s, reflect.TypeOf(s))

	// []rune to string
	runes := []rune{'a', 'b', 'c'}
	fmt.Println(runes, reflect.TypeOf(runes))
	s = string(runes)
	fmt.Println(s)

	// float64 to string은 변환하는 것이 복잡하여 생략하겠습니다.
	// 궁금하시다면 구글링해보세요!
}
