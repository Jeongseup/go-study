package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "student"
	password = "" // NOT REQUIERD
	dbname   = "study"
)

type Post struct {
}

func (post *Post) Create(db *sql.DB, statement string) (err error) {
	// insert
	insertStmt := `insert into "students"("name", "roll") values('john', 'non')`
	// https://www.youtube.com/watch?v=sxaFWBggS3U

	stmt, err := db.Prepare(statement)
	defer stmt.Close()

	err = stmt.QueryRow(post.Content)
	_, e := db.Exec(insertStmt)
	CheckError(e)

}

func main() {
	// db connection
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db open
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// db close
	defer db.Close()

	// db check
	err = db.Ping()
	CheckError(err)
	fmt.Println("connected!")

	statement := `insert into students (name, roll) values($1, $2) returning id`

	// dynamic
	//insertDynStmt := `insert into "students"("name", "roll") values($1, $2)`
	// _, err = db.Exec(insertDynStmt, "test", "2")
	// CheckError(err)

	// update
	// updateStmt := `update "students" set "name"=$1, "roll"=$2 where "id"=$3`
	// _, e := db.Exec(updateStmt, "Mary", 10, 2)
	// CheckError(e)

	// Delete
	// deleteStmt := `delete from "students" where id=$1`
	// _, e := db.Exec(deleteStmt, 1)
	// CheckError(e)
	rows, err := db.Query(`select * from "students"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var roll string

		err = rows.Scan(&id, &name, &roll)
		CheckError(err)

		fmt.Println(id, name, roll)
		fmt.Println(1)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
