package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString = ""
	Port             = 0
)

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		Port = 3306
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"))
}
