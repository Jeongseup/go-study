package mydatabase

import "database/sql"

// Exec 함수는 이전 예제의 exec함수를 대체
func Exec(db *sql.DB) error {
	defer db.Exec("DROP TABLE example")
	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db, "Aaron"); err != nil {
		return err
	}

	return nil
}
