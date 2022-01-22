package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	MYSQL_USER     string = "root"
	MYSQL_PASS     string = "secret"
	MYSQL_DATABASE string = "tugas_15_b31"
	MYSQL_HOST            = "localhost"
	MYSQL_PORT            = 33769
)

var (
	dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PASS, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)
)

func DbUrl() string {
	return dsn
}
