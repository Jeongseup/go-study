package main

import "fmt"

func main() {
	health := 10
	if health > 10 {
		fmt.Println("health 변수가 최대치를 넘었습니다. 10으로 재조정 됩니다.")
		health = 10
	} else if health == 10 {
		fmt.Println("health 변수가 최대치 입니다.")
	} else if health <= 3 {
		fmt.Println("health 변수가 3 이하 입니다. 건강하세요!")
	}
	fmt.Println("health:", health)
}
