package main

import "fmt"

func main() {
	// map 초기화
	somemap := map[string]string{"some key1": "some value1"}
	fmt.Println(somemap)

	// map 값 추가
	somemap["some key2"] = "some value2"
	fmt.Println(somemap)

	// map 값 가져오기
	fmt.Println(somemap["some key1"])

	// map 값 제거
	delete(somemap, "some key1")

	fmt.Println("somemap is", somemap)

	// push
	studentAndAges := map[string]int{}
	studentAndAges["김땡땡"] = 13
	studentAndAges["이땡땡"] = 15
	studentAndAges["박땡땡"] = 14
	studentAndAges["최땡땡"] = 17

	// print
	fmt.Println("김땡땡의 나이:", studentAndAges["김땡땡"])
	fmt.Println("이땡땡의 나이:", studentAndAges["이땡땡"])

	// delete
	delete(studentAndAges, "박땡땡")

	// print all
	fmt.Println(studentAndAges)
}
