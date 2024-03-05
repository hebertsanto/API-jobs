package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() {

	var err error

	DB_URL := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", DB_URL)

	if err != nil {
		panic("some error ocurred trying to connect to database" + err.Error())
	}

	fmt.Println("Conected to database successfully!")
}

func GetDB() *sql.DB {
	return db
}
