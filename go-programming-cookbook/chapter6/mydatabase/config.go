package mydatabase

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-pg/pg"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "student"
	dbname = "gocookbook"
)

// Example 구조체는 쿼리에 대한 결과를 저장
type Example struct {
	Name    string
	Created *time.Time
}

// Setup함수는 데이터베이스를 설정하고 연결 풀이 설정된 데이터베이스를 반환
func Setup() (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	// close database
	// defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected!")

	return db, nil
}
