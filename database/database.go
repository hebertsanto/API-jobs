package database

import (
	"database/sql"
	"fmt"
	"vagas/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() {

	var err error

	db, err = sql.Open("postgres", config.Sc)

	if err != nil {
		panic("some error ocurred trying to connect to database" + err.Error())
	}

	fmt.Println("Conected to database successfully!")
}

func GetDB() *sql.DB {
	return db
}
