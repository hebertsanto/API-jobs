package config

import (
	"log"
	"os"

	"strconv"

	"github.com/joho/godotenv"
)

var (
	Sc = ""

	Port = 0
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 3000
	}

	Sc = os.Getenv("DATABASE_URL")

}
