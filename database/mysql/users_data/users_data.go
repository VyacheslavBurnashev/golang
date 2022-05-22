package users_data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Client   *sql.DB
	username = "root"
	host     = "127.0.0.1:3306"
	schema   = "users"
)

func init() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("can't loading .env file")
	}
	password := os.Getenv("DB_PASSWORD")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("db is running")
}
