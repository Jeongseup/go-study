package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("go mysql tutorial")

	// db 객체 얻기
	connStr := "postgres://student:@localhost/study?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	//db.ping() 만약 db연결이 실패할 경우 왜 실패한건지 에러찍는 용도
	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}

	// db가 전부 사용한 다음 맨 마지막에 닫도록 해주는 명령
	defer db.Close()

	// 디비 연결 스터디
	// fmt.Println("successfully connected")

	// 디비 데이터 insert 스터디
	// insert, err := db.Query(`insert into users values('seup')`)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	results, err := db.Query(`select * from users`)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}

	defer results.Close()
}
