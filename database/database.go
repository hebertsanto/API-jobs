package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() {

	var err error

	db, err = sql.Open("postgres", "postgresql://docker:docker@localhost:8080/jobs?sslmode=disable")

	if err != nil {
		panic(err)
	}

	fmt.Println("Conected to database successfully!")
}

func GetDB() *sql.DB {
	return db
}
