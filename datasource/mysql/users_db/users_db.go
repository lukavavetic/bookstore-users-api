package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	MYSQL_USER   = "DB_USER"
	MYSQL_PASS   = "DB_PASS"
	MYSQL_HOST   = "DB_HOST"
	MYSQL_SCHEMA = "DB_SCHEMA"

)

var (
	Client *sql.DB

	DbUser   = os.Getenv(MYSQL_USER)
	DbPass   = os.Getenv(MYSQL_PASS)
	DbHost   = os.Getenv(MYSQL_HOST)
	DbSchema = os.Getenv(MYSQL_SCHEMA)
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		DbUser, DbPass, DbHost, DbSchema,
	)

	var err error
	Client, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	println("success")
}