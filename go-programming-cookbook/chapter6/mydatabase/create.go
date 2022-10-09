package mydatabase

import (
	"database/sql"

	// _ "github.com/lib/pq"
	_ "github.com/go-pg/pg"
)

// Create함수는 example 이라는 이름의 테이블을 생성하고 데이터를 채운다
func Create(db *sql.DB) error {
	// DB 생성
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`); err != nil {
		return err
	}
	return nil
}
